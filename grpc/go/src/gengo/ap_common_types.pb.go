// Code generated by protoc-gen-go.
// source: ap_common_types.proto
// DO NOT EDIT!

/*
Package access_point is a generated protocol buffer package.

It is generated from these files:
	ap_common_types.proto
	ap_global.proto
	ap_packet.proto
	ap_stats.proto
	ap_version.proto

It has these top-level messages:
	APErrorStatus
	APRadio
	APSsid
	APInitMsg
	APInitMsgRsp
	APGlobalNotif
	APGlobalsGetMsg
	APGlobalsGetMsgRsp
	APPacketHdr
	APPacketsMsg
	APPacketsMsgRsp
	APStatsRequest
	APStatsMsg
	APSystemStatsMsgRsp
	MemInfo
	SlabInfo
	APMemoryStatsMsgRsp
	APDNSStatsMsgRsp
	IPv4Route
	APRoutingStatsMsgRsp
	MulticastCounter
	WLAN
	WLANEntry
	APWLANStatsMsgRsp
	RadioUtilization
	RadioCounters
	DfsState
	RadioEntry
	APRadioStatsMsgRsp
	APClientEntry
	APClientStatsMsgRsp
	APInterfaceEntry
	APInterfaceStatsMsgRsp
	APStatsMsgRsp
*/
package access_point

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Registration Operations.
type APRegOp int32

const (
	// Reserved. 0x0
	APRegOp_AP_REGOP_RESERVED APRegOp = 0
	// Register. 0x1
	APRegOp_AP_REGOP_REGISTER APRegOp = 1
	// Un-Register. 0x2
	APRegOp_AP_REGOP_UNREGISTER APRegOp = 2
	// End Of File. 0x3
	APRegOp_AP_REGOP_EOF APRegOp = 3
)

var APRegOp_name = map[int32]string{
	0: "AP_REGOP_RESERVED",
	1: "AP_REGOP_REGISTER",
	2: "AP_REGOP_UNREGISTER",
	3: "AP_REGOP_EOF",
}
var APRegOp_value = map[string]int32{
	"AP_REGOP_RESERVED":   0,
	"AP_REGOP_REGISTER":   1,
	"AP_REGOP_UNREGISTER": 2,
	"AP_REGOP_EOF":        3,
}

func (x APRegOp) String() string {
	return proto.EnumName(APRegOp_name, int32(x))
}
func (APRegOp) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Object Operations.
type APObjectOp int32

const (
	// Reserved. 0x0
	APObjectOp_AP_OBJOP_RESERVED APObjectOp = 0
	// Add. 0x1
	APObjectOp_AP_OBJOP_ADD APObjectOp = 1
	// Update. 0x2
	APObjectOp_AP_OBJOP_UPDATE APObjectOp = 2
	// Delete. 0x3
	APObjectOp_AP_OBJOP_DELETE APObjectOp = 3
)

var APObjectOp_name = map[int32]string{
	0: "AP_OBJOP_RESERVED",
	1: "AP_OBJOP_ADD",
	2: "AP_OBJOP_UPDATE",
	3: "AP_OBJOP_DELETE",
}
var APObjectOp_value = map[string]int32{
	"AP_OBJOP_RESERVED": 0,
	"AP_OBJOP_ADD":      1,
	"AP_OBJOP_UPDATE":   2,
	"AP_OBJOP_DELETE":   3,
}

func (x APObjectOp) String() string {
	return proto.EnumName(APObjectOp_name, int32(x))
}
func (APObjectOp) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Notification Operations.
type APNotifOp int32

const (
	// Reserved. 0x0
	APNotifOp_AP_NOTIFOP_RESERVED APNotifOp = 0
	// Enable. 0x1
	APNotifOp_AP_NOTIFOP_ENABLE APNotifOp = 1
	// Disable. 0x2
	APNotifOp_AP_NOTIFOP_DISABLE APNotifOp = 2
)

