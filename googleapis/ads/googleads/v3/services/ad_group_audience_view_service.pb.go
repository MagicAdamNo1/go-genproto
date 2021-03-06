// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/ad_group_audience_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [AdGroupAudienceViewService.GetAdGoupAudienceView][].
type GetAdGroupAudienceViewRequest struct {
	// Required. The resource name of the ad group audience view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupAudienceViewRequest) Reset()         { *m = GetAdGroupAudienceViewRequest{} }
func (m *GetAdGroupAudienceViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupAudienceViewRequest) ProtoMessage()    {}
func (*GetAdGroupAudienceViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_88c843f8d448d38a, []int{0}
}

func (m *GetAdGroupAudienceViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupAudienceViewRequest.Unmarshal(m, b)
}
func (m *GetAdGroupAudienceViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupAudienceViewRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupAudienceViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupAudienceViewRequest.Merge(m, src)
}
func (m *GetAdGroupAudienceViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupAudienceViewRequest.Size(m)
}
func (m *GetAdGroupAudienceViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupAudienceViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupAudienceViewRequest proto.InternalMessageInfo

func (m *GetAdGroupAudienceViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupAudienceViewRequest)(nil), "google.ads.googleads.v3.services.GetAdGroupAudienceViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/ad_group_audience_view_service.proto", fileDescriptor_88c843f8d448d38a)
}

var fileDescriptor_88c843f8d448d38a = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x31, 0x8b, 0xd4, 0x40,
	0x18, 0x65, 0x23, 0x08, 0x06, 0x6d, 0x52, 0xe8, 0x19, 0x95, 0x5b, 0x8e, 0x2b, 0xe4, 0x90, 0x19,
	0x30, 0x72, 0xc8, 0x88, 0xca, 0x2c, 0xc8, 0x5a, 0x1d, 0x87, 0xc2, 0x82, 0x12, 0x08, 0x73, 0x99,
	0xcf, 0x38, 0x90, 0xcc, 0xc4, 0x99, 0x24, 0x07, 0x8a, 0x8d, 0x85, 0x7f, 0xc0, 0xda, 0xc6, 0xd2,
	0x9f, 0x72, 0xad, 0x9d, 0x95, 0x82, 0x95, 0x3f, 0xc1, 0x4a, 0x72, 0x33, 0x93, 0xdb, 0xc3, 0xc4,
	0xed, 0x1e, 0xfb, 0xde, 0xf7, 0xbe, 0x37, 0xef, 0xdb, 0x84, 0x4f, 0x0a, 0xa5, 0x8a, 0x12, 0x30,
	0xe3, 0x06, 0x5b, 0xd8, 0xa3, 0x2e, 0xc1, 0x06, 0x74, 0x27, 0x72, 0x30, 0x98, 0xf1, 0xac, 0xd0,
	0xaa, 0xad, 0x33, 0xd6, 0x72, 0x01, 0x32, 0x87, 0xac, 0x13, 0x70, 0x9c, 0x39, 0x1e, 0xd5, 0x5a,
	0x35, 0x2a, 0x9a, 0xdb, 0x59, 0xc4, 0xb8, 0x41, 0x83, 0x0d, 0xea, 0x12, 0xe4, 0x6d, 0xe2, 0x47,
	0x53, 0x8b, 0x34, 0x18, 0xd5, 0xea, 0xe9, 0x4d, 0x76, 0x43, 0x7c, 0xd3, 0xcf, 0xd7, 0x02, 0x33,
	0x29, 0x55, 0xc3, 0x1a, 0xa1, 0xa4, 0x71, 0xec, 0xb5, 0x35, 0x36, 0x2f, 0x05, 0xc8, 0xc6, 0x11,
	0xdb, 0x6b, 0xc4, 0x2b, 0x01, 0x25, 0xcf, 0x8e, 0xe0, 0x35, 0xeb, 0x84, 0xd2, 0x4e, 0x70, 0x7d,
	0x4d, 0xe0, 0xa3, 0x58, 0x6a, 0xe7, 0x6d, 0x78, 0x6b, 0x09, 0x0d, 0xe5, 0xcb, 0x3e, 0x14, 0x75,
	0x99, 0x56, 0x02, 0x8e, 0x9f, 0xc1, 0x9b, 0x16, 0x4c, 0x13, 0xbd, 0x08, 0xaf, 0xf8, 0x91, 0x4c,
	0xb2, 0x0a, 0xb6, 0x66, 0xf3, 0xd9, 0xed, 0x4b, 0x8b, 0x7b, 0x3f, 0x68, 0xf0, 0x87, 0xa2, 0xf0,
	0xce, 0x59, 0x13, 0x0e, 0xd5, 0xc2, 0xa0, 0x5c, 0x55, 0x78, 0xcc, 0xf3, 0xb2, 0xb7, 0x3a, 0x60,
	0x15, 0xdc, 0xfd, 0x1c, 0x84, 0xf1, 0x88, 0xea, 0xb9, 0xad, 0x33, 0xfa, 0x39, 0x0b, 0xaf, 0x8e,
	0x67, 0x8b, 0x1e, 0xa3, 0x4d, 0xb7, 0x40, 0xff, 0x7d, 0x55, 0xbc, 0x3f, 0x69, 0x30, 0x9c, 0x0a,
	0x8d, 0x8c, 0xef, 0x1c, 0x7c, 0xa7, 0xe7, 0xeb, 0xf8, 0xf0, 0xed, 0xd7, 0xa7, 0xe0, 0x7e, 0xb4,
	0xdf, 0x5f, 0xf9, 0xdd, 0x39, 0xe6, 0x61, 0xde, 0x9a, 0x46, 0x55, 0xa0, 0x0d, 0xde, 0xc3, 0xec,
	0x5f, 0x2f, 0x83, 0xf7, 0xde, 0xc7, 0x37, 0x4e, 0xe8, 0xd6, 0x54, 0x83, 0x8b, 0x8f, 0x41, 0xb8,
	0x9b, 0xab, 0x6a, 0xe3, 0x5b, 0x17, 0xdb, 0xd3, 0x2d, 0x1e, 0xf6, 0x57, 0x3e, 0x9c, 0xbd, 0x7c,
	0xea, 0x4c, 0x0a, 0x55, 0x32, 0x59, 0x20, 0xa5, 0x0b, 0x5c, 0x80, 0x3c, 0xfd, 0x0f, 0xe0, 0xb3,
	0xb5, 0xd3, 0x9f, 0xc8, 0x03, 0x0f, 0xbe, 0x04, 0x17, 0x96, 0x94, 0x7e, 0x0d, 0xe6, 0x4b, 0x6b,
	0x48, 0xb9, 0x41, 0x16, 0xf6, 0x68, 0x95, 0x20, 0xb7, 0xd8, 0x9c, 0x78, 0x49, 0x4a, 0xb9, 0x49,
	0x07, 0x49, 0xba, 0x4a, 0x52, 0x2f, 0xf9, 0x1d, 0xec, 0xda, 0xdf, 0x09, 0xa1, 0xdc, 0x10, 0x32,
	0x88, 0x08, 0x59, 0x25, 0x84, 0x78, 0xd9, 0xd1, 0xc5, 0xd3, 0x9c, 0xc9, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xed, 0xe5, 0x93, 0xa6, 0xc9, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdGroupAudienceViewServiceClient is the client API for AdGroupAudienceViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupAudienceViewServiceClient interface {
	// Returns the requested ad group audience view in full detail.
	GetAdGroupAudienceView(ctx context.Context, in *GetAdGroupAudienceViewRequest, opts ...grpc.CallOption) (*resources.AdGroupAudienceView, error)
}

type adGroupAudienceViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupAudienceViewServiceClient(cc grpc.ClientConnInterface) AdGroupAudienceViewServiceClient {
	return &adGroupAudienceViewServiceClient{cc}
}

func (c *adGroupAudienceViewServiceClient) GetAdGroupAudienceView(ctx context.Context, in *GetAdGroupAudienceViewRequest, opts ...grpc.CallOption) (*resources.AdGroupAudienceView, error) {
	out := new(resources.AdGroupAudienceView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.AdGroupAudienceViewService/GetAdGroupAudienceView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupAudienceViewServiceServer is the server API for AdGroupAudienceViewService service.
type AdGroupAudienceViewServiceServer interface {
	// Returns the requested ad group audience view in full detail.
	GetAdGroupAudienceView(context.Context, *GetAdGroupAudienceViewRequest) (*resources.AdGroupAudienceView, error)
}

// UnimplementedAdGroupAudienceViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupAudienceViewServiceServer struct {
}

func (*UnimplementedAdGroupAudienceViewServiceServer) GetAdGroupAudienceView(ctx context.Context, req *GetAdGroupAudienceViewRequest) (*resources.AdGroupAudienceView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdGroupAudienceView not implemented")
}

func RegisterAdGroupAudienceViewServiceServer(s *grpc.Server, srv AdGroupAudienceViewServiceServer) {
	s.RegisterService(&_AdGroupAudienceViewService_serviceDesc, srv)
}

func _AdGroupAudienceViewService_GetAdGroupAudienceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupAudienceViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAudienceViewServiceServer).GetAdGroupAudienceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.AdGroupAudienceViewService/GetAdGroupAudienceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAudienceViewServiceServer).GetAdGroupAudienceView(ctx, req.(*GetAdGroupAudienceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupAudienceViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.AdGroupAudienceViewService",
	HandlerType: (*AdGroupAudienceViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupAudienceView",
			Handler:    _AdGroupAudienceViewService_GetAdGroupAudienceView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/ad_group_audience_view_service.proto",
}
