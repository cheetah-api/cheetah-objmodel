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

// Packet Get Notifications message
type APPacketsGetNotifMsg struct {
}

func (m *APPacketsGetNotifMsg) Reset()                    { *m = APPacketsGetNotifMsg{} }
func (m *APPacketsGetNotifMsg) String() string            { return proto.CompactTextString(m) }
func (*APPacketsGetNotifMsg) ProtoMessage()               {}
func (*APPacketsGetNotifMsg) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// Packet Registration messages
type APPacketsRegMsg struct {
	// Registration Operation
	Oper APRegOp `protobuf:"varint,1,opt,name=Oper,enum=cheetah.APRegOp" json:"Oper,omitempty"`
	// List of packet registrations
	PacketHdr []*APPacketHdr `protobuf:"bytes,2,rep,name=PacketHdr" json:"PacketHdr,omitempty"`
}

func (m *APPacketsRegMsg) Reset()                    { *m = APPacketsRegMsg{} }
func (m *APPacketsRegMsg) String() string            { return proto.CompactTextString(m) }
func (*APPacketsRegMsg) ProtoMessage()               {}
func (*APPacketsRegMsg) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *APPacketsRegMsg) GetOper() APRegOp {
	if m != nil {
		return m.Oper
	}
	return APRegOp_AP_REGOP_RESERVED
}

func (m *APPacketsRegMsg) GetPacketHdr() []*APPacketHdr {
	if m != nil {
		return m.PacketHdr
	}
	return nil
}

// Packet Registration message response
type APPacketsRegMsgRsp struct {
	// Summary result of the bulk operation (refer to enum APErrorStatus)
	//
	// In general, the ErrStatus consists of:
	// 1. AP_SUCCESS: signifies that the entire bulk operation was successful.
	//         In this case, the Results list is empty.
	// 2. AP_EINVAL: signifies that the operation failed for one or more entries
	//         In this case, Results holds the result for each individual entry
	//         that was in error
	ErrStatus *APErrorStatus `protobuf:"bytes,1,opt,name=ErrStatus" json:"ErrStatus,omitempty"`
	// In case of errors, this field contains all the entries that were in error
	Results *APPacketsRegMsg `protobuf:"bytes,2,opt,name=Results" json:"Results,omitempty"`
}

func (m *APPacketsRegMsgRsp) Reset()                    { *m = APPacketsRegMsgRsp{} }
func (m *APPacketsRegMsgRsp) String() string            { return proto.CompactTextString(m) }
func (*APPacketsRegMsgRsp) ProtoMessage()               {}
func (*APPacketsRegMsgRsp) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *APPacketsRegMsgRsp) GetErrStatus() *APErrorStatus {
	if m != nil {
		return m.ErrStatus
	}
	return nil
}

func (m *APPacketsRegMsgRsp) GetResults() *APPacketsRegMsg {
	if m != nil {
		return m.Results
	}
	return nil
}

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
func (*APPacketHdr) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

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
func (*APPacketsMsgRsp) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

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
	proto.RegisterType((*APPacketsGetNotifMsg)(nil), "cheetah.APPacketsGetNotifMsg")
	proto.RegisterType((*APPacketsRegMsg)(nil), "cheetah.APPacketsRegMsg")
	proto.RegisterType((*APPacketsRegMsgRsp)(nil), "cheetah.APPacketsRegMsgRsp")
	proto.RegisterType((*APPacketHdr)(nil), "cheetah.APPacketHdr")
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
	// APPacketsRegMsg.Oper = AP_REGOP_REGISTER
	//     Packet registration: Sends a list of Packet registration messages
	//     and expects a list of registration responses.
	//
	// APPacketsRegMsg.Oper = AP_REGOP_UNREGISTER
	//     Packet unregistration: Sends a list of Packet unregistration messages
	//     and expects a list of unregistration responses.
	//
	APPacketsRegOp(ctx context.Context, in *APPacketsRegMsg, opts ...grpc.CallOption) (*APPacketsRegMsgRsp, error)
	// This call is used to get a stream of packet notifications matching the
	// set of registrations performed with APPacketsRegOp().
	// The caller must maintain the GRPC channel as long as
	// there is interest in packet notifications. Only sessions that were
	// created through this API will be notified to caller.
	APPacketsInitNotif(ctx context.Context, in *APPacketsGetNotifMsg, opts ...grpc.CallOption) (APPackets_APPacketsInitNotifClient, error)
}

