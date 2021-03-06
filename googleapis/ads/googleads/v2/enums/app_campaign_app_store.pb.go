// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/app_campaign_app_store.proto

package enums

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

// Enum describing app campaign app store.
type AppCampaignAppStoreEnum_AppCampaignAppStore int32

const (
	// Not specified.
	AppCampaignAppStoreEnum_UNSPECIFIED AppCampaignAppStoreEnum_AppCampaignAppStore = 0
	// Used for return value only. Represents value unknown in this version.
	AppCampaignAppStoreEnum_UNKNOWN AppCampaignAppStoreEnum_AppCampaignAppStore = 1
	// Apple app store.
	AppCampaignAppStoreEnum_APPLE_APP_STORE AppCampaignAppStoreEnum_AppCampaignAppStore = 2
	// Google play.
	AppCampaignAppStoreEnum_GOOGLE_APP_STORE AppCampaignAppStoreEnum_AppCampaignAppStore = 3
)

var AppCampaignAppStoreEnum_AppCampaignAppStore_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "APPLE_APP_STORE",
	3: "GOOGLE_APP_STORE",
}

var AppCampaignAppStoreEnum_AppCampaignAppStore_value = map[string]int32{
	"UNSPECIFIED":      0,
	"UNKNOWN":          1,
	"APPLE_APP_STORE":  2,
	"GOOGLE_APP_STORE": 3,
}

func (x AppCampaignAppStoreEnum_AppCampaignAppStore) String() string {
	return proto.EnumName(AppCampaignAppStoreEnum_AppCampaignAppStore_name, int32(x))
}

func (AppCampaignAppStoreEnum_AppCampaignAppStore) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b3345cd5808d0893, []int{0, 0}
}

// The application store that distributes mobile applications.
type AppCampaignAppStoreEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppCampaignAppStoreEnum) Reset()         { *m = AppCampaignAppStoreEnum{} }
func (m *AppCampaignAppStoreEnum) String() string { return proto.CompactTextString(m) }
func (*AppCampaignAppStoreEnum) ProtoMessage()    {}
func (*AppCampaignAppStoreEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3345cd5808d0893, []int{0}
}

func (m *AppCampaignAppStoreEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppCampaignAppStoreEnum.Unmarshal(m, b)
}
func (m *AppCampaignAppStoreEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppCampaignAppStoreEnum.Marshal(b, m, deterministic)
}
func (m *AppCampaignAppStoreEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppCampaignAppStoreEnum.Merge(m, src)
}
func (m *AppCampaignAppStoreEnum) XXX_Size() int {
	return xxx_messageInfo_AppCampaignAppStoreEnum.Size(m)
}
func (m *AppCampaignAppStoreEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AppCampaignAppStoreEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AppCampaignAppStoreEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.AppCampaignAppStoreEnum_AppCampaignAppStore", AppCampaignAppStoreEnum_AppCampaignAppStore_name, AppCampaignAppStoreEnum_AppCampaignAppStore_value)
	proto.RegisterType((*AppCampaignAppStoreEnum)(nil), "google.ads.googleads.v2.enums.AppCampaignAppStoreEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/app_campaign_app_store.proto", fileDescriptor_b3345cd5808d0893)
}

var fileDescriptor_b3345cd5808d0893 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x75, 0x1d, 0x28, 0x64, 0x0f, 0x2b, 0x9d, 0xa0, 0x88, 0x7b, 0xd8, 0x3e, 0x20, 0x81, 0xfa,
	0x16, 0x9f, 0xb2, 0x59, 0xcb, 0x50, 0xda, 0xe0, 0xdc, 0x04, 0x29, 0x96, 0xb8, 0x96, 0x50, 0x58,
	0x93, 0xb0, 0x74, 0x03, 0x7f, 0xc7, 0x47, 0x3f, 0xc5, 0x4f, 0x11, 0xfc, 0x07, 0x49, 0xb2, 0x0d,
	0x84, 0xe9, 0x4b, 0x38, 0xb9, 0xe7, 0x9e, 0xc3, 0xb9, 0x07, 0x60, 0x2e, 0x25, 0x5f, 0x96, 0x88,
	0x15, 0x1a, 0x39, 0x68, 0xd0, 0x26, 0x44, 0xa5, 0x58, 0xd7, 0x1a, 0x31, 0xa5, 0xf2, 0x05, 0xab,
	0x15, 0xab, 0xb8, 0xc8, 0xcd, 0x47, 0x37, 0x72, 0x55, 0x42, 0xb5, 0x92, 0x8d, 0x0c, 0xfa, 0x4e,
	0x00, 0x59, 0xa1, 0xe1, 0x5e, 0x0b, 0x37, 0x21, 0xb4, 0xda, 0x8b, 0xcb, 0x9d, 0xb5, 0xaa, 0x10,
	0x13, 0x42, 0x36, 0xac, 0xa9, 0xa4, 0xd0, 0x4e, 0x3c, 0x7c, 0x03, 0x67, 0x44, 0xa9, 0xf1, 0xd6,
	0x9b, 0x28, 0x35, 0x35, 0xce, 0x91, 0x58, 0xd7, 0xc3, 0x17, 0xd0, 0x3b, 0x40, 0x05, 0x5d, 0xd0,
	0x99, 0x25, 0x53, 0x1a, 0x8d, 0x27, 0xb7, 0x93, 0xe8, 0xc6, 0x3f, 0x0a, 0x3a, 0xe0, 0x64, 0x96,
	0xdc, 0x25, 0xe9, 0x53, 0xe2, 0xb7, 0x82, 0x1e, 0xe8, 0x12, 0x4a, 0xef, 0xa3, 0x9c, 0x50, 0x9a,
	0x4f, 0x1f, 0xd3, 0x87, 0xc8, 0xf7, 0x82, 0x53, 0xe0, 0xc7, 0x69, 0x1a, 0xff, 0x9a, 0xb6, 0x47,
	0xdf, 0x2d, 0x30, 0x58, 0xc8, 0x1a, 0xfe, 0x1b, 0x7f, 0x74, 0x7e, 0x20, 0x03, 0x35, 0xd1, 0x69,
	0xeb, 0x79, 0xb4, 0x95, 0x72, 0xb9, 0x64, 0x82, 0x43, 0xb9, 0xe2, 0x88, 0x97, 0xc2, 0x1e, 0xb6,
	0x6b, 0x51, 0x55, 0xfa, 0x8f, 0x52, 0xaf, 0xed, 0xfb, 0xee, 0xb5, 0x63, 0x42, 0x3e, 0xbc, 0x7e,
	0xec, 0xac, 0x48, 0xa1, 0xa1, 0x83, 0x06, 0xcd, 0x43, 0x68, 0xaa, 0xd0, 0x9f, 0x3b, 0x3e, 0x23,
	0x85, 0xce, 0xf6, 0x7c, 0x36, 0x0f, 0x33, 0xcb, 0x7f, 0x79, 0x03, 0x37, 0xc4, 0x98, 0x14, 0x1a,
	0xe3, 0xfd, 0x06, 0xc6, 0xf3, 0x10, 0x63, 0xbb, 0xf3, 0x7a, 0x6c, 0x83, 0x5d, 0xfd, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x25, 0xdd, 0x06, 0x91, 0xec, 0x01, 0x00, 0x00,
}