var APNotifOp_name = map[int32]string{
	0: "AP_NOTIFOP_RESERVED",
	1: "AP_NOTIFOP_ENABLE",
	2: "AP_NOTIFOP_DISABLE",
}
var APNotifOp_value = map[string]int32{
	"AP_NOTIFOP_RESERVED": 0,
	"AP_NOTIFOP_ENABLE":   1,
	"AP_NOTIFOP_DISABLE":  2,
}

func (x APNotifOp) String() string {
	return proto.EnumName(APNotifOp_name, int32(x))
}
func (APNotifOp) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type APErrorStatus_APErrno int32

const (
	// Success, no errors detected. 0x0.
	APErrorStatus_AP_SUCCESS APErrorStatus_APErrno = 0
	// Client is not connected.
	// The client is expected to remain connected after init and version
	// validation, RPC calls can fail with this error code otherwise.
	// Refer to RPC APGlobalInitNotif. 0x1
	APErrorStatus_AP_NOT_CONNECTED APErrorStatus_APErrno = 1
	// Operation must be retried. 0x2
	APErrorStatus_AP_EAGAIN APErrorStatus_APErrno = 2
	// One or more components does not have sufficient memory. 0x3
	APErrorStatus_AP_ENOMEM APErrorStatus_APErrno = 3
	// Too many outstanding requests. 0x4
	APErrorStatus_AP_EBUSY APErrorStatus_APErrno = 4
	// One or more arguments are invalid. 0x5
	APErrorStatus_AP_EINVAL APErrorStatus_APErrno = 5
	// Unsupported version. 0x6
	APErrorStatus_AP_UNSUPPORTED_VER APErrorStatus_APErrno = 6
	// Not Available. 0x7
	APErrorStatus_AP_NOT_AVAILABLE APErrorStatus_APErrno = 7
	// Stream mode not supported. 0x8
	APErrorStatus_AP_STREAM_NOT_SUPPORTED APErrorStatus_APErrno = 8
	// Operation not supported. 0x9
	APErrorStatus_AP_ENOTSUP APErrorStatus_APErrno = 9
	// One or more objects is errored:
	// Each object must be individually examined. 0xa
	APErrorStatus_AP_SOME_ERR APErrorStatus_APErrno = 10
	// Operation Timed out.
	// The result of the operation is undeterministic (success or fail). 0xb
	APErrorStatus_AP_TIMEOUT APErrorStatus_APErrno = 11
	// Due to some event, the client will no longer receive notification
	// events on this channel. 0xc
	// Such events include:
	// - Notification Session was hijacked by another client.
	APErrorStatus_AP_NOTIF_TERM APErrorStatus_APErrno = 12
	// Offset for INIT errors. 0x500
	APErrorStatus_AP_INIT_START_OFFSET APErrorStatus_APErrno = 1280
	// Success, no errors detected - clear state.
	// This error is returned on the first-ever initialization, or,
	// when a fatal event has occured and all previous state was lost. 0x501
	APErrorStatus_AP_INIT_STATE_CLEAR APErrorStatus_APErrno = 1281
	// Success, no errors detected - previous state fully recovered.
	// This error is returned on a client re-initialization with
	// successful recovery of state. 0x502
	APErrorStatus_AP_INIT_STATE_READY APErrorStatus_APErrno = 1282
	// Server software incompatible with client software version. 0x503
	APErrorStatus_AP_INIT_UNSUPPORTED_VER APErrorStatus_APErrno = 1283
	// Initialization request received while server is not ready. 0x504
	APErrorStatus_AP_INIT_SERVER_NOT_INITIALIZED APErrorStatus_APErrno = 1284
	// Server operational mode change from stream to non-stream
	// or vice-versa failed. 0x505
	APErrorStatus_AP_INIT_SERVER_MODE_CHANGE_FAILED APErrorStatus_APErrno = 1285
)

