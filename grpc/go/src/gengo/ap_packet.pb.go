// Code generated by protoc-gen-go.
// source: ap_packet.proto
// DO NOT EDIT!

package cheetah

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The message type defining the category of packets to be retrieved
type APMsgType int32

const (
	// Reserved. 0x0
	APMsgType_AP_MSG_TYPE_RESERVED APMsgType = 0
	// Traditional IEEE_TYPE_MANAGEMENT frames
	APMsgType_AP_MSG_TYPE_MGMT APMsgType = 1
	// Traditional IEEE_TYPE_CONTROL frames
	APMsgType_AP_MSG_TYPE_CTRL APMsgType = 2
	// Combination of IEEE_TYPE_DATA frames as well as other types (e.g, QOS_DATA)
	APMsgType_AP_MSG_TYPE_DATA APMsgType = 3
	// Cisco proprietary frames
	APMsgType_AP_MSG_TYPE_CISCO APMsgType = 4
)

var APMsgType_name = map[int32]string{
	0: "AP_MSG_TYPE_RESERVED",
	1: "AP_MSG_TYPE_MGMT",
	2: "AP_MSG_TYPE_CTRL",
	3: "AP_MSG_TYPE_DATA",
	4: "AP_MSG_TYPE_CISCO",
}
var APMsgType_value = map[string]int32{
	"AP_MSG_TYPE_RESERVED": 0,
	"AP_MSG_TYPE_MGMT":     1,
	"AP_MSG_TYPE_CTRL":     2,
	"AP_MSG_TYPE_DATA":     3,
	"AP_MSG_TYPE_CISCO":    4,
}

