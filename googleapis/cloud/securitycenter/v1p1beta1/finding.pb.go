// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1p1beta1/finding.proto

package securitycenter

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// The state of the finding.
type Finding_State int32

const (
	// Unspecified state.
	Finding_STATE_UNSPECIFIED Finding_State = 0
	// The finding requires attention and has not been addressed yet.
	Finding_ACTIVE Finding_State = 1
	// The finding has been fixed, triaged as a non-issue or otherwise addressed
	// and is no longer active.
	Finding_INACTIVE Finding_State = 2
)

var Finding_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "ACTIVE",
	2: "INACTIVE",
}

var Finding_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"ACTIVE":            1,
	"INACTIVE":          2,
}

func (x Finding_State) String() string {
	return proto.EnumName(Finding_State_name, int32(x))
}

func (Finding_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_238d4dee941eabf6, []int{0, 0}
}

// Cloud Security Command Center (Cloud SCC) finding.
//
// A finding is a record of assessment data (security, risk, health or privacy)
// ingested into Cloud SCC for presentation, notification, analysis,
// policy testing, and enforcement. For example, an XSS vulnerability in an
// App Engine application is a finding.
type Finding struct {
	// The relative resource name of this finding. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/{organization_id}/sources/{source_id}/findings/{finding_id}"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The relative resource name of the source the finding belongs to. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// This field is immutable after creation time.
	// For example:
	// "organizations/{organization_id}/sources/{source_id}"
	Parent string `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	// For findings on Google Cloud Platform (GCP) resources, the full resource
	// name of the GCP resource this finding is for. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	// When the finding is for a non-GCP resource, the resourceName can be a
	// customer or partner defined string.
	// This field is immutable after creation time.
	ResourceName string `protobuf:"bytes,3,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The state of the finding.
	State Finding_State `protobuf:"varint,4,opt,name=state,proto3,enum=google.cloud.securitycenter.v1p1beta1.Finding_State" json:"state,omitempty"`
	// The additional taxonomy group within findings from a given source.
	// This field is immutable after creation time.
	// Example: "XSS_FLASH_INJECTION"
	Category string `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	// The URI that, if available, points to a web page outside of Cloud SCC
	// where additional information about the finding can be found. This field is
	// guaranteed to be either empty or a well formed URL.
	ExternalUri string `protobuf:"bytes,6,opt,name=external_uri,json=externalUri,proto3" json:"external_uri,omitempty"`
	// Source specific properties. These properties are managed by the source
	// that writes the finding. The key names in the source_properties map must be
	// between 1 and 255 characters, and must start with a letter and contain
	// alphanumeric characters or underscores only.
	SourceProperties map[string]*_struct.Value `protobuf:"bytes,7,rep,name=source_properties,json=sourceProperties,proto3" json:"source_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output only. User specified security marks. These marks are entirely
	// managed by the user and come from the SecurityMarks resource that belongs
	// to the finding.
	SecurityMarks *SecurityMarks `protobuf:"bytes,8,opt,name=security_marks,json=securityMarks,proto3" json:"security_marks,omitempty"`
	// The time at which the event took place. For example, if the finding
	// represents an open firewall it would capture the time the detector believes
	// the firewall became open. The accuracy is determined by the detector.
	EventTime *timestamp.Timestamp `protobuf:"bytes,9,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// The time at which the finding was created in Cloud SCC.
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Finding) Reset()         { *m = Finding{} }
func (m *Finding) String() string { return proto.CompactTextString(m) }
func (*Finding) ProtoMessage()    {}
func (*Finding) Descriptor() ([]byte, []int) {
	return fileDescriptor_238d4dee941eabf6, []int{0}
}

func (m *Finding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Finding.Unmarshal(m, b)
}
func (m *Finding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Finding.Marshal(b, m, deterministic)
}
func (m *Finding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Finding.Merge(m, src)
}
func (m *Finding) XXX_Size() int {
	return xxx_messageInfo_Finding.Size(m)
}
func (m *Finding) XXX_DiscardUnknown() {
	xxx_messageInfo_Finding.DiscardUnknown(m)
}

var xxx_messageInfo_Finding proto.InternalMessageInfo

func (m *Finding) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Finding) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Finding) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Finding) GetState() Finding_State {
	if m != nil {
		return m.State
	}
	return Finding_STATE_UNSPECIFIED
}

func (m *Finding) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Finding) GetExternalUri() string {
	if m != nil {
		return m.ExternalUri
	}
	return ""
}

func (m *Finding) GetSourceProperties() map[string]*_struct.Value {
	if m != nil {
		return m.SourceProperties
	}
	return nil
}

func (m *Finding) GetSecurityMarks() *SecurityMarks {
	if m != nil {
		return m.SecurityMarks
	}
	return nil
}

func (m *Finding) GetEventTime() *timestamp.Timestamp {
	if m != nil {
		return m.EventTime
	}
	return nil
}

func (m *Finding) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.securitycenter.v1p1beta1.Finding_State", Finding_State_name, Finding_State_value)
	proto.RegisterType((*Finding)(nil), "google.cloud.securitycenter.v1p1beta1.Finding")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "google.cloud.securitycenter.v1p1beta1.Finding.SourcePropertiesEntry")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1p1beta1/finding.proto", fileDescriptor_238d4dee941eabf6)
}

var fileDescriptor_238d4dee941eabf6 = []byte{
	// 646 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdf, 0x6e, 0x12, 0x4d,
	0x14, 0xff, 0x16, 0x0a, 0x6d, 0x87, 0xb6, 0xa1, 0x93, 0xb4, 0xd9, 0x8f, 0x98, 0x88, 0x35, 0x4d,
	0x30, 0x31, 0x3b, 0xa1, 0xf5, 0xa2, 0x6e, 0x6f, 0xa4, 0x94, 0x1a, 0x4c, 0x6c, 0x08, 0x50, 0x2e,
	0xb4, 0x09, 0x0e, 0xdb, 0xe9, 0x3a, 0xe9, 0xee, 0xcc, 0x3a, 0x3b, 0x4b, 0xc4, 0xa6, 0x2f, 0xe4,
	0x0b, 0xf8, 0x0e, 0x3e, 0x85, 0xd7, 0x7d, 0x07, 0x13, 0xb3, 0x33, 0x3b, 0x08, 0x68, 0x52, 0xbc,
	0x9b, 0x73, 0x7e, 0x7f, 0xce, 0xe1, 0x9c, 0xc3, 0x82, 0x43, 0x9f, 0x73, 0x3f, 0x20, 0xc8, 0x0b,
	0x78, 0x72, 0x85, 0x62, 0xe2, 0x25, 0x82, 0xca, 0x89, 0x47, 0x98, 0x24, 0x02, 0x8d, 0xeb, 0x51,
	0x7d, 0x44, 0x24, 0xae, 0xa3, 0x6b, 0xca, 0xae, 0x28, 0xf3, 0x9d, 0x48, 0x70, 0xc9, 0xe1, 0xbe,
	0x16, 0x39, 0x4a, 0xe4, 0xcc, 0x8b, 0x9c, 0xa9, 0xa8, 0xf2, 0x28, 0xf3, 0xc6, 0x11, 0x45, 0x98,
	0x31, 0x2e, 0xb1, 0xa4, 0x9c, 0xc5, 0xda, 0xa4, 0xf2, 0x78, 0x06, 0xbd, 0xa6, 0x24, 0xb8, 0x1a,
	0x8e, 0xc8, 0x47, 0x3c, 0xa6, 0x5c, 0x64, 0x84, 0xff, 0x67, 0x08, 0x82, 0xc4, 0x3c, 0x11, 0x1e,
	0xc9, 0x20, 0x77, 0xb9, 0xae, 0x0d, 0x30, 0x0c, 0xb1, 0xb8, 0x31, 0x75, 0x4d, 0x57, 0x2a, 0x1a,
	0x25, 0xd7, 0x28, 0x96, 0x22, 0xf1, 0xe4, 0x42, 0x57, 0x53, 0x54, 0xd2, 0x90, 0xc4, 0x12, 0x87,
	0x91, 0x26, 0xec, 0x7d, 0x2b, 0x82, 0xd5, 0x33, 0x3d, 0x0d, 0x08, 0xc1, 0x0a, 0xc3, 0x21, 0xb1,
	0xad, 0xaa, 0x55, 0x5b, 0xef, 0xaa, 0x37, 0xdc, 0x05, 0xc5, 0x08, 0x0b, 0xc2, 0xa4, 0x9d, 0x53,
	0xd9, 0x2c, 0x82, 0x4f, 0xc1, 0xa6, 0xf9, 0x11, 0x43, 0x25, 0xca, 0x2b, 0x78, 0xc3, 0x24, 0xcf,
	0x53, 0xf1, 0x1b, 0x50, 0x88, 0x25, 0x96, 0xc4, 0x5e, 0xa9, 0x5a, 0xb5, 0xad, 0x83, 0x17, 0xce,
	0x52, 0x83, 0x76, 0xb2, 0x7e, 0x9c, 0x5e, 0xaa, 0xed, 0x6a, 0x0b, 0x58, 0x01, 0x6b, 0x1e, 0x96,
	0xc4, 0xe7, 0x62, 0x62, 0x17, 0x54, 0xad, 0x69, 0x0c, 0x9f, 0x80, 0x0d, 0xf2, 0x59, 0x12, 0xc1,
	0x70, 0x30, 0x4c, 0x04, 0xb5, 0x8b, 0x0a, 0x2f, 0x99, 0xdc, 0x85, 0xa0, 0xf0, 0x13, 0xd8, 0xce,
	0xba, 0x8d, 0x04, 0x8f, 0x88, 0x90, 0x94, 0xc4, 0xf6, 0x6a, 0x35, 0x5f, 0x2b, 0x1d, 0x9c, 0xfe,
	0x6b, 0x5b, 0xca, 0xa7, 0x33, 0xb5, 0x69, 0x31, 0x29, 0x26, 0xdd, 0x72, 0xbc, 0x90, 0x86, 0x1f,
	0xc0, 0xd6, 0xfc, 0xc6, 0xec, 0xb5, 0xaa, 0x55, 0x2b, 0x2d, 0x3d, 0x86, 0x5e, 0x06, 0xbc, 0x4d,
	0xb5, 0x27, 0xf9, 0x1f, 0x8d, 0x7c, 0x77, 0x33, 0x9e, 0xcd, 0xc1, 0x97, 0x00, 0x90, 0x31, 0x61,
	0x72, 0x98, 0x6e, 0xd5, 0x5e, 0x57, 0xee, 0x15, 0xe3, 0x6e, 0x56, 0xee, 0xf4, 0xcd, 0xca, 0xbb,
	0xeb, 0x8a, 0x9d, 0xc6, 0xf0, 0x18, 0x94, 0x3c, 0x41, 0xb0, 0x24, 0x5a, 0x0b, 0x1e, 0xd4, 0x02,
	0x4d, 0x4f, 0x13, 0x95, 0xf7, 0x60, 0xe7, 0xaf, 0x43, 0x80, 0x65, 0x90, 0xbf, 0x21, 0x93, 0xec,
	0x80, 0xd2, 0x27, 0x7c, 0x0e, 0x0a, 0x63, 0x1c, 0x24, 0x44, 0x9d, 0x4f, 0xe9, 0x60, 0xf7, 0x8f,
	0x0a, 0x83, 0x14, 0xed, 0x6a, 0x92, 0x9b, 0x3b, 0xb2, 0xf6, 0x8e, 0x40, 0x41, 0x2d, 0x1e, 0xee,
	0x80, 0xed, 0x5e, 0xbf, 0xd1, 0x6f, 0x0d, 0x2f, 0xce, 0x7b, 0x9d, 0x56, 0xb3, 0x7d, 0xd6, 0x6e,
	0x9d, 0x96, 0xff, 0x83, 0x00, 0x14, 0x1b, 0xcd, 0x7e, 0x7b, 0xd0, 0x2a, 0x5b, 0x70, 0x03, 0xac,
	0xb5, 0xcf, 0xb3, 0x28, 0xe7, 0x06, 0xf7, 0x0d, 0x0a, 0xf6, 0x17, 0xe6, 0xa9, 0xeb, 0xe1, 0x88,
	0xc6, 0x8e, 0xc7, 0x43, 0x64, 0x6e, 0xfd, 0x15, 0x17, 0x3e, 0x66, 0xf4, 0x8b, 0xfe, 0x0f, 0xa3,
	0xdb, 0xd9, 0xf0, 0x0e, 0xe9, 0x5d, 0xc6, 0xe8, 0x56, 0x3f, 0xee, 0xcc, 0x07, 0x23, 0x46, 0xb7,
	0xd9, 0xeb, 0xee, 0xe4, 0xa7, 0x05, 0x9e, 0x79, 0x3c, 0x5c, 0x6e, 0x99, 0x1d, 0xeb, 0x5d, 0x2f,
	0x23, 0xfa, 0x3c, 0xc0, 0xcc, 0x77, 0xb8, 0xf0, 0x91, 0x4f, 0x98, 0x9a, 0x03, 0xfa, 0xdd, 0xe4,
	0x03, 0xdf, 0x80, 0xe3, 0x79, 0xe0, 0x6b, 0x6e, 0xff, 0xb5, 0x76, 0x6d, 0xaa, 0xf2, 0xe6, 0x64,
	0x9a, 0xba, 0xfc, 0xa0, 0xde, 0xa9, 0x9f, 0xa4, 0xb2, 0xef, 0x86, 0x77, 0xa9, 0x78, 0x97, 0xf3,
	0xbc, 0xcb, 0x81, 0xb1, 0xbf, 0xcf, 0xd5, 0x34, 0xcf, 0x75, 0x15, 0xd1, 0x75, 0xe7, 0x99, 0xae,
	0x9b, 0x52, 0x95, 0xe5, 0xa8, 0xa8, 0x5a, 0x3f, 0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0x46, 0x85,
	0xf0, 0x05, 0x73, 0x05, 0x00, 0x00,
}