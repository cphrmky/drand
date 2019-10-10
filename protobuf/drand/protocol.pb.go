// Code generated by protoc-gen-go. DO NOT EDIT.
// source: drand/protocol.proto

package drand // import "github.com/dedis/drand/protobuf/drand"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import dkg "github.com/dedis/drand/protobuf/crypto/dkg"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// BeaconRequest holds a link to a previous signature, a timestamp and the
// partial signature for this beacon. All participants send and collects many of
// theses partial beacon packets to recreate locally one beacon
type BeaconRequest struct {
	Round       uint64 `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	PreviousSig []byte `protobuf:"bytes,2,opt,name=previous_sig,json=previousSig,proto3" json:"previous_sig,omitempty"`
	// To prove the issuer comes from a valid node in the group
	// It is a group point prefixed by the index of the issuer
	PartialSig           []byte   `protobuf:"bytes,3,opt,name=partial_sig,json=partialSig,proto3" json:"partial_sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BeaconRequest) Reset()         { *m = BeaconRequest{} }
func (m *BeaconRequest) String() string { return proto.CompactTextString(m) }
func (*BeaconRequest) ProtoMessage()    {}
func (*BeaconRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_protocol_e5f7e0a88bc2ab94, []int{0}
}
func (m *BeaconRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeaconRequest.Unmarshal(m, b)
}
func (m *BeaconRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeaconRequest.Marshal(b, m, deterministic)
}
func (dst *BeaconRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeaconRequest.Merge(dst, src)
}
func (m *BeaconRequest) XXX_Size() int {
	return xxx_messageInfo_BeaconRequest.Size(m)
}
func (m *BeaconRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BeaconRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BeaconRequest proto.InternalMessageInfo

func (m *BeaconRequest) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *BeaconRequest) GetPreviousSig() []byte {
	if m != nil {
		return m.PreviousSig
	}
	return nil
}

func (m *BeaconRequest) GetPartialSig() []byte {
	if m != nil {
		return m.PartialSig
	}
	return nil
}

type BeaconResponse struct {
	PartialSig           []byte   `protobuf:"bytes,1,opt,name=partial_sig,json=partialSig,proto3" json:"partial_sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BeaconResponse) Reset()         { *m = BeaconResponse{} }
func (m *BeaconResponse) String() string { return proto.CompactTextString(m) }
func (*BeaconResponse) ProtoMessage()    {}
func (*BeaconResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_protocol_e5f7e0a88bc2ab94, []int{1}
}
func (m *BeaconResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeaconResponse.Unmarshal(m, b)
}
func (m *BeaconResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeaconResponse.Marshal(b, m, deterministic)
}
func (dst *BeaconResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeaconResponse.Merge(dst, src)
}
func (m *BeaconResponse) XXX_Size() int {
	return xxx_messageInfo_BeaconResponse.Size(m)
}
func (m *BeaconResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BeaconResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BeaconResponse proto.InternalMessageInfo

func (m *BeaconResponse) GetPartialSig() []byte {
	if m != nil {
		return m.PartialSig
	}
	return nil
}

type SetupPacket struct {
	Dkg                  *dkg.Packet `protobuf:"bytes,1,opt,name=dkg,proto3" json:"dkg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SetupPacket) Reset()         { *m = SetupPacket{} }
func (m *SetupPacket) String() string { return proto.CompactTextString(m) }
func (*SetupPacket) ProtoMessage()    {}
func (*SetupPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_protocol_e5f7e0a88bc2ab94, []int{2}
}
func (m *SetupPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetupPacket.Unmarshal(m, b)
}
func (m *SetupPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetupPacket.Marshal(b, m, deterministic)
}
func (dst *SetupPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetupPacket.Merge(dst, src)
}
func (m *SetupPacket) XXX_Size() int {
	return xxx_messageInfo_SetupPacket.Size(m)
}
func (m *SetupPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_SetupPacket.DiscardUnknown(m)
}

var xxx_messageInfo_SetupPacket proto.InternalMessageInfo

func (m *SetupPacket) GetDkg() *dkg.Packet {
	if m != nil {
		return m.Dkg
	}
	return nil
}

// Reshare is a wrapper around a Setup packet for resharing operation that
// serves two purposes: - indicate to non-leader old nodes that they should
// generate and send their deals - indicate to which new group are we resharing.
// drand should keep a list of new ready-to-operate groups allowed.
type ResharePacket struct {
	Dkg                  *dkg.Packet `protobuf:"bytes,1,opt,name=dkg,proto3" json:"dkg,omitempty"`
	GroupHash            string      `protobuf:"bytes,2,opt,name=group_hash,json=groupHash,proto3" json:"group_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ResharePacket) Reset()         { *m = ResharePacket{} }
func (m *ResharePacket) String() string { return proto.CompactTextString(m) }
func (*ResharePacket) ProtoMessage()    {}
func (*ResharePacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_protocol_e5f7e0a88bc2ab94, []int{3}
}
func (m *ResharePacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResharePacket.Unmarshal(m, b)
}
func (m *ResharePacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResharePacket.Marshal(b, m, deterministic)
}
func (dst *ResharePacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResharePacket.Merge(dst, src)
}
func (m *ResharePacket) XXX_Size() int {
	return xxx_messageInfo_ResharePacket.Size(m)
}
func (m *ResharePacket) XXX_DiscardUnknown() {
	xxx_messageInfo_ResharePacket.DiscardUnknown(m)
}

var xxx_messageInfo_ResharePacket proto.InternalMessageInfo

func (m *ResharePacket) GetDkg() *dkg.Packet {
	if m != nil {
		return m.Dkg
	}
	return nil
}

func (m *ResharePacket) GetGroupHash() string {
	if m != nil {
		return m.GroupHash
	}
	return ""
}

func init() {
	proto.RegisterType((*BeaconRequest)(nil), "drand.BeaconRequest")
	proto.RegisterType((*BeaconResponse)(nil), "drand.BeaconResponse")
	proto.RegisterType((*SetupPacket)(nil), "drand.SetupPacket")
	proto.RegisterType((*ResharePacket)(nil), "drand.ResharePacket")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProtocolClient is the client API for Protocol service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProtocolClient interface {
	// Setup is doing the DKG setup phase
	Setup(ctx context.Context, in *SetupPacket, opts ...grpc.CallOption) (*Empty, error)
	// Reshare performs the resharing phase
	Reshare(ctx context.Context, in *ResharePacket, opts ...grpc.CallOption) (*Empty, error)
	// NewBeacon asks for a partial signature to another node
	NewBeacon(ctx context.Context, in *BeaconRequest, opts ...grpc.CallOption) (*BeaconResponse, error)
}

type protocolClient struct {
	cc *grpc.ClientConn
}

func NewProtocolClient(cc *grpc.ClientConn) ProtocolClient {
	return &protocolClient{cc}
}

func (c *protocolClient) Setup(ctx context.Context, in *SetupPacket, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/drand.Protocol/Setup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) Reshare(ctx context.Context, in *ResharePacket, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/drand.Protocol/Reshare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) NewBeacon(ctx context.Context, in *BeaconRequest, opts ...grpc.CallOption) (*BeaconResponse, error) {
	out := new(BeaconResponse)
	err := c.cc.Invoke(ctx, "/drand.Protocol/NewBeacon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProtocolServer is the server API for Protocol service.
type ProtocolServer interface {
	// Setup is doing the DKG setup phase
	Setup(context.Context, *SetupPacket) (*Empty, error)
	// Reshare performs the resharing phase
	Reshare(context.Context, *ResharePacket) (*Empty, error)
	// NewBeacon asks for a partial signature to another node
	NewBeacon(context.Context, *BeaconRequest) (*BeaconResponse, error)
}

func RegisterProtocolServer(s *grpc.Server, srv ProtocolServer) {
	s.RegisterService(&_Protocol_serviceDesc, srv)
}

func _Protocol_Setup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupPacket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).Setup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drand.Protocol/Setup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).Setup(ctx, req.(*SetupPacket))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_Reshare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResharePacket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).Reshare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drand.Protocol/Reshare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).Reshare(ctx, req.(*ResharePacket))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_NewBeacon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeaconRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).NewBeacon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drand.Protocol/NewBeacon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).NewBeacon(ctx, req.(*BeaconRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Protocol_serviceDesc = grpc.ServiceDesc{
	ServiceName: "drand.Protocol",
	HandlerType: (*ProtocolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Setup",
			Handler:    _Protocol_Setup_Handler,
		},
		{
			MethodName: "Reshare",
			Handler:    _Protocol_Reshare_Handler,
		},
		{
			MethodName: "NewBeacon",
			Handler:    _Protocol_NewBeacon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "drand/protocol.proto",
}

func init() { proto.RegisterFile("drand/protocol.proto", fileDescriptor_protocol_e5f7e0a88bc2ab94) }

var fileDescriptor_protocol_e5f7e0a88bc2ab94 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xa9, 0x73, 0xea, 0xde, 0x6e, 0x82, 0xa1, 0xc2, 0x28, 0x0c, 0xe7, 0x40, 0x9c, 0xa0,
	0x2d, 0xce, 0x8b, 0xe7, 0x81, 0xe0, 0x45, 0x19, 0xdd, 0xcd, 0xcb, 0xc8, 0x9a, 0x98, 0x86, 0x6d,
	0x4d, 0xcc, 0x1f, 0x65, 0x5f, 0xc6, 0xcf, 0x2a, 0x4d, 0x3a, 0xd9, 0xe6, 0xc5, 0x43, 0xa1, 0xef,
	0xef, 0x7d, 0x1e, 0xfa, 0xe4, 0x69, 0x20, 0x22, 0x0a, 0x97, 0x24, 0x95, 0x4a, 0x18, 0x91, 0x8b,
	0x65, 0xe2, 0x5e, 0x50, 0xd3, 0xd1, 0x38, 0xca, 0xd5, 0x5a, 0x1a, 0x91, 0x92, 0x05, 0xab, 0x1e,
	0xbf, 0x8c, 0xcf, 0xbc, 0x85, 0xae, 0xa4, 0x59, 0x7b, 0x34, 0xe0, 0xd0, 0x19, 0x53, 0x9c, 0x8b,
	0x32, 0xa3, 0x1f, 0x96, 0x6a, 0x83, 0x22, 0x68, 0x2a, 0x61, 0x4b, 0xd2, 0x0d, 0xfa, 0xc1, 0xf0,
	0x30, 0xf3, 0x03, 0xba, 0x84, 0xb6, 0x54, 0xf4, 0x93, 0x0b, 0xab, 0x67, 0x9a, 0xb3, 0xee, 0x41,
	0x3f, 0x18, 0xb6, 0xb3, 0x70, 0xc3, 0xa6, 0x9c, 0xa1, 0x0b, 0x08, 0x25, 0x56, 0x86, 0xe3, 0xa5,
	0x53, 0x34, 0x9c, 0x02, 0x6a, 0x34, 0xe5, 0x6c, 0x70, 0x0f, 0xa7, 0x9b, 0x4f, 0x69, 0x29, 0x4a,
	0x4d, 0xf7, 0x2d, 0xc1, 0x1f, 0xcb, 0x2d, 0x84, 0x53, 0x6a, 0xac, 0x9c, 0xe0, 0x7c, 0x41, 0x0d,
	0xea, 0x41, 0x83, 0x2c, 0xbc, 0x2e, 0x1c, 0x85, 0x49, 0x75, 0x30, 0xbf, 0xc9, 0x2a, 0x3e, 0x78,
	0x81, 0x4e, 0x46, 0x75, 0x81, 0x15, 0xfd, 0x97, 0x1e, 0xf5, 0x00, 0x98, 0x12, 0x56, 0xce, 0x0a,
	0xac, 0x0b, 0x77, 0xa4, 0x56, 0xd6, 0x72, 0xe4, 0x19, 0xeb, 0x62, 0xf4, 0x1d, 0xc0, 0xc9, 0xa4,
	0x6e, 0x17, 0xdd, 0x40, 0xd3, 0x25, 0x41, 0x28, 0x71, 0x25, 0x26, 0x5b, 0xb9, 0xe2, 0x76, 0xcd,
	0x9e, 0xaa, 0x62, 0xd1, 0x1d, 0x1c, 0xd7, 0x31, 0x50, 0x54, 0x2f, 0x76, 0x62, 0xed, 0xc9, 0x1f,
	0xa1, 0xf5, 0x4a, 0xbf, 0x7c, 0x33, 0xbf, 0x86, 0x9d, 0x7f, 0x12, 0x9f, 0xef, 0x51, 0x5f, 0xdf,
	0xf8, 0xfa, 0xed, 0x8a, 0x71, 0x53, 0xd8, 0x79, 0x92, 0x8b, 0x55, 0x4a, 0x28, 0xe1, 0x3a, 0xdd,
	0xba, 0x14, 0x73, 0xfb, 0xee, 0xc7, 0xf9, 0x91, 0x9b, 0x1f, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x43, 0x13, 0x9d, 0xd7, 0x33, 0x02, 0x00, 0x00,
}