var APErrorStatus_APErrno_name = map[int32]string{
	0:    "AP_SUCCESS",
	1:    "AP_NOT_CONNECTED",
	2:    "AP_EAGAIN",
	3:    "AP_ENOMEM",
	4:    "AP_EBUSY",
	5:    "AP_EINVAL",
	6:    "AP_UNSUPPORTED_VER",
	7:    "AP_NOT_AVAILABLE",
	8:    "AP_STREAM_NOT_SUPPORTED",
	9:    "AP_ENOTSUP",
	10:   "AP_SOME_ERR",
	11:   "AP_TIMEOUT",
	12:   "AP_NOTIF_TERM",
	1280: "AP_INIT_START_OFFSET",
	1281: "AP_INIT_STATE_CLEAR",
	1282: "AP_INIT_STATE_READY",
	1283: "AP_INIT_UNSUPPORTED_VER",
	1284: "AP_INIT_SERVER_NOT_INITIALIZED",
	1285: "AP_INIT_SERVER_MODE_CHANGE_FAILED",
}
var APErrorStatus_APErrno_value = map[string]int32{
	"AP_SUCCESS":                        0,
	"AP_NOT_CONNECTED":                  1,
	"AP_EAGAIN":                         2,
	"AP_ENOMEM":                         3,
	"AP_EBUSY":                          4,
	"AP_EINVAL":                         5,
	"AP_UNSUPPORTED_VER":                6,
	"AP_NOT_AVAILABLE":                  7,
	"AP_STREAM_NOT_SUPPORTED":           8,
	"AP_ENOTSUP":                        9,
	"AP_SOME_ERR":                       10,
	"AP_TIMEOUT":                        11,
	"AP_NOTIF_TERM":                     12,
	"AP_INIT_START_OFFSET":              1280,
	"AP_INIT_STATE_CLEAR":               1281,
	"AP_INIT_STATE_READY":               1282,
	"AP_INIT_UNSUPPORTED_VER":           1283,
	"AP_INIT_SERVER_NOT_INITIALIZED":    1284,
	"AP_INIT_SERVER_MODE_CHANGE_FAILED": 1285,
}

func (x APErrorStatus_APErrno) String() string {
	return proto.EnumName(APErrorStatus_APErrno_name, int32(x))
}
func (APErrorStatus_APErrno) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Status codes, including errors and success codes.
// All errors are defined below.
type APErrorStatus struct {
	Status APErrorStatus_APErrno `protobuf:"varint,1,opt,name=Status,enum=access_point.APErrorStatus_APErrno" json:"Status,omitempty"`
}

func (m *APErrorStatus) Reset()                    { *m = APErrorStatus{} }
func (m *APErrorStatus) String() string            { return proto.CompactTextString(m) }
func (*APErrorStatus) ProtoMessage()               {}
func (*APErrorStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *APErrorStatus) GetStatus() APErrorStatus_APErrno {
	if m != nil {
		return m.Status
	}
	return APErrorStatus_AP_SUCCESS
}

// Radio info.
type APRadio struct {
	// One of radio name or handle
	//
	// Types that are valid to be assigned to Radio:
	//	*APRadio_Name
	//	*APRadio_Handle
	Radio isAPRadio_Radio `protobuf_oneof:"Radio"`
}

func (m *APRadio) Reset()                    { *m = APRadio{} }
func (m *APRadio) String() string            { return proto.CompactTextString(m) }
func (*APRadio) ProtoMessage()               {}
func (*APRadio) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isAPRadio_Radio interface {
	isAPRadio_Radio()
}

type APRadio_Name struct {
	Name string `protobuf:"bytes,1,opt,name=Name,oneof"`
}
type APRadio_Handle struct {
	Handle uint32 `protobuf:"varint,2,opt,name=Handle,oneof"`
}

func (*APRadio_Name) isAPRadio_Radio()   {}
func (*APRadio_Handle) isAPRadio_Radio() {}

func (m *APRadio) GetRadio() isAPRadio_Radio {
	if m != nil {
		return m.Radio
	}
	return nil
}

func (m *APRadio) GetName() string {
	if x, ok := m.GetRadio().(*APRadio_Name); ok {
		return x.Name
	}
	return ""
}