type aPPacketsClient struct {
	cc *grpc.ClientConn
}

func NewAPPacketsClient(cc *grpc.ClientConn) APPacketsClient {
	return &aPPacketsClient{cc}
}

func (c *aPPacketsClient) APPacketsRegOp(ctx context.Context, in *APPacketsRegMsg, opts ...grpc.CallOption) (*APPacketsRegMsgRsp, error) {
	out := new(APPacketsRegMsgRsp)
	err := grpc.Invoke(ctx, "/cheetah.APPackets/APPacketsRegOp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPPacketsClient) APPacketsInitNotif(ctx context.Context, in *APPacketsGetNotifMsg, opts ...grpc.CallOption) (APPackets_APPacketsInitNotifClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_APPackets_serviceDesc.Streams[0], c.cc, "/cheetah.APPackets/APPacketsInitNotif", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPPacketsAPPacketsInitNotifClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type APPackets_APPacketsInitNotifClient interface {
	Recv() (*APPacketsMsgRsp, error)
	grpc.ClientStream
}

type aPPacketsAPPacketsInitNotifClient struct {
	grpc.ClientStream
}

func (x *aPPacketsAPPacketsInitNotifClient) Recv() (*APPacketsMsgRsp, error) {
	m := new(APPacketsMsgRsp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for APPackets service

type APPacketsServer interface {
	// APPacketsRegMsg.Oper = AP_REGOP_REGISTER
	//     Packet registration: Sends a list of Packet registration messages
	//     and expects a list of registration responses.
	//
	// APPacketsRegMsg.Oper = AP_REGOP_UNREGISTER
	//     Packet unregistration: Sends a list of Packet unregistration messages
	//     and expects a list of unregistration responses.
	//
	APPacketsRegOp(context.Context, *APPacketsRegMsg) (*APPacketsRegMsgRsp, error)
	// This call is used to get a stream of packet notifications matching the
	// set of registrations performed with APPacketsRegOp().
	// The caller must maintain the GRPC channel as long as
	// there is interest in packet notifications. Only sessions that were
	// created through this API will be notified to caller.
	APPacketsInitNotif(*APPacketsGetNotifMsg, APPackets_APPacketsInitNotifServer) error
}

func RegisterAPPacketsServer(s *grpc.Server, srv APPacketsServer) {
	s.RegisterService(&_APPackets_serviceDesc, srv)
}

func _APPackets_APPacketsRegOp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(APPacketsRegMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APPacketsServer).APPacketsRegOp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheetah.APPackets/APPacketsRegOp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APPacketsServer).APPacketsRegOp(ctx, req.(*APPacketsRegMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _APPackets_APPacketsInitNotif_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(APPacketsGetNotifMsg)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APPacketsServer).APPacketsInitNotif(m, &aPPacketsAPPacketsInitNotifServer{stream})
}

type APPackets_APPacketsInitNotifServer interface {
	Send(*APPacketsMsgRsp) error
	grpc.ServerStream
}

type aPPacketsAPPacketsInitNotifServer struct {
	grpc.ServerStream
}

func (x *aPPacketsAPPacketsInitNotifServer) Send(m *APPacketsMsgRsp) error {
	return x.ServerStream.SendMsg(m)
}

var _APPackets_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cheetah.APPackets",
	HandlerType: (*APPacketsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "APPacketsRegOp",
			Handler:    _APPackets_APPacketsRegOp_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "APPacketsInitNotif",
			Handler:       _APPackets_APPacketsInitNotif_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ap_packet.proto",
}

func init() { proto.RegisterFile("ap_packet.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0x9b, 0x4c,
	0x14, 0x0d, 0x98, 0x7c, 0xfe, 0x3c, 0x69, 0x93, 0xe9, 0x28, 0x3f, 0xc4, 0x89, 0x2b, 0xd7, 0xea,
	0x22, 0xb2, 0x2a, 0xb7, 0xa2, 0x7d, 0x01, 0x8c, 0x91, 0x6d, 0xc9, 0x04, 0x34, 0x90, 0x4a, 0x5d,
	0x21, 0x42, 0x08, 0xb1, 0x62, 0x1b, 0xc4, 0x8c, 0x17, 0xd9, 0xe4, 0x0d, 0xfa, 0x18, 0x5d, 0x77,
	0xd9, 0x67, 0xe8, 0x53, 0xa5, 0x9a, 0x01, 0xdb, 0x60, 0x8c, 0x54, 0xa9, 0xdb, 0x7b, 0xce, 0xb9,
	0x67, 0xee, 0x0f, 0x17, 0x70, 0xe4, 0xc5, 0x6e, 0xec, 0xf9, 0x8f, 0x01, 0xed, 0xc5, 0x49, 0x44,
	0x23, 0x54, 0xf7, 0x1f, 0x82, 0x80, 0x7a, 0x0f, 0xcd, 0x13, 0x2f, 0x76, 0xfd, 0x68, 0x3e, 0x8f,
	0x16, 0x2e, 0x7d, 0x8a, 0x03, 0x92, 0xe2, 0x9d, 0x53, 0x70, 0xac, 0x5a, 0x16, 0x57, 0x90, 0x61,
	0x40, 0xaf, 0x23, 0x3a, 0xbd, 0x37, 0x48, 0xd8, 0x79, 0x04, 0x47, 0xeb, 0x38, 0x0e, 0x42, 0x83,
	0x84, 0xe8, 0x3d, 0x90, 0xcc, 0x38, 0x48, 0x64, 0xa1, 0x2d, 0x5c, 0x1d, 0x2a, 0xb0, 0x97, 0x65,
	0xee, 0xa9, 0x16, 0x0e, 0x42, 0x33, 0xc6, 0x1c, 0x45, 0x0a, 0x68, 0xa4, 0xb2, 0xd1, 0x5d, 0x22,
	0x8b, 0xed, 0xda, 0xd5, 0x81, 0x72, 0x9c, 0xa3, 0xae, 0x31, 0xbc, 0xa1, 0x75, 0x9e, 0x01, 0xda,
	0x32, 0xc3, 0x24, 0x46, 0x5f, 0x40, 0x43, 0x4f, 0x12, 0x9b, 0x7a, 0x74, 0x49, 0xb8, 0xe9, 0x81,
	0x72, 0x9a, 0xcb, 0xa4, 0x27, 0x49, 0x94, 0xa1, 0x78, 0x43, 0x44, 0x0a, 0xa8, 0xe3, 0x80, 0x2c,
	0x67, 0x94, 0xc8, 0x22, 0xd7, 0xc8, 0x25, 0xf7, 0x95, 0xc7, 0x8a, 0xd8, 0xf9, 0x2e, 0x82, 0x83,
	0xdc, 0xd3, 0xd0, 0x07, 0x50, 0x37, 0x48, 0xe8, 0x3c, 0xc5, 0x41, 0x56, 0x2c, 0xca, 0xe5, 0xc8,
	0x10, 0xbc, 0xa2, 0xa0, 0x8f, 0x40, 0x9a, 0x87, 0x73, 0xca, 0xed, 0x0e, 0x95, 0xf3, 0x3c, 0x35,
	0x9c, 0x53, 0x83, 0x84, 0xf6, 0xf2, 0x96, 0xb5, 0x7c, 0xb4, 0x87, 0x39, 0x91, 0x09, 0x7c, 0x9a,
	0xcc, 0xe4, 0x5a, 0x49, 0xa0, 0xd1, 0x64, 0x56, 0x14, 0x30, 0x22, 0x13, 0xdc, 0x79, 0xd4, 0x93,
	0xa5, 0x92, 0x60, 0xe0, 0x51, 0xaf, 0x28, 0x60, 0x44, 0xa4, 0x80, 0x7d, 0x7f, 0x4a, 0xfc, 0x48,
	0xde, 0xe7, 0x8a, 0x66, 0xde, 0x82, 0xc5, 0x0b, 0x92, 0x94, 0xda, 0x6f, 0x80, 0x7a, 0x16, 0xeb,
	0xfc, 0x12, 0x72, 0xd3, 0xff, 0xc7, 0x69, 0x14, 0xb6, 0x41, 0xf8, 0x8b, 0x6d, 0x40, 0x97, 0x2b,
	0xcd, 0x24, 0x58, 0xf0, 0x1e, 0xbd, 0xc6, 0x9b, 0xc0, 0x06, 0xed, 0x2f, 0xef, 0x79, 0x43, 0x5e,
	0xe1, 0x4d, 0xa0, 0xfb, 0x0c, 0x1a, 0xeb, 0x09, 0x21, 0x99, 0xed, 0xb6, 0x6b, 0xd8, 0x43, 0xd7,
	0xf9, 0x66, 0xe9, 0x2e, 0xd6, 0x6d, 0x1d, 0x7f, 0xd5, 0x07, 0x70, 0x0f, 0x1d, 0x03, 0x98, 0x47,
	0x8c, 0xa1, 0xe1, 0x40, 0x61, 0x3b, 0xaa, 0x39, 0x78, 0x02, 0xc5, 0xed, 0xe8, 0x40, 0x75, 0x54,
	0x58, 0x43, 0x27, 0xe0, 0x4d, 0x81, 0x3b, 0xb6, 0x35, 0x13, 0x4a, 0xdd, 0x9f, 0x02, 0x63, 0x17,
	0xe7, 0x8e, 0xda, 0xe0, 0x92, 0x71, 0x87, 0x86, 0xc3, 0x05, 0xf6, 0x4d, 0x7f, 0xfb, 0x3d, 0x2d,
	0x70, 0xbe, 0x8b, 0xa1, 0xda, 0xb6, 0xa9, 0x41, 0x01, 0x5d, 0x02, 0x79, 0x27, 0x7c, 0xe3, 0x8c,
	0xa0, 0x58, 0x25, 0xb6, 0xb0, 0xd9, 0xd7, 0xa1, 0x84, 0x5a, 0xe0, 0x6c, 0xa7, 0x78, 0x32, 0x81,
	0x2f, 0x2f, 0xb5, 0xae, 0xcd, 0x1e, 0x5c, 0xdc, 0xbb, 0xec, 0xc1, 0xac, 0xfe, 0xea, 0x07, 0x9f,
	0xed, 0x62, 0xac, 0x92, 0xfe, 0xe6, 0x6d, 0x28, 0x2e, 0x67, 0x96, 0x95, 0xf5, 0xaf, 0x2a, 0xeb,
	0x05, 0xcf, 0x5a, 0x62, 0xa8, 0xd8, 0x5a, 0x37, 0xa1, 0x04, 0x0e, 0x46, 0x9a, 0x05, 0xc5, 0x2a,
	0xa9, 0xae, 0x5a, 0x50, 0xaa, 0x92, 0x8e, 0x35, 0xc3, 0x82, 0xff, 0x67, 0xb5, 0x94, 0x5d, 0xb3,
	0x5a, 0x28, 0x9b, 0xf4, 0xd6, 0x57, 0x83, 0xde, 0x81, 0x16, 0xab, 0x9f, 0x4d, 0xbd, 0xaa, 0x98,
	0xd4, 0xb4, 0x4c, 0xb9, 0x1e, 0xb0, 0x6a, 0xde, 0x56, 0xa0, 0x99, 0xab, 0xf2, 0x43, 0x60, 0x9b,
	0x9c, 0x7d, 0x82, 0x68, 0x08, 0x0e, 0xf3, 0xc7, 0xcb, 0x8c, 0x51, 0xe5, 0x55, 0x6b, 0x5e, 0x54,
	0xde, 0x3b, 0x12, 0x23, 0x33, 0x77, 0x69, 0xc7, 0x8b, 0x69, 0x7a, 0xef, 0x51, 0xab, 0x2c, 0xc9,
	0xfd, 0x0b, 0x9a, 0x3b, 0xbc, 0xd2, 0x74, 0x9f, 0x84, 0xdb, 0xff, 0xf8, 0x6f, 0xe4, 0xf3, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0xa6, 0x1d, 0x2a, 0x79, 0x06, 0x00, 0x00,
}