func (x APMsgType) String() string {
	return proto.EnumName(APMsgType_name, int32(x))
}
func (APMsgType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// The message subtype for the AP_MSG_TYPE_MGMT message type category
type APMgmtMsgSubtype int32

const (
	// Reserved. 0x0
	APMgmtMsgSubtype_AP_MGMT_MSG_SUBTYPE_RESERVED APMgmtMsgSubtype = 0
	// Association packets
	APMgmtMsgSubtype_AP_MGMT_MSG_SUBTYPE_ASSOC APMgmtMsgSubtype = 1
	// Authentication packets
	APMgmtMsgSubtype_AP_MGMT_MSG_SUBTYPE_AUTH APMgmtMsgSubtype = 2
	// Probe packets
	APMgmtMsgSubtype_AP_MGMT_MSG_SUBTYPE_PROBE APMgmtMsgSubtype = 4
	// All packets
	APMgmtMsgSubtype_AP_MGMT_MSG_SUBTYPE_ALL APMgmtMsgSubtype = 65535
)

var APMgmtMsgSubtype_name = map[int32]string{
	0:     "AP_MGMT_MSG_SUBTYPE_RESERVED",
	1:     "AP_MGMT_MSG_SUBTYPE_ASSOC",
	2:     "AP_MGMT_MSG_SUBTYPE_AUTH",
	4:     "AP_MGMT_MSG_SUBTYPE_PROBE",
	65535: "AP_MGMT_MSG_SUBTYPE_ALL",
}
var APMgmtMsgSubtype_value = map[string]int32{
	"AP_MGMT_MSG_SUBTYPE_RESERVED": 0,
	"AP_MGMT_MSG_SUBTYPE_ASSOC":    1,
	"AP_MGMT_MSG_SUBTYPE_AUTH":     2,
	"AP_MGMT_MSG_SUBTYPE_PROBE":    4,
	"AP_MGMT_MSG_SUBTYPE_ALL":      65535,
}

func (x APMgmtMsgSubtype) String() string {
	return proto.EnumName(APMgmtMsgSubtype_name, int32(x))
}
func (APMgmtMsgSubtype) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

// The message subtype for the AP_MSG_TYPE_CTRL message type category
type APCtrlMsgSubtype int32

const (
	// Reserved. 0x0
	APCtrlMsgSubtype_AP_CTRL_MSG_SUBTYPE_RESERVED APCtrlMsgSubtype = 0
	// All packets
	APCtrlMsgSubtype_AP_CTRL_MSG_SUBTYPE_ALL APCtrlMsgSubtype = 65535
)

var APCtrlMsgSubtype_name = map[int32]string{
	0:     "AP_CTRL_MSG_SUBTYPE_RESERVED",
	65535: "AP_CTRL_MSG_SUBTYPE_ALL",
}
var APCtrlMsgSubtype_value = map[string]int32{
	"AP_CTRL_MSG_SUBTYPE_RESERVED": 0,
	"AP_CTRL_MSG_SUBTYPE_ALL":      65535,
}

func (x APCtrlMsgSubtype) String() string {
	return proto.EnumName(APCtrlMsgSubtype_name, int32(x))
}
func (APCtrlMsgSubtype) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

// The message subtype for the AP_MSG_TYPE_DATA message type category
type APDataMsgSubtype int32

const (
	// Reserved. 0x0
	APDataMsgSubtype_AP_DATA_MSG_SUBTYPE_RESERVED APDataMsgSubtype = 0
	// ARP
	APDataMsgSubtype_AP_DATA_MSG_SUBTYPE_ARP APDataMsgSubtype = 1
	// DHCP
	APDataMsgSubtype_AP_DATA_MSG_SUBTYPE_DHCP APDataMsgSubtype = 2
	// EAP
	APDataMsgSubtype_AP_DATA_MSG_SUBTYPE_EAP APDataMsgSubtype = 4
	// ICMP
	APDataMsgSubtype_AP_DATA_MSG_SUBTYPE_ICMP APDataMsgSubtype = 8
	// All packets
	APDataMsgSubtype_AP_DATA_MSG_SUBTYPE_ALL APDataMsgSubtype = 65535
)

var APDataMsgSubtype_name = map[int32]string{
	0:     "AP_DATA_MSG_SUBTYPE_RESERVED",
	1:     "AP_DATA_MSG_SUBTYPE_ARP",
	2:     "AP_DATA_MSG_SUBTYPE_DHCP",
	4:     "AP_DATA_MSG_SUBTYPE_EAP",
	8:     "AP_DATA_MSG_SUBTYPE_ICMP",
	65535: "AP_DATA_MSG_SUBTYPE_ALL",
}
var APDataMsgSubtype_value = map[string]int32{
	"AP_DATA_MSG_SUBTYPE_RESERVED": 0,
	"AP_DATA_MSG_SUBTYPE_ARP":      1,
	"AP_DATA_MSG_SUBTYPE_DHCP":     2,
	"AP_DATA_MSG_SUBTYPE_EAP":      4,
	"AP_DATA_MSG_SUBTYPE_ICMP":     8,
	"AP_DATA_MSG_SUBTYPE_ALL":      65535,
}

func (x APDataMsgSubtype) String() string {
	return proto.EnumName(APDataMsgSubtype_name, int32(x))
}
func (APDataMsgSubtype) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

// The message subtype for the AP_MSG_TYPE_CISCO message type category
type APCiscoMsgSubtype int32

const (
	// Reserved. 0x0
	APCiscoMsgSubtype_AP_CISCO_MSG_SUBTYPE_RESERVED APCiscoMsgSubtype = 0
	// NDP
	APCiscoMsgSubtype_AP_CISCO_MSG_SUBTYPE_NDP APCiscoMsgSubtype = 1
	// All packets
	APCiscoMsgSubtype_AP_CISCO_MSG_SUBTYPE_ALL APCiscoMsgSubtype = 65535
)

var APCiscoMsgSubtype_name = map[int32]string{
	0:     "AP_CISCO_MSG_SUBTYPE_RESERVED",
	1:     "AP_CISCO_MSG_SUBTYPE_NDP",
	65535: "AP_CISCO_MSG_SUBTYPE_ALL",
}
var APCiscoMsgSubtype_value = map[string]int32{
	"AP_CISCO_MSG_SUBTYPE_RESERVED": 0,
	"AP_CISCO_MSG_SUBTYPE_NDP":      1,
	"AP_CISCO_MSG_SUBTYPE_ALL":      65535,
}

func (x APCiscoMsgSubtype) String() string {
	return proto.EnumName(APCiscoMsgSubtype_name, int32(x))
}
func (APCiscoMsgSubtype) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

// The packet type/subtype definition
type APPacketHdr struct {
	// AP Message Type, e.g.
	//    AP_MSG_TYPE_MGMT
	MsgType APMsgType `protobuf:"varint,1,opt,name=MsgType,enum=cheetah.APMsgType" json:"MsgType,omitempty"`
	// AP Message Subtype
	//
	// It could be a mask for the request message, e.g:
	//     AP_MGMT_MSG_SUBTYPE_ASSOC | AP_MGMT_MSG_SUBTYPE_AUTH, or just
	//     AP_DATA_MSG_SUBTYPE_ICMP
	//
	// It must be unary for the response message
	//
	//
	// Types that are valid to be assigned to Subtype:
	//	*APPacketHdr_Mgmt
	//	*APPacketHdr_Ctrl
	//	*APPacketHdr_Data
	//	*APPacketHdr_Cisco
	Subtype isAPPacketHdr_Subtype `protobuf_oneof:"Subtype"`
}

func (m *APPacketHdr) Reset()                    { *m = APPacketHdr{} }
func (m *APPacketHdr) String() string            { return proto.CompactTextString(m) }
func (*APPacketHdr) ProtoMessage()               {}
func (*APPacketHdr) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type isAPPacketHdr_Subtype interface {
	isAPPacketHdr_Subtype()
}

type APPacketHdr_Mgmt struct {
	Mgmt APMgmtMsgSubtype `protobuf:"varint,2,opt,name=mgmt,enum=cheetah.APMgmtMsgSubtype,oneof"`
}
type APPacketHdr_Ctrl struct {
	Ctrl APCtrlMsgSubtype `protobuf:"varint,3,opt,name=ctrl,enum=cheetah.APCtrlMsgSubtype,oneof"`
}
type APPacketHdr_Data struct {
	Data APDataMsgSubtype `protobuf:"varint,4,opt,name=data,enum=cheetah.APDataMsgSubtype,oneof"`
}
type APPacketHdr_Cisco struct {
	Cisco APCiscoMsgSubtype `protobuf:"varint,5,opt,name=cisco,enum=cheetah.APCiscoMsgSubtype,oneof"`
}

func (*APPacketHdr_Mgmt) isAPPacketHdr_Subtype()  {}
func (*APPacketHdr_Ctrl) isAPPacketHdr_Subtype()  {}
func (*APPacketHdr_Data) isAPPacketHdr_Subtype()  {}
func (*APPacketHdr_Cisco) isAPPacketHdr_Subtype() {}

func (m *APPacketHdr) GetSubtype() isAPPacketHdr_Subtype {
	if m != nil {
		return m.Subtype
	}
	return nil
}

func (m *APPacketHdr) GetMsgType() APMsgType {
	if m != nil {
		return m.MsgType
	}
	return APMsgType_AP_MSG_TYPE_RESERVED
}

func (m *APPacketHdr) GetMgmt() APMgmtMsgSubtype {
	if x, ok := m.GetSubtype().(*APPacketHdr_Mgmt); ok {
		return x.Mgmt
	}
	return APMgmtMsgSubtype_AP_MGMT_MSG_SUBTYPE_RESERVED
}

func (m *APPacketHdr) GetCtrl() APCtrlMsgSubtype {
	if x, ok := m.GetSubtype().(*APPacketHdr_Ctrl); ok {
		return x.Ctrl
	}
	return APCtrlMsgSubtype_AP_CTRL_MSG_SUBTYPE_RESERVED
}

func (m *APPacketHdr) GetData() APDataMsgSubtype {
	if x, ok := m.GetSubtype().(*APPacketHdr_Data); ok {
		return x.Data
	}
	return APDataMsgSubtype_AP_DATA_MSG_SUBTYPE_RESERVED
}

func (m *APPacketHdr) GetCisco() APCiscoMsgSubtype {
	if x, ok := m.GetSubtype().(*APPacketHdr_Cisco); ok {
		return x.Cisco
	}
	return APCiscoMsgSubtype_AP_CISCO_MSG_SUBTYPE_RESERVED
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*APPacketHdr) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _APPacketHdr_OneofMarshaler, _APPacketHdr_OneofUnmarshaler, _APPacketHdr_OneofSizer, []interface{}{
		(*APPacketHdr_Mgmt)(nil),
		(*APPacketHdr_Ctrl)(nil),
		(*APPacketHdr_Data)(nil),
		(*APPacketHdr_Cisco)(nil),
	}
}

