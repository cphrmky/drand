package main

import (
	"bytes"
	"encoding/binary"
	"sync"
	"time"

	"gopkg.in/dedis/kyber.v1/share"

	"github.com/dedis/drand/bls"
	"github.com/nikkolasg/slog"
)

// How much time can a signature timestamp differ from our local time
var maxTimestampDelta = 10 * time.Second

// Beacon holds the logic to initiate, and react to the TBLS protocol. Each time
// a full signature can be recosntructed, it saves it to the given Store.
type Beacon struct {
	r         *Router
	share     *Share
	group     *Group
	pub       *share.PubPoly
	store     Store
	threshold int
	sync.Mutex

	pendingSigs map[string][]*bls.ThresholdSig

	ticker *time.Ticker
}

// newBlsBeacon
func newBlsBeacon(sh *Share, group *Group, r *Router, s Store) *Beacon {
	return &Beacon{
		r:         r,
		group:     group,
		share:     sh,
		pub:       share.NewPubPoly(g2, g2.Point().Base(), sh.Commits),
		threshold: len(sh.Commits),
		store:     s,
	}
}

// RandomBeacon starts periodically the TBLS protocol. The seed is the first
// message signed alongside with the current timestamp. All subsequent
// signatures are chained:
// s_i+1 = SIG(s_i || timestamp)
// For the moment, each resulting signature is stored in a file named
// beacons/<timestamp>.sig (because of FileStore).
func (b *Beacon) Start(seed []byte, period time.Duration) {
	b.Lock()
	b.ticker = time.NewTicker(period)
	b.Unlock()

	var msg = seed
	var counter uint64 = 1
	var failed uint64
	for _ = range b.ticker.C {
		now := time.Now().Unix()
		b.genPartialSignature(msg, now)
		packet := &DrandPacket{
			Beacon: &BeaconPacket{
				Request: &BeaconRequest{
					PreviousSig: msg,
					Timestamp:   now,
				},
			},
		}
		if err := b.r.Broadcast(b.group, packet); err != nil {
			failed++
			slog.Infof("beacon: start round %d failed (%d total failed)", counter, failed)
			continue
		}
		counter++
		slog.Infof("beacon: start round %d correct", counter)
	}
}

// processBeaconPacket looks if the packet is a signature request or a signature
// reply and acts accordingly.
func (b *Beacon) processBeaconPacket(pub *Public, msg *BeaconPacket) {
	switch {
	case msg.Request != nil:
		b.processBeaconRequest(pub, msg.Request)
	case msg.Reply != nil:
		b.processBeaconSignature(pub, msg.Reply)
	default:
		slog.Info("beacon received unknown bls beacon message")
	}
}

// processBeaconRequest process the beacon packet in two steps:
// 1- verify that the new timestamp is close enough to our time
// 2- generates and saves a new threshold partial signature for
//    the new message m_i = H(sig_i-1 || timestamp)
// 3- broadcast that partial signature to the whole group
func (b *Beacon) processBeaconRequest(pub *Public, msg *BeaconRequest) {
	// 1
	now := time.Now()
	leaderTime := time.Unix(msg.Timestamp, 0)
	if now.Sub(leaderTime) > maxTimestampDelta {
		slog.Info("blsbeacon received out-of-range timestamp signature request: ", now.Sub(leaderTime))
		return
	}
	// 2-
	sig := b.genPartialSignature(msg.PreviousSig, msg.Timestamp)
	packet := &DrandPacket{
		Beacon: &BeaconPacket{
			Reply: &BeaconReply{
				Request:   msg,
				Signature: sig,
			},
		},
	}
	// 3-
	go func() {
		if err := b.r.Broadcast(b.group, packet); err != nil {
			slog.Info("blsBeacon error broadcast partial signature: ", err)
		}
	}()
}

// processBeaconSignature does the following:
// 1- checks if the given partial signature is valid.
// 2- check if we already recovered the full signatures (by looking at the
// signature folder)
// 3- Saves it in memory and if there is enough threshold partial signatures for
// the message, it reconstructs the full bls signature and saves it to a file.
func (b *Beacon) processBeaconSignature(pub *Public, sig *BeaconReply) {
	b.Lock()
	defer b.Unlock()
	// 1-
	msg := message(sig.Request.PreviousSig, sig.Request.Timestamp)
	if !bls.ThresholdVerify(pairing, b.pub, msg, sig.Signature) {
		slog.Info("blsBeacon received invalid partial signature")
		return
	}

	// 2-
	if b.store.SignatureExists(sig.Request.Timestamp) {
		slog.Infof("blsBeacon already reconstructed signature %d", sig.Request.Timestamp)
		return
	}

	d := digest(msg)
	for _, s := range b.pendingSigs[d] {
		if s.Index == sig.Signature.Index {
			slog.Debug("blsbeacon already received partial signature for same message")
			return
		}
	}
	b.pendingSigs[d] = append(b.pendingSigs[d], sig.Signature)

	// 3-
	if len(b.pendingSigs[d]) < b.threshold {
		slog.Debugf("blsBeacon: not enough partial signature yet %d/%d", len(b.pendingSigs[d]), b.threshold)
		return
	}

	slog.Debug("blsBeacon: full signature recovery")
	fullSig, err := bls.AggregateSignatures(pairing, b.pub, msg, b.pendingSigs[d], len(b.group.List), b.threshold)
	if err != nil {
		slog.Info("blsBeacon: full signature recovery failed for ts %d: %s", sig.Request.Timestamp, err)
		return
	}
	delete(b.pendingSigs, d)

	if b.store.SaveSignature(NewBeaconSignature(sig.Request, fullSig)); err != nil {
		slog.Infof("blsBeacon: error saving signature: %s", err)
		return
	}
	slog.Print("blsBeacon: reconstructed and save full signature")
}

func (b *Beacon) Stop() {
	b.Lock()
	defer b.Unlock()
	b.ticker.Stop()
}

func (b *Beacon) genPartialSignature(oldSig []byte, time int64) *bls.ThresholdSig {
	newMessage := message(oldSig, time)
	thresholdSign := bls.ThresholdSign(pairing, b.share.Share, newMessage)
	b.Lock()
	defer b.Unlock()
	digestM := digest(newMessage)
	b.pendingSigs[digestM] = append(b.pendingSigs[digestM], thresholdSign)
	return thresholdSign
}

// message returns the message out of the signature and the timestamp as what
// gets signed during a round of the TBLS protocol.
func message(previousSig []byte, ts int64) []byte {
	var buff bytes.Buffer
	binary.Write(&buff, binary.LittleEndian, ts)
	buff.Write(previousSig)
	return buff.Bytes()
}

// digest returns a compact representation of the given message
func digest(msg []byte) string {
	return string(pairing.Hash().Sum(msg))
}