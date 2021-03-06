// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ap_version.proto

package cheetah

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Access Point API version.
// This is used in the Global init message exchange to handshake client/server
// Version numbers.
type APVersion int32

const (
	APVersion_AP_VERSION_UNUSED APVersion = 0
	APVersion_AP_MAJOR_VERSION  APVersion = 0
	APVersion_AP_MINOR_VERSION  APVersion = 0
	APVersion_AP_SUB_VERSION    APVersion = 1
)

var APVersion_name = map[int32]string{
	0: "AP_VERSION_UNUSED",
	// Duplicate value: 0: "AP_MAJOR_VERSION",
	// Duplicate value: 0: "AP_MINOR_VERSION",
	1: "AP_SUB_VERSION",
}
var APVersion_value = map[string]int32{
	"AP_VERSION_UNUSED": 0,
	"AP_MAJOR_VERSION":  0,
	"AP_MINOR_VERSION":  0,
	"AP_SUB_VERSION":    1,
}

func (x APVersion) String() string {
	return proto.EnumName(APVersion_name, int32(x))
}
func (APVersion) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func init() {
	proto.RegisterEnum("cheetah.APVersion", APVersion_name, APVersion_value)
}

func init() { proto.RegisterFile("ap_version.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2c, 0x88, 0x2f,
	0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xce,
	0x48, 0x4d, 0x2d, 0x49, 0xcc, 0xd0, 0x4a, 0xe3, 0xe2, 0x74, 0x0c, 0x08, 0x83, 0xc8, 0x09, 0x89,
	0x72, 0x09, 0x3a, 0x06, 0xc4, 0x87, 0xb9, 0x06, 0x05, 0x7b, 0xfa, 0xfb, 0xc5, 0x87, 0xfa, 0x85,
	0x06, 0xbb, 0xba, 0x08, 0x30, 0x08, 0x89, 0x70, 0x09, 0x38, 0x06, 0xc4, 0xfb, 0x3a, 0x7a, 0xf9,
	0x07, 0xc1, 0x24, 0x11, 0xa2, 0x9e, 0x7e, 0x28, 0xa2, 0x42, 0x5c, 0x7c, 0x8e, 0x01, 0xf1, 0xc1,
	0xa1, 0x4e, 0x70, 0x31, 0x46, 0x29, 0x26, 0x01, 0xc6, 0x24, 0x36, 0xb0, 0xbd, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xfc, 0x6b, 0x8a, 0xd2, 0x8b, 0x00, 0x00, 0x00,
}