func _APPacketHdr_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*APPacketHdr)
	// Subtype
	switch x := m.Subtype.(type) {
	case *APPacketHdr_Mgmt:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Mgmt))
	case *APPacketHdr_Ctrl:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Ctrl))
	case *APPacketHdr_Data:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Data))
	case *APPacketHdr_Cisco:
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Cisco))
	case nil:
	default:
		return fmt.Errorf("APPacketHdr.Subtype has unexpected type %T", x)
	}
	return nil
}

func _APPacketHdr_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*APPacketHdr)
	switch tag {
	case 2: // Subtype.mgmt
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Subtype = &APPacketHdr_Mgmt{APMgmtMsgSubtype(x)}
		return true, err
	case 3: // Subtype.ctrl
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Subtype = &APPacketHdr_Ctrl{APCtrlMsgSubtype(x)}
		return true, err
	case 4: // Subtype.data
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Subtype = &APPacketHdr_Data{APDataMsgSubtype(x)}
		return true, err
	case 5: // Subtype.cisco
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Subtype = &APPacketHdr_Cisco{APCiscoMsgSubtype(x)}
		return true, err
	default:
		return false, nil
	}
}

func _APPacketHdr_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*APPacketHdr)
	// Subtype
	switch x := m.Subtype.(type) {
	case *APPacketHdr_Mgmt:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Mgmt))
	case *APPacketHdr_Ctrl:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Ctrl))
	case *APPacketHdr_Data:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Data))
	case *APPacketHdr_Cisco:
		n += proto.SizeVarint(5<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Cisco))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Packet message request
