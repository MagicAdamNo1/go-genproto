// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/click_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
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

// Request message for [ClickViewService.GetClickView][google.ads.googleads.v2.services.ClickViewService.GetClickView].
type GetClickViewRequest struct {
	// Required. The resource name of the click view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetClickViewRequest) Reset()         { *m = GetClickViewRequest{} }
func (m *GetClickViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetClickViewRequest) ProtoMessage()    {}
func (*GetClickViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fade0b32fb74bb5, []int{0}
}

func (m *GetClickViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClickViewRequest.Unmarshal(m, b)
}
func (m *GetClickViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClickViewRequest.Marshal(b, m, deterministic)
}
func (m *GetClickViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClickViewRequest.Merge(m, src)
}
func (m *GetClickViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetClickViewRequest.Size(m)
}
func (m *GetClickViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClickViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetClickViewRequest proto.InternalMessageInfo

func (m *GetClickViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetClickViewRequest)(nil), "google.ads.googleads.v2.services.GetClickViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/click_view_service.proto", fileDescriptor_7fade0b32fb74bb5)
}

var fileDescriptor_7fade0b32fb74bb5 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x8a, 0xd3, 0x40,
	0x1c, 0x27, 0x11, 0x04, 0x43, 0x05, 0x89, 0x88, 0x35, 0x0a, 0x96, 0xd2, 0x83, 0x94, 0x32, 0x03,
	0x11, 0x0f, 0x8e, 0x78, 0x98, 0x2a, 0xd4, 0x93, 0x16, 0x85, 0x1c, 0x24, 0x10, 0xa6, 0xc9, 0xbf,
	0x71, 0x30, 0xc9, 0xd4, 0x4c, 0x9a, 0x1e, 0xc4, 0x8b, 0xaf, 0xe0, 0x1b, 0x78, 0xf4, 0x0d, 0x7c,
	0x85, 0x5e, 0xbd, 0x79, 0xda, 0xc3, 0x9e, 0x76, 0xdf, 0x60, 0xd9, 0xc3, 0x92, 0x4e, 0x26, 0x4d,
	0x97, 0x0d, 0xbd, 0xfd, 0x98, 0xdf, 0xc7, 0xff, 0x6b, 0xac, 0x97, 0xb1, 0x10, 0x71, 0x02, 0x98,
	0x45, 0x12, 0x2b, 0x58, 0xa1, 0xd2, 0xc5, 0x12, 0xf2, 0x92, 0x87, 0x20, 0x71, 0x98, 0xf0, 0xf0,
	0x6b, 0x50, 0x72, 0xd8, 0x04, 0xf5, 0x1b, 0x5a, 0xe5, 0xa2, 0x10, 0xf6, 0x40, 0xe9, 0x11, 0x8b,
	0x24, 0x6a, 0xac, 0xa8, 0x74, 0x91, 0xb6, 0x3a, 0x6e, 0x57, 0x78, 0x0e, 0x52, 0xac, 0xf3, 0xc3,
	0x74, 0x95, 0xea, 0x3c, 0xd1, 0x9e, 0x15, 0xc7, 0x2c, 0xcb, 0x44, 0xc1, 0x0a, 0x2e, 0x32, 0x59,
	0xb3, 0x0f, 0x5b, 0x6c, 0x98, 0x70, 0xc8, 0x8a, 0x9a, 0x78, 0xda, 0x22, 0x96, 0x1c, 0x92, 0x28,
	0x58, 0xc0, 0x17, 0x56, 0x72, 0x91, 0xd7, 0x82, 0x47, 0x2d, 0x81, 0x2e, 0xaf, 0xa8, 0xe1, 0xd2,
	0xba, 0x3f, 0x83, 0xe2, 0x4d, 0xd5, 0x89, 0xc7, 0x61, 0xf3, 0x11, 0xbe, 0xad, 0x41, 0x16, 0xf6,
	0x07, 0xeb, 0xae, 0x16, 0x06, 0x19, 0x4b, 0xa1, 0x6f, 0x0c, 0x8c, 0x67, 0x77, 0xa6, 0xe3, 0x13,
	0x6a, 0x5e, 0xd0, 0x91, 0x35, 0xdc, 0xcf, 0x5c, 0xa3, 0x15, 0x97, 0x28, 0x14, 0x29, 0xde, 0x27,
	0xf5, 0x74, 0xc0, 0x7b, 0x96, 0x82, 0x7b, 0x6e, 0x58, 0xf7, 0x1a, 0xee, 0x93, 0x5a, 0x92, 0xfd,
	0xd7, 0xb0, 0x7a, 0xed, 0xea, 0xf6, 0x0b, 0x74, 0x6c, 0xaf, 0xe8, 0x86, 0x6e, 0x9d, 0x49, 0xa7,
	0xad, 0x59, 0x36, 0x6a, 0x4c, 0xc3, 0xb7, 0xff, 0xe9, 0xe1, 0x70, 0x3f, 0xff, 0x9d, 0xfe, 0x32,
	0x91, 0x3d, 0xa9, 0xae, 0xf3, 0xfd, 0x80, 0x79, 0x1d, 0xae, 0x65, 0x21, 0x52, 0xc8, 0x25, 0x1e,
	0xab, 0x73, 0x55, 0x09, 0x12, 0x8f, 0x7f, 0x38, 0x8f, 0xb7, 0xb4, 0xdf, 0xb5, 0x85, 0xe9, 0xa5,
	0x61, 0x8d, 0x42, 0x91, 0x1e, 0x9d, 0x66, 0xfa, 0xe0, 0xfa, 0x4e, 0xe6, 0xd5, 0x55, 0xe6, 0xc6,
	0xe7, 0x77, 0xb5, 0x35, 0x16, 0x09, 0xcb, 0x62, 0x24, 0xf2, 0x18, 0xc7, 0x90, 0xed, 0x6e, 0x86,
	0xf7, 0xc5, 0xba, 0xbf, 0xee, 0x2b, 0x0d, 0x7e, 0x9b, 0xb7, 0x66, 0x94, 0xfe, 0x31, 0x07, 0x33,
	0x15, 0x48, 0x23, 0x89, 0x14, 0xac, 0x90, 0xe7, 0xa2, 0xba, 0xb0, 0xdc, 0x6a, 0x89, 0x4f, 0x23,
	0xe9, 0x37, 0x12, 0xdf, 0x73, 0x7d, 0x2d, 0x39, 0x33, 0x47, 0xea, 0x9d, 0x10, 0x1a, 0x49, 0x42,
	0x1a, 0x11, 0x21, 0x9e, 0x4b, 0x88, 0x96, 0x2d, 0x6e, 0xef, 0xfa, 0x7c, 0x7e, 0x15, 0x00, 0x00,
	0xff, 0xff, 0x90, 0x5e, 0x10, 0x05, 0x61, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ClickViewServiceClient is the client API for ClickViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClickViewServiceClient interface {
	// Returns the requested click view in full detail.
	GetClickView(ctx context.Context, in *GetClickViewRequest, opts ...grpc.CallOption) (*resources.ClickView, error)
}

type clickViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClickViewServiceClient(cc grpc.ClientConnInterface) ClickViewServiceClient {
	return &clickViewServiceClient{cc}
}

func (c *clickViewServiceClient) GetClickView(ctx context.Context, in *GetClickViewRequest, opts ...grpc.CallOption) (*resources.ClickView, error) {
	out := new(resources.ClickView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.ClickViewService/GetClickView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClickViewServiceServer is the server API for ClickViewService service.
type ClickViewServiceServer interface {
	// Returns the requested click view in full detail.
	GetClickView(context.Context, *GetClickViewRequest) (*resources.ClickView, error)
}

// UnimplementedClickViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClickViewServiceServer struct {
}

func (*UnimplementedClickViewServiceServer) GetClickView(ctx context.Context, req *GetClickViewRequest) (*resources.ClickView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClickView not implemented")
}

func RegisterClickViewServiceServer(s *grpc.Server, srv ClickViewServiceServer) {
	s.RegisterService(&_ClickViewService_serviceDesc, srv)
}

func _ClickViewService_GetClickView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClickViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClickViewServiceServer).GetClickView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.ClickViewService/GetClickView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClickViewServiceServer).GetClickView(ctx, req.(*GetClickViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClickViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.ClickViewService",
	HandlerType: (*ClickViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClickView",
			Handler:    _ClickViewService_GetClickView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/click_view_service.proto",
}
