// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/partial_failure_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Enum describing possible partial failure errors.
type PartialFailureErrorEnum_PartialFailureError int32

const (
	// Enum unspecified.
	PartialFailureErrorEnum_UNSPECIFIED PartialFailureErrorEnum_PartialFailureError = 0
	// The received error code is not known in this version.
	PartialFailureErrorEnum_UNKNOWN PartialFailureErrorEnum_PartialFailureError = 1
	// The partial failure field was false in the request.
	// This method requires this field be set to true.
	PartialFailureErrorEnum_PARTIAL_FAILURE_MODE_REQUIRED PartialFailureErrorEnum_PartialFailureError = 2
)

var PartialFailureErrorEnum_PartialFailureError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PARTIAL_FAILURE_MODE_REQUIRED",
}

var PartialFailureErrorEnum_PartialFailureError_value = map[string]int32{
	"UNSPECIFIED":                   0,
	"UNKNOWN":                       1,
	"PARTIAL_FAILURE_MODE_REQUIRED": 2,
}

func (x PartialFailureErrorEnum_PartialFailureError) String() string {
	return proto.EnumName(PartialFailureErrorEnum_PartialFailureError_name, int32(x))
}

func (PartialFailureErrorEnum_PartialFailureError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6ff0908682efac2c, []int{0, 0}
}

// Container for enum describing possible partial failure errors.
type PartialFailureErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartialFailureErrorEnum) Reset()         { *m = PartialFailureErrorEnum{} }
func (m *PartialFailureErrorEnum) String() string { return proto.CompactTextString(m) }
func (*PartialFailureErrorEnum) ProtoMessage()    {}
func (*PartialFailureErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff0908682efac2c, []int{0}
}

func (m *PartialFailureErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartialFailureErrorEnum.Unmarshal(m, b)
}
func (m *PartialFailureErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartialFailureErrorEnum.Marshal(b, m, deterministic)
}
func (m *PartialFailureErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartialFailureErrorEnum.Merge(m, src)
}
func (m *PartialFailureErrorEnum) XXX_Size() int {
	return xxx_messageInfo_PartialFailureErrorEnum.Size(m)
}
func (m *PartialFailureErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PartialFailureErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PartialFailureErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.PartialFailureErrorEnum_PartialFailureError", PartialFailureErrorEnum_PartialFailureError_name, PartialFailureErrorEnum_PartialFailureError_value)
	proto.RegisterType((*PartialFailureErrorEnum)(nil), "google.ads.googleads.v3.errors.PartialFailureErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/partial_failure_error.proto", fileDescriptor_6ff0908682efac2c)
}

var fileDescriptor_6ff0908682efac2c = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4e, 0x83, 0x40,
	0x14, 0x86, 0x2d, 0x26, 0x9a, 0x4c, 0x17, 0x36, 0xb8, 0xd0, 0x18, 0x6d, 0x22, 0x07, 0x18, 0x16,
	0xec, 0xc6, 0xd5, 0x54, 0xa6, 0x0d, 0xb1, 0x52, 0x44, 0xc1, 0xc4, 0x90, 0x90, 0x51, 0x70, 0x42,
	0x42, 0x67, 0x70, 0x86, 0xf6, 0x40, 0x2e, 0x3d, 0x8a, 0x47, 0x71, 0xeb, 0x05, 0x0c, 0x3c, 0xdb,
	0x55, 0x75, 0xc5, 0x9f, 0xc7, 0xff, 0xfd, 0xef, 0x9f, 0x87, 0x88, 0x50, 0x4a, 0xd4, 0xa5, 0xcb,
	0x0b, 0xe3, 0x82, 0xec, 0xd4, 0xda, 0x73, 0x4b, 0xad, 0x95, 0x36, 0x6e, 0xc3, 0x75, 0x5b, 0xf1,
	0x3a, 0x7f, 0xe5, 0x55, 0xbd, 0xd2, 0x65, 0xde, 0x8f, 0x71, 0xa3, 0x55, 0xab, 0xec, 0x31, 0x00,
	0x98, 0x17, 0x06, 0x6f, 0x59, 0xbc, 0xf6, 0x30, 0xb0, 0x67, 0xe7, 0x9b, 0xec, 0xa6, 0x72, 0xb9,
	0x94, 0xaa, 0xe5, 0x6d, 0xa5, 0xa4, 0x01, 0xda, 0x79, 0x43, 0x27, 0x11, 0x84, 0x4f, 0x21, 0x9b,
	0x75, 0x14, 0x93, 0xab, 0xa5, 0x93, 0xa2, 0xe3, 0x1d, 0xbf, 0xec, 0x23, 0x34, 0x4c, 0xc2, 0xfb,
	0x88, 0x5d, 0x07, 0xd3, 0x80, 0xf9, 0xa3, 0x3d, 0x7b, 0x88, 0x0e, 0x93, 0xf0, 0x26, 0x5c, 0x3c,
	0x86, 0xa3, 0x81, 0x7d, 0x89, 0x2e, 0x22, 0x1a, 0x3f, 0x04, 0x74, 0x9e, 0x4f, 0x69, 0x30, 0x4f,
	0x62, 0x96, 0xdf, 0x2e, 0x7c, 0x96, 0xc7, 0xec, 0x2e, 0x09, 0x62, 0xe6, 0x8f, 0xac, 0xc9, 0xf7,
	0x00, 0x39, 0x2f, 0x6a, 0x89, 0xff, 0xef, 0x3d, 0x39, 0xdd, 0xb1, 0x3c, 0xea, 0x3a, 0x47, 0x83,
	0x27, 0xff, 0x97, 0x15, 0xaa, 0xe6, 0x52, 0x60, 0xa5, 0x85, 0x2b, 0x4a, 0xd9, 0xbf, 0x68, 0x73,
	0xbf, 0xa6, 0x32, 0x7f, 0x9d, 0xf3, 0x0a, 0x3e, 0xef, 0xd6, 0xfe, 0x8c, 0xd2, 0x0f, 0x6b, 0x3c,
	0x83, 0x30, 0x5a, 0x18, 0x0c, 0xb2, 0x53, 0xa9, 0x87, 0xfb, 0x95, 0xe6, 0x73, 0x63, 0xc8, 0x68,
	0x61, 0xb2, 0xad, 0x21, 0x4b, 0xbd, 0x0c, 0x0c, 0x5f, 0x96, 0x03, 0x53, 0x42, 0x68, 0x61, 0x08,
	0xd9, 0x5a, 0x08, 0x49, 0x3d, 0x42, 0xc0, 0xf4, 0x7c, 0xd0, 0xb7, 0xf3, 0x7e, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x77, 0x25, 0x10, 0x84, 0xeb, 0x01, 0x00, 0x00,
}