func (m *APRadio) GetHandle() uint32 {
	if x, ok := m.GetRadio().(*APRadio_Handle); ok {
		return x.Handle
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*APRadio) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _APRadio_OneofMarshaler, _APRadio_OneofUnmarshaler, _APRadio_OneofSizer, []interface{}{
		(*APRadio_Name)(nil),
		(*APRadio_Handle)(nil),
	}
}

func _APRadio_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*APRadio)
	// Radio
	switch x := m.Radio.(type) {
	case *APRadio_Name:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Name)
	case *APRadio_Handle:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Handle))
	case nil:
	default:
		return fmt.Errorf("APRadio.Radio has unexpected type %T", x)
	}
	return nil
}

func _APRadio_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*APRadio)
	switch tag {
	case 1: // Radio.Name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Radio = &APRadio_Name{x}
		return true, err
	case 2: // Radio.Handle
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Radio = &APRadio_Handle{uint32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _APRadio_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*APRadio)
	// Radio
	switch x := m.Radio.(type) {
	case *APRadio_Name:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Name)))
		n += len(x.Name)
	case *APRadio_Handle:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Handle))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Ssid info.
type APSsid struct {
	// One of ssid name or handle
	//
	// Types that are valid to be assigned to Ssid:
	//	*APSsid_Name
	//	*APSsid_Handle
	Ssid isAPSsid_Ssid `protobuf_oneof:"Ssid"`
}

func (m *APSsid) Reset()                    { *m = APSsid{} }
func (m *APSsid) String() string            { return proto.CompactTextString(m) }
func (*APSsid) ProtoMessage()               {}
func (*APSsid) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isAPSsid_Ssid interface {
	isAPSsid_Ssid()
}

type APSsid_Name struct {
	Name string `protobuf:"bytes,1,opt,name=Name,oneof"`
}
type APSsid_Handle struct {
	Handle uint32 `protobuf:"varint,2,opt,name=Handle,oneof"`
}

func (*APSsid_Name) isAPSsid_Ssid()   {}
func (*APSsid_Handle) isAPSsid_Ssid() {}

func (m *APSsid) GetSsid() isAPSsid_Ssid {
	if m != nil {
		return m.Ssid
	}
	return nil
}

func (m *APSsid) GetName() string {
	if x, ok := m.GetSsid().(*APSsid_Name); ok {
		return x.Name
	}
	return ""
}

func (m *APSsid) GetHandle() uint32 {
	if x, ok := m.GetSsid().(*APSsid_Handle); ok {
		return x.Handle
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*APSsid) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _APSsid_OneofMarshaler, _APSsid_OneofUnmarshaler, _APSsid_OneofSizer, []interface{}{
		(*APSsid_Name)(nil),
		(*APSsid_Handle)(nil),
	}
}

func _APSsid_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*APSsid)
	// Ssid
	switch x := m.Ssid.(type) {
	case *APSsid_Name:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Name)
	case *APSsid_Handle:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Handle))
	case nil:
	default:
		return fmt.Errorf("APSsid.Ssid has unexpected type %T", x)
	}
	return nil
}

func _APSsid_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*APSsid)
	switch tag {
	case 1: // Ssid.Name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Ssid = &APSsid_Name{x}
		return true, err
	case 2: // Ssid.Handle
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Ssid = &APSsid_Handle{uint32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _APSsid_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*APSsid)
	// Ssid
	switch x := m.Ssid.(type) {
	case *APSsid_Name:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Name)))
		n += len(x.Name)
	case *APSsid_Handle:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Handle))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*APErrorStatus)(nil), "access_point.APErrorStatus")
	proto.RegisterType((*APRadio)(nil), "access_point.APRadio")
	proto.RegisterType((*APSsid)(nil), "access_point.APSsid")
	proto.RegisterEnum("access_point.APRegOp", APRegOp_name, APRegOp_value)
	proto.RegisterEnum("access_point.APObjectOp", APObjectOp_name, APObjectOp_value)
	proto.RegisterEnum("access_point.APNotifOp", APNotifOp_name, APNotifOp_value)
	proto.RegisterEnum("access_point.APErrorStatus_APErrno", APErrorStatus_APErrno_name, APErrorStatus_APErrno_value)
}