type APPacketsMsg struct {
	// List of packet registrations
	PacketHdr []*APPacketHdr `protobuf:"bytes,1,rep,name=PacketHdr" json:"PacketHdr,omitempty"`
}

func (m *APPacketsMsg) Reset()                    { *m = APPacketsMsg{} }
func (m *APPacketsMsg) String() string            { return proto.CompactTextString(m) }
func (*APPacketsMsg) ProtoMessage()               {}
func (*APPacketsMsg) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *APPacketsMsg) GetPacketHdr() []*APPacketHdr {
	if m != nil {
		return m.PacketHdr
	}
	return nil
}

// Packet message response
type APPacketsMsgRsp struct {
	// Corresponding error code
	ErrStatus *APErrorStatus `protobuf:"bytes,1,opt,name=ErrStatus" json:"ErrStatus,omitempty"`
	// Type of pushed packet
	PacketHdr *APPacketHdr `protobuf:"bytes,2,opt,name=PacketHdr" json:"PacketHdr,omitempty"`
	// Length of pushed packet
	PacketLen uint32 `protobuf:"varint,3,opt,name=PacketLen" json:"PacketLen,omitempty"`
	// Buffer carrying the packet
	PacketBuf []byte `protobuf:"bytes,4,opt,name=PacketBuf,proto3" json:"PacketBuf,omitempty"`
}

