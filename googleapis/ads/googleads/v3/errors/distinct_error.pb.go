// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/distinct_error.proto

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

// Enum describing possible distinct errors.
type DistinctErrorEnum_DistinctError int32

const (
	// Enum unspecified.
	DistinctErrorEnum_UNSPECIFIED DistinctErrorEnum_DistinctError = 0
	// The received error code is not known in this version.
	DistinctErrorEnum_UNKNOWN DistinctErrorEnum_DistinctError = 1
	// Duplicate element.
	DistinctErrorEnum_DUPLICATE_ELEMENT DistinctErrorEnum_DistinctError = 2
	// Duplicate type.
	DistinctErrorEnum_DUPLICATE_TYPE DistinctErrorEnum_DistinctError = 3
)

var DistinctErrorEnum_DistinctError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "DUPLICATE_ELEMENT",
	3: "DUPLICATE_TYPE",
}

var DistinctErrorEnum_DistinctError_value = map[string]int32{
	"UNSPECIFIED":       0,
	"UNKNOWN":           1,
	"DUPLICATE_ELEMENT": 2,
	"DUPLICATE_TYPE":    3,
}

func (x DistinctErrorEnum_DistinctError) String() string {
	return proto.EnumName(DistinctErrorEnum_DistinctError_name, int32(x))
}

func (DistinctErrorEnum_DistinctError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_273e984be5b3360c, []int{0, 0}
}

// Container for enum describing possible distinct errors.
type DistinctErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DistinctErrorEnum) Reset()         { *m = DistinctErrorEnum{} }
func (m *DistinctErrorEnum) String() string { return proto.CompactTextString(m) }
func (*DistinctErrorEnum) ProtoMessage()    {}
func (*DistinctErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_273e984be5b3360c, []int{0}
}

func (m *DistinctErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistinctErrorEnum.Unmarshal(m, b)
}
func (m *DistinctErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistinctErrorEnum.Marshal(b, m, deterministic)
}
func (m *DistinctErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistinctErrorEnum.Merge(m, src)
}
func (m *DistinctErrorEnum) XXX_Size() int {
	return xxx_messageInfo_DistinctErrorEnum.Size(m)
}
func (m *DistinctErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DistinctErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DistinctErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.DistinctErrorEnum_DistinctError", DistinctErrorEnum_DistinctError_name, DistinctErrorEnum_DistinctError_value)
	proto.RegisterType((*DistinctErrorEnum)(nil), "google.ads.googleads.v3.errors.DistinctErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/distinct_error.proto", fileDescriptor_273e984be5b3360c)
}

var fileDescriptor_273e984be5b3360c = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdf, 0x4a, 0xc3, 0x30,
	0x18, 0xc5, 0x5d, 0x07, 0x0a, 0x19, 0x6a, 0x17, 0xf0, 0x46, 0x64, 0x17, 0x7d, 0x80, 0xe4, 0x22,
	0x77, 0xf1, 0x2a, 0x5b, 0xe3, 0x18, 0xce, 0x5a, 0x70, 0xad, 0x7f, 0x28, 0x8c, 0xba, 0x94, 0x50,
	0x58, 0x93, 0xd2, 0xd4, 0x3d, 0x90, 0x97, 0x3e, 0x8a, 0x8f, 0x22, 0xf8, 0x0e, 0xd2, 0x66, 0xad,
	0xec, 0x42, 0xaf, 0x7a, 0xfa, 0xf1, 0x3b, 0xe7, 0x3b, 0xf9, 0x00, 0x91, 0x5a, 0xcb, 0x6d, 0x86,
	0x53, 0x61, 0xb0, 0x95, 0x8d, 0xda, 0x11, 0x9c, 0x55, 0x95, 0xae, 0x0c, 0x16, 0xb9, 0xa9, 0x73,
	0xb5, 0xa9, 0xd7, 0xed, 0x3f, 0x2a, 0x2b, 0x5d, 0x6b, 0x38, 0xb1, 0x24, 0x4a, 0x85, 0x41, 0xbd,
	0x09, 0xed, 0x08, 0xb2, 0xa6, 0xcb, 0xab, 0x2e, 0xb4, 0xcc, 0x71, 0xaa, 0x94, 0xae, 0xd3, 0x3a,
	0xd7, 0xca, 0x58, 0xb7, 0x57, 0x80, 0xb1, 0xbf, 0x4f, 0xe5, 0x0d, 0xcf, 0xd5, 0x5b, 0xe1, 0x3d,
	0x81, 0xd3, 0x83, 0x21, 0x3c, 0x07, 0xa3, 0x28, 0x78, 0x08, 0xf9, 0x6c, 0x71, 0xb3, 0xe0, 0xbe,
	0x7b, 0x04, 0x47, 0xe0, 0x24, 0x0a, 0x6e, 0x83, 0xfb, 0xc7, 0xc0, 0x1d, 0xc0, 0x0b, 0x30, 0xf6,
	0xa3, 0x70, 0xb9, 0x98, 0xb1, 0x15, 0x5f, 0xf3, 0x25, 0xbf, 0xe3, 0xc1, 0xca, 0x75, 0x20, 0x04,
	0x67, 0xbf, 0xe3, 0xd5, 0x73, 0xc8, 0xdd, 0xe1, 0xf4, 0x7b, 0x00, 0xbc, 0x8d, 0x2e, 0xd0, 0xff,
	0x9d, 0xa7, 0xf0, 0x60, 0x7d, 0xd8, 0x34, 0x0d, 0x07, 0x2f, 0xfe, 0xde, 0x25, 0xf5, 0x36, 0x55,
	0x12, 0xe9, 0x4a, 0x62, 0x99, 0xa9, 0xf6, 0x1d, 0xdd, 0xb9, 0xca, 0xdc, 0xfc, 0x75, 0xbd, 0x6b,
	0xfb, 0x79, 0x77, 0x86, 0x73, 0xc6, 0x3e, 0x9c, 0xc9, 0xdc, 0x86, 0x31, 0x61, 0x90, 0x95, 0x8d,
	0x8a, 0x09, 0x6a, 0x57, 0x9a, 0xcf, 0x0e, 0x48, 0x98, 0x30, 0x49, 0x0f, 0x24, 0x31, 0x49, 0x2c,
	0xf0, 0xe5, 0x78, 0x76, 0x4a, 0x29, 0x13, 0x86, 0xd2, 0x1e, 0xa1, 0x34, 0x26, 0x94, 0x5a, 0xe8,
	0xf5, 0xb8, 0x6d, 0x47, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x41, 0x9d, 0xfa, 0xda, 0x01,
	0x00, 0x00,
}