func init() { proto.RegisterFile("ap_common_types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4f, 0x4f, 0x1b, 0x3f,
	0x10, 0xcd, 0x3f, 0x02, 0x0c, 0x01, 0x06, 0x03, 0x3f, 0xf2, 0x53, 0xab, 0x8a, 0x06, 0xa9, 0x42,
	0x1c, 0x38, 0xb4, 0xc7, 0x5e, 0xea, 0x64, 0x27, 0xc9, 0x56, 0xbb, 0xf6, 0xca, 0xf6, 0x46, 0xa2,
	0x17, 0x2b, 0x84, 0xb4, 0x4a, 0x55, 0xb2, 0x51, 0xb2, 0x3d, 0xf4, 0xd6, 0xd2, 0xf6, 0x13, 0xf5,
	0x0b, 0x56, 0x76, 0x36, 0x11, 0x20, 0x2e, 0xbd, 0x44, 0x79, 0xef, 0xcd, 0xb3, 0xdf, 0x78, 0x66,
	0xe1, 0x78, 0x38, 0xb3, 0xa3, 0xec, 0xf6, 0x36, 0x9b, 0xda, 0xfc, 0xdb, 0x6c, 0xbc, 0xb8, 0x9c,
	0xcd, 0xb3, 0x3c, 0x63, 0x8d, 0xe1, 0x68, 0x34, 0x5e, 0x2c, 0xec, 0x2c, 0x9b, 0x4c, 0xf3, 0xd6,
	0x5d, 0x0d, 0x76, 0x79, 0x42, 0xf3, 0x79, 0x36, 0xd7, 0xf9, 0x30, 0xff, 0xba, 0x60, 0x6f, 0xa1,
	0xbe, 0xfc, 0xd7, 0x2c, 0x9f, 0x96, 0xcf, 0xf7, 0x5e, 0x9f, 0x5d, 0xde, 0x37, 0x5c, 0x3e, 0x28,
	0x5e, 0xa2, 0x69, 0xa6, 0x0a, 0x4b, 0xeb, 0x4f, 0x15, 0x36, 0x0b, 0x8e, 0xed, 0x01, 0xf0, 0xc4,
	0xea, 0xb4, 0xd3, 0x21, 0xad, 0xb1, 0xc4, 0x8e, 0x00, 0x79, 0x62, 0x85, 0x34, 0xb6, 0x23, 0x85,
	0xa0, 0x8e, 0xa1, 0x00, 0xcb, 0x6c, 0x17, 0xb6, 0x79, 0x62, 0x89, 0xf7, 0x78, 0x28, 0xb0, 0xb2,
	0x82, 0x42, 0xc6, 0x14, 0x63, 0x95, 0x35, 0x60, 0xcb, 0xc1, 0x76, 0xaa, 0xaf, 0xb0, 0xb6, 0x12,
	0x43, 0x31, 0xe0, 0x11, 0x6e, 0xb0, 0xff, 0x80, 0xf1, 0xc4, 0xa6, 0x42, 0xa7, 0x49, 0x22, 0x95,
	0xa1, 0xc0, 0x0e, 0x48, 0x61, 0xfd, 0xde, 0x45, 0x7c, 0xc0, 0xc3, 0x88, 0xb7, 0x23, 0xc2, 0x4d,
	0xf6, 0x0c, 0x4e, 0x5c, 0x1c, 0xa3, 0x88, 0xc7, 0x5e, 0x5c, 0xdb, 0x70, 0xab, 0xc8, 0x4a, 0x42,
	0x1a, 0x9d, 0x26, 0xb8, 0xcd, 0xf6, 0x61, 0xc7, 0x15, 0xcb, 0x98, 0x2c, 0x29, 0x85, 0x50, 0x14,
	0x98, 0x30, 0x26, 0x99, 0x1a, 0xdc, 0x61, 0x07, 0xee, 0xd9, 0xdc, 0x31, 0x61, 0xd7, 0x1a, 0x52,
	0x31, 0x36, 0xd8, 0xff, 0x70, 0xc4, 0x13, 0x1b, 0x8a, 0xd0, 0x58, 0x6d, 0xb8, 0x32, 0x56, 0x76,
	0xbb, 0x9a, 0x0c, 0x7e, 0x07, 0xd6, 0x84, 0xc3, 0x7b, 0x92, 0x21, 0xdb, 0x89, 0x88, 0x2b, 0xfc,
	0xf1, 0x84, 0xa2, 0x88, 0x07, 0x57, 0x78, 0x07, 0xec, 0xb9, 0xcf, 0xeb, 0x95, 0xc7, 0x2d, 0xfe,
	0x04, 0x76, 0x06, 0x2f, 0xd6, 0x3e, 0x52, 0x03, 0x52, 0xbe, 0x25, 0x87, 0x43, 0x1e, 0x85, 0x1f,
	0x28, 0xc0, 0x5f, 0xc0, 0x5e, 0xc1, 0xcb, 0x47, 0x45, 0xb1, 0x0c, 0xc8, 0x76, 0xfa, 0x5c, 0xf4,
	0xc8, 0x76, 0x79, 0x18, 0x51, 0x80, 0xbf, 0xa1, 0xd5, 0x76, 0x43, 0x53, 0xc3, 0x9b, 0x49, 0xc6,
	0x8e, 0xa0, 0x26, 0x86, 0xb7, 0x63, 0x3f, 0xfb, 0xed, 0x7e, 0x49, 0x79, 0xc4, 0x9a, 0x50, 0xef,
	0x0f, 0xa7, 0x37, 0x5f, 0xc6, 0xcd, 0xca, 0x69, 0xf9, 0x7c, 0xb7, 0x5f, 0x52, 0x05, 0x6e, 0x6f,
	0xc2, 0x86, 0x37, 0xb6, 0xde, 0x41, 0x9d, 0x27, 0x7a, 0x31, 0xb9, 0xf9, 0xe7, 0x23, 0xea, 0x50,
	0x73, 0xbe, 0x8b, 0x6b, 0x9f, 0x62, 0xfc, 0x49, 0xce, 0xd8, 0x31, 0x1c, 0xf0, 0xc4, 0x2a, 0xea,
	0x49, 0xf7, 0xeb, 0xb3, 0x07, 0x58, 0x7a, 0x44, 0xf7, 0x42, 0x6d, 0x48, 0x61, 0x99, 0x9d, 0xf8,
	0x37, 0x5c, 0xd2, 0xa9, 0x58, 0x0b, 0x15, 0x86, 0xd0, 0x58, 0x0b, 0x24, 0xbb, 0x58, 0xbd, 0xb0,
	0x6e, 0x8c, 0xf2, 0xfa, 0xf3, 0x78, 0x94, 0xaf, 0xaf, 0x91, 0xed, 0xf7, 0x0f, 0xaf, 0x59, 0xda,
	0x96, 0x34, 0x0f, 0xdc, 0x92, 0x1e, 0xc2, 0xfe, 0x9a, 0x49, 0x93, 0x80, 0x1b, 0xc2, 0xca, 0x03,
	0x32, 0xa0, 0x88, 0x0c, 0x61, 0xf5, 0x42, 0xbb, 0x15, 0x15, 0x59, 0x3e, 0xf9, 0x28, 0x67, 0x45,
	0x30, 0xbf, 0x24, 0x4f, 0x35, 0xb2, 0x12, 0x48, 0xf8, 0x15, 0x2d, 0x17, 0x0b, 0xbd, 0xa2, 0x83,
	0x50, 0x7b, 0xbe, 0x72, 0x5d, 0xf7, 0x5f, 0xee, 0x9b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbb,
	0x00, 0x5c, 0x18, 0xd2, 0x03, 0x00, 0x00,
}