func (m *APPacketsMsgRsp) Reset()                    { *m = APPacketsMsgRsp{} }
func (m *APPacketsMsgRsp) String() string            { return proto.CompactTextString(m) }
func (*APPacketsMsgRsp) ProtoMessage()               {}
func (*APPacketsMsgRsp) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *APPacketsMsgRsp) GetErrStatus() *APErrorStatus {
	if m != nil {
		return m.ErrStatus
	}
	return nil
}

func (m *APPacketsMsgRsp) GetPacketHdr() *APPacketHdr {
	if m != nil {
		return m.PacketHdr
	}
	return nil
}

func (m *APPacketsMsgRsp) GetPacketLen() uint32 {
	if m != nil {
		return m.PacketLen
	}
	return 0
}

func (m *APPacketsMsgRsp) GetPacketBuf() []byte {
	if m != nil {
		return m.PacketBuf
	}
	return nil
}

func init() {
	proto.RegisterType((*APPacketHdr)(nil), "cheetah.APPacketHdr")
	proto.RegisterType((*APPacketsMsg)(nil), "cheetah.APPacketsMsg")
	proto.RegisterType((*APPacketsMsgRsp)(nil), "cheetah.APPacketsMsgRsp")
	proto.RegisterEnum("cheetah.APMsgType", APMsgType_name, APMsgType_value)
	proto.RegisterEnum("cheetah.APMgmtMsgSubtype", APMgmtMsgSubtype_name, APMgmtMsgSubtype_value)
	proto.RegisterEnum("cheetah.APCtrlMsgSubtype", APCtrlMsgSubtype_name, APCtrlMsgSubtype_value)
	proto.RegisterEnum("cheetah.APDataMsgSubtype", APDataMsgSubtype_name, APDataMsgSubtype_value)
	proto.RegisterEnum("cheetah.APCiscoMsgSubtype", APCiscoMsgSubtype_name, APCiscoMsgSubtype_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for APPackets service

type APPacketsClient interface {
	// Registration RPC for packet types to be pushed
	APPacketsGet(ctx context.Context, in *APPacketsMsg, opts ...grpc.CallOption) (APPackets_APPacketsGetClient, error)
}

type aPPacketsClient struct {
	cc *grpc.ClientConn
}

func NewAPPacketsClient(cc *grpc.ClientConn) APPacketsClient {
	return &aPPacketsClient{cc}
}

func (c *aPPacketsClient) APPacketsGet(ctx context.Context, in *APPacketsMsg, opts ...grpc.CallOption) (APPackets_APPacketsGetClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_APPackets_serviceDesc.Streams[0], c.cc, "/cheetah.APPackets/APPacketsGet", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPPacketsAPPacketsGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type APPackets_APPacketsGetClient interface {
	Recv() (*APPacketsMsgRsp, error)
	grpc.ClientStream
}

type aPPacketsAPPacketsGetClient struct {
	grpc.ClientStream
}

func (x *aPPacketsAPPacketsGetClient) Recv() (*APPacketsMsgRsp, error) {
	m := new(APPacketsMsgRsp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for APPackets service

type APPacketsServer interface {
	// Registration RPC for packet types to be pushed
	APPacketsGet(*APPacketsMsg, APPackets_APPacketsGetServer) error
}

func RegisterAPPacketsServer(s *grpc.Server, srv APPacketsServer) {
	s.RegisterService(&_APPackets_serviceDesc, srv)
}

func _APPackets_APPacketsGet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(APPacketsMsg)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APPacketsServer).APPacketsGet(m, &aPPacketsAPPacketsGetServer{stream})
}

type APPackets_APPacketsGetServer interface {
	Send(*APPacketsMsgRsp) error
	grpc.ServerStream
}

type aPPacketsAPPacketsGetServer struct {
	grpc.ServerStream
}

func (x *aPPacketsAPPacketsGetServer) Send(m *APPacketsMsgRsp) error {
	return x.ServerStream.SendMsg(m)
}

var _APPackets_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cheetah.APPackets",
	HandlerType: (*APPacketsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "APPacketsGet",
			Handler:       _APPackets_APPacketsGet_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ap_packet.proto",
}

func init() { proto.RegisterFile("ap_packet.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xc1, 0x6e, 0x9b, 0x40,
	0x10, 0x86, 0x03, 0x76, 0xea, 0x7a, 0x9d, 0x36, 0xdb, 0x95, 0xdd, 0x10, 0xd7, 0xae, 0x5c, 0x9f,
	0x22, 0xab, 0x72, 0x2b, 0xda, 0x17, 0xc0, 0x18, 0xd9, 0x91, 0x4c, 0xbc, 0x5a, 0x48, 0xa5, 0x9e,
	0x10, 0x21, 0x94, 0x54, 0x0d, 0x01, 0xc1, 0xfa, 0x90, 0x4b, 0xdf, 0xa0, 0xcf, 0xd2, 0x63, 0x9f,
	0xa1, 0x4f, 0x95, 0x6a, 0xd7, 0x6b, 0x03, 0x06, 0xd4, 0xeb, 0xfc, 0xff, 0x37, 0xb3, 0x33, 0xcc,
	0x00, 0x4e, 0xdd, 0xd8, 0x89, 0x5d, 0xef, 0x87, 0x4f, 0xa7, 0x71, 0x12, 0xd1, 0x08, 0xb5, 0xbc,
	0x3b, 0xdf, 0xa7, 0xee, 0x5d, 0xbf, 0xe7, 0xc6, 0x8e, 0x17, 0x85, 0x61, 0xf4, 0xe0, 0xd0, 0xc7,
	0xd8, 0x4f, 0xb7, 0xfa, 0xf8, 0x97, 0x0c, 0x3a, 0x1a, 0xc6, 0x1c, 0x59, 0xde, 0x26, 0xe8, 0x3d,
	0x68, 0x99, 0x69, 0x60, 0x3f, 0xc6, 0xbe, 0x22, 0x8d, 0xa4, 0x8b, 0x97, 0x2a, 0x9a, 0x8a, 0x0c,
	0x53, 0x0d, 0x0b, 0x85, 0xec, 0x2c, 0xe8, 0x03, 0x68, 0x86, 0x41, 0x48, 0x15, 0x99, 0x5b, 0xcf,
	0xf3, 0xd6, 0x20, 0xa4, 0x66, 0x1a, 0x58, 0x9b, 0x1b, 0x56, 0x6d, 0x79, 0x44, 0xb8, 0x91, 0x01,
	0x1e, 0x4d, 0xee, 0x95, 0x46, 0x09, 0xd0, 0x69, 0x72, 0x5f, 0x04, 0x98, 0x91, 0x01, 0xb7, 0x2e,
	0x75, 0x95, 0x66, 0x09, 0x98, 0xbb, 0xd4, 0x2d, 0x02, 0xcc, 0x88, 0x54, 0x70, 0xec, 0x7d, 0x4f,
	0xbd, 0x48, 0x39, 0xe6, 0x44, 0x3f, 0x5f, 0x82, 0xc5, 0x0b, 0xc8, 0xd6, 0x3a, 0x6b, 0x83, 0x96,
	0x88, 0x8d, 0x67, 0xe0, 0x64, 0x37, 0x8e, 0xd4, 0x4c, 0x03, 0xa4, 0x82, 0xf6, 0x7e, 0x38, 0x8a,
	0x34, 0x6a, 0x5c, 0x74, 0xd4, 0x6e, 0x2e, 0xe5, 0x5e, 0x23, 0x99, 0x6d, 0xfc, 0x47, 0x02, 0xa7,
	0xf9, 0x24, 0x24, 0x8d, 0xd1, 0x67, 0xd0, 0x36, 0x92, 0xc4, 0xa2, 0x2e, 0xdd, 0xa4, 0x7c, 0xb2,
	0x1d, 0xf5, 0x75, 0x2e, 0x8f, 0x91, 0x24, 0x91, 0x50, 0x49, 0x66, 0x2c, 0x56, 0x97, 0x39, 0xf5,
	0xbf, 0xea, 0x68, 0xb0, 0x63, 0x56, 0xfe, 0x03, 0x9f, 0xf3, 0x0b, 0x92, 0x05, 0x32, 0x75, 0xb6,
	0xf9, 0xc6, 0x87, 0x7a, 0x42, 0xb2, 0xc0, 0xe4, 0x27, 0x68, 0xef, 0xbf, 0x32, 0x52, 0x40, 0x57,
	0xc3, 0x8e, 0x69, 0x2d, 0x1c, 0xfb, 0x2b, 0x36, 0x1c, 0x62, 0x58, 0x06, 0xf9, 0x62, 0xcc, 0xe1,
	0x11, 0xea, 0x02, 0x98, 0x57, 0xcc, 0x85, 0x69, 0x43, 0xe9, 0x30, 0xaa, 0xdb, 0x64, 0x05, 0xe5,
	0xc3, 0xe8, 0x5c, 0xb3, 0x35, 0xd8, 0x40, 0x3d, 0xf0, 0xaa, 0xe0, 0xbd, 0xb4, 0xf4, 0x35, 0x6c,
	0x4e, 0x7e, 0x4b, 0xcc, 0x5d, 0xdc, 0x1d, 0x34, 0x02, 0x03, 0xe6, 0x5d, 0x98, 0x36, 0x07, 0xac,
	0xeb, 0xd9, 0xe1, 0x7b, 0x86, 0xe0, 0xbc, 0xca, 0xa1, 0x59, 0xd6, 0x5a, 0x87, 0x12, 0x1a, 0x00,
	0xa5, 0x52, 0xbe, 0xb6, 0x97, 0x50, 0xae, 0x83, 0x31, 0x59, 0xcf, 0x0c, 0xd8, 0x44, 0x43, 0x70,
	0x56, 0x09, 0xaf, 0x56, 0xf0, 0xe9, 0xa9, 0x31, 0xb1, 0xd8, 0x83, 0x8b, 0xbb, 0x2b, 0x1e, 0xcc,
	0xfa, 0xaf, 0x7f, 0xf0, 0x59, 0x95, 0x63, 0x97, 0xf4, 0x2f, 0x1f, 0x43, 0x71, 0xc1, 0x45, 0x56,
	0x36, 0xbf, 0xba, 0xac, 0x6f, 0x78, 0xd6, 0x92, 0x43, 0x23, 0x78, 0x3f, 0x84, 0x92, 0x38, 0x5f,
	0xea, 0x18, 0xca, 0x75, 0xa8, 0xa1, 0x61, 0xd8, 0xac, 0x43, 0x2f, 0x75, 0x13, 0xc3, 0xe7, 0xa2,
	0x97, 0x72, 0x55, 0xd1, 0x0b, 0x65, 0x5f, 0xfa, 0xe0, 0xf2, 0xd0, 0x3b, 0x30, 0x64, 0xfd, 0xb3,
	0xaf, 0x5e, 0xd7, 0xcc, 0xb6, 0x68, 0xd9, 0x72, 0x35, 0x67, 0xdd, 0xbc, 0xad, 0x51, 0x45, 0x55,
	0xf5, 0x8a, 0x2d, 0xb2, 0xb8, 0x40, 0xa4, 0xe5, 0x6e, 0x7a, 0xe1, 0x53, 0xd4, 0x2b, 0x9d, 0x10,
	0xbb, 0xd2, 0xbe, 0x52, 0x19, 0x26, 0x69, 0xfc, 0x51, 0xba, 0x79, 0xc6, 0xff, 0x96, 0x9f, 0xfe,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x48, 0x4e, 0x04, 0x60, 0x05, 0x00, 0x00,
}
