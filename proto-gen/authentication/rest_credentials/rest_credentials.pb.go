// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rest_credentials.proto

package rest_credentials

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

//Credentials values generated for a rest application.
type RestAppDefinition struct {
	ProjectId            string   `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret         string   `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	AuthUri              string   `protobuf:"bytes,4,opt,name=auth_uri,json=authUri,proto3" json:"auth_uri,omitempty"`
	RedirectUris         []string `protobuf:"bytes,5,rep,name=redirect_uris,json=redirectUris,proto3" json:"redirect_uris,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestAppDefinition) Reset()         { *m = RestAppDefinition{} }
func (m *RestAppDefinition) String() string { return proto.CompactTextString(m) }
func (*RestAppDefinition) ProtoMessage()    {}
func (*RestAppDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_2217ae0d6dc6a76b, []int{0}
}

func (m *RestAppDefinition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestAppDefinition.Unmarshal(m, b)
}
func (m *RestAppDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestAppDefinition.Marshal(b, m, deterministic)
}
func (m *RestAppDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestAppDefinition.Merge(m, src)
}
func (m *RestAppDefinition) XXX_Size() int {
	return xxx_messageInfo_RestAppDefinition.Size(m)
}
func (m *RestAppDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_RestAppDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_RestAppDefinition proto.InternalMessageInfo

func (m *RestAppDefinition) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *RestAppDefinition) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *RestAppDefinition) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *RestAppDefinition) GetAuthUri() string {
	if m != nil {
		return m.AuthUri
	}
	return ""
}

func (m *RestAppDefinition) GetRedirectUris() []string {
	if m != nil {
		return m.RedirectUris
	}
	return nil
}

//Create a response for the rest application with credentials data.
type CreateRestResponse struct {
	Rest                 *RestAppDefinition `protobuf:"bytes,1,opt,name=rest,proto3" json:"rest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateRestResponse) Reset()         { *m = CreateRestResponse{} }
func (m *CreateRestResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRestResponse) ProtoMessage()    {}
func (*CreateRestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2217ae0d6dc6a76b, []int{1}
}

func (m *CreateRestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRestResponse.Unmarshal(m, b)
}
func (m *CreateRestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRestResponse.Marshal(b, m, deterministic)
}
func (m *CreateRestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRestResponse.Merge(m, src)
}
func (m *CreateRestResponse) XXX_Size() int {
	return xxx_messageInfo_CreateRestResponse.Size(m)
}
func (m *CreateRestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRestResponse proto.InternalMessageInfo

func (m *CreateRestResponse) GetRest() *RestAppDefinition {
	if m != nil {
		return m.Rest
	}
	return nil
}

// Data needed to create credentials.
type CreateRestRequest struct {
	ProjectId            string   `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	RedirectUris         string   `protobuf:"bytes,2,opt,name=redirect_uris,json=redirectUris,proto3" json:"redirect_uris,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRestRequest) Reset()         { *m = CreateRestRequest{} }
func (m *CreateRestRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRestRequest) ProtoMessage()    {}
func (*CreateRestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2217ae0d6dc6a76b, []int{2}
}

func (m *CreateRestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRestRequest.Unmarshal(m, b)
}
func (m *CreateRestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRestRequest.Marshal(b, m, deterministic)
}
func (m *CreateRestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRestRequest.Merge(m, src)
}
func (m *CreateRestRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRestRequest.Size(m)
}
func (m *CreateRestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRestRequest proto.InternalMessageInfo

func (m *CreateRestRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *CreateRestRequest) GetRedirectUris() string {
	if m != nil {
		return m.RedirectUris
	}
	return ""
}

func init() {
	proto.RegisterType((*RestAppDefinition)(nil), "rest_credentials.RestAppDefinition")
	proto.RegisterType((*CreateRestResponse)(nil), "rest_credentials.CreateRestResponse")
	proto.RegisterType((*CreateRestRequest)(nil), "rest_credentials.CreateRestRequest")
}

func init() { proto.RegisterFile("rest_credentials.proto", fileDescriptor_2217ae0d6dc6a76b) }

var fileDescriptor_2217ae0d6dc6a76b = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcb, 0x6e, 0xd3, 0x4c,
	0x18, 0x95, 0x9d, 0xde, 0x32, 0xff, 0x8f, 0x48, 0x07, 0x84, 0x8c, 0x01, 0x69, 0xea, 0x22, 0x84,
	0xa2, 0xd6, 0x0e, 0x21, 0x08, 0x29, 0x6c, 0x28, 0x85, 0x45, 0x45, 0x91, 0x90, 0x43, 0x84, 0x60,
	0x13, 0x0d, 0xf6, 0x17, 0x67, 0x2a, 0x67, 0x66, 0x98, 0x4b, 0x53, 0x96, 0xf0, 0x08, 0x20, 0x16,
	0xe5, 0x39, 0x78, 0x13, 0x5e, 0x81, 0x07, 0x41, 0x1e, 0xa7, 0xa2, 0x60, 0x04, 0x2b, 0xcb, 0xe7,
	0xe2, 0x73, 0xe6, 0xd8, 0x46, 0x57, 0x14, 0x68, 0x33, 0xc9, 0x14, 0xe4, 0xc0, 0x0d, 0xa3, 0xa5,
	0x8e, 0xa5, 0x12, 0x46, 0xe0, 0xce, 0xef, 0x78, 0x78, 0xbd, 0x10, 0xa2, 0x28, 0x21, 0xa1, 0x92,
	0x25, 0x94, 0x73, 0x61, 0xa8, 0x61, 0x82, 0x2f, 0xf5, 0xe1, 0x8e, 0xbb, 0x64, 0xbb, 0x05, 0xf0,
	0x5d, 0xbd, 0xa0, 0x45, 0x01, 0x2a, 0x11, 0xd2, 0x29, 0x9a, 0xea, 0xe8, 0xab, 0x87, 0x36, 0x53,
	0xd0, 0x66, 0x4f, 0xca, 0xc7, 0x30, 0x65, 0x9c, 0x55, 0x24, 0xbe, 0x81, 0x90, 0x54, 0xe2, 0x08,
	0x32, 0x33, 0x61, 0x79, 0xe0, 0x11, 0xef, 0x76, 0x3b, 0x6d, 0x2f, 0x91, 0x83, 0x1c, 0x5f, 0x43,
	0xed, 0xac, 0x64, 0xc0, 0x1d, 0xeb, 0x3b, 0x76, 0xa3, 0x06, 0x0e, 0x72, 0xbc, 0x8d, 0x2e, 0x2c,
	0x49, 0x0d, 0x99, 0x02, 0x13, 0xb4, 0x9c, 0xe0, 0xff, 0x1a, 0x1c, 0x39, 0x0c, 0x5f, 0x45, 0x1b,
	0xd4, 0x9a, 0xd9, 0xc4, 0x2a, 0x16, 0xac, 0x38, 0x7e, 0xbd, 0xba, 0x1f, 0x2b, 0x56, 0xf9, 0x15,
	0xe4, 0x4c, 0x55, 0xe1, 0x56, 0x31, 0x1d, 0xac, 0x92, 0x56, 0xe5, 0x3f, 0x03, 0xc7, 0x8a, 0xe9,
	0xe8, 0x19, 0xc2, 0xfb, 0x0a, 0xa8, 0x81, 0xaa, 0x7b, 0x0a, 0x5a, 0x0a, 0xae, 0x01, 0xdf, 0x47,
	0x2b, 0xd5, 0x58, 0xae, 0xf0, 0x7f, 0xfd, 0xed, 0xb8, 0xb1, 0x68, 0xe3, 0xa4, 0xa9, 0x33, 0x44,
	0x2f, 0xd1, 0xe6, 0xf9, 0xc7, 0xbd, 0xb5, 0xa0, 0xcd, 0xbf, 0x46, 0x68, 0xf4, 0xac, 0x87, 0xf8,
	0xa5, 0x67, 0xff, 0xd4, 0x43, 0x97, 0x5c, 0xa8, 0x35, 0xb3, 0xfd, 0x9f, 0x45, 0xf0, 0x7b, 0x0f,
	0xad, 0xd5, 0x89, 0xf8, 0x0f, 0x35, 0x1b, 0x5d, 0xc2, 0x9b, 0x7f, 0x17, 0xd5, 0xe7, 0x8f, 0x76,
	0x3e, 0x7c, 0xfb, 0xfe, 0xc9, 0xbf, 0x15, 0x6d, 0x25, 0x52, 0x28, 0xc3, 0x68, 0x52, 0x99, 0x92,
	0x73, 0xa6, 0xa4, 0x00, 0x0e, 0x8a, 0x1a, 0x18, 0x7a, 0xdd, 0x47, 0xa7, 0xfe, 0xc7, 0xbd, 0xcf,
	0x3e, 0x9e, 0xa2, 0x8b, 0xcf, 0x9d, 0x9a, 0x8c, 0x40, 0x1d, 0xb3, 0x0c, 0x74, 0x74, 0x88, 0x82,
	0xa7, 0x4f, 0x5e, 0x8d, 0xc8, 0x54, 0x28, 0x32, 0x7f, 0x47, 0x38, 0x2c, 0x88, 0x5e, 0x72, 0xf8,
	0xf2, 0xcc, 0x18, 0xa9, 0x87, 0x89, 0x8b, 0xb0, 0xdc, 0xea, 0x38, 0x13, 0xf3, 0x30, 0x9c, 0x96,
	0xf4, 0x98, 0x59, 0x1d, 0x97, 0x34, 0xa3, 0xc6, 0x6a, 0xfb, 0xb0, 0x98, 0x53, 0x56, 0x56, 0x5c,
	0x7f, 0xf5, 0x4e, 0xdc, 0x8b, 0x7b, 0xd1, 0xfa, 0xb2, 0x54, 0xd7, 0xf7, 0xfc, 0x7e, 0x87, 0x4a,
	0x59, 0xb2, 0xcc, 0x7d, 0x8a, 0xc9, 0x91, 0x16, 0x7c, 0xd8, 0x40, 0xd2, 0x07, 0xa8, 0x35, 0xe8,
	0x0d, 0xf0, 0x00, 0x75, 0x53, 0x30, 0x56, 0x71, 0xc8, 0xc9, 0x62, 0x06, 0x9c, 0x98, 0x19, 0x10,
	0x05, 0x5a, 0x58, 0x95, 0x01, 0xc9, 0x05, 0x68, 0xc2, 0x85, 0x21, 0x70, 0xc2, 0xb4, 0x89, 0xf1,
	0x1a, 0x5a, 0xf9, 0xe2, 0x7b, 0xeb, 0xe9, 0x16, 0x6a, 0xdd, 0xeb, 0xf5, 0x70, 0x88, 0x3a, 0x63,
	0x0e, 0x27, 0x12, 0x32, 0x03, 0x39, 0x01, 0xa5, 0x84, 0x3a, 0x93, 0xbc, 0x0e, 0x51, 0x80, 0xf0,
	0x8b, 0xc3, 0x11, 0xa9, 0x5e, 0x4e, 0x35, 0x50, 0x1d, 0x8e, 0xfd, 0x0d, 0xef, 0xcd, 0x9a, 0xfb,
	0x3b, 0xee, 0xfe, 0x08, 0x00, 0x00, 0xff, 0xff, 0x89, 0x14, 0x6b, 0x56, 0x95, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RestAuthCredentialsClient is the client API for RestAuthCredentials service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RestAuthCredentialsClient interface {
	// Create new Credentials for rest backend
	Create(ctx context.Context, in *CreateRestRequest, opts ...grpc.CallOption) (*CreateRestResponse, error)
}

type restAuthCredentialsClient struct {
	cc *grpc.ClientConn
}

func NewRestAuthCredentialsClient(cc *grpc.ClientConn) RestAuthCredentialsClient {
	return &restAuthCredentialsClient{cc}
}

func (c *restAuthCredentialsClient) Create(ctx context.Context, in *CreateRestRequest, opts ...grpc.CallOption) (*CreateRestResponse, error) {
	out := new(CreateRestResponse)
	err := c.cc.Invoke(ctx, "/rest_credentials.RestAuthCredentials/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestAuthCredentialsServer is the server API for RestAuthCredentials service.
type RestAuthCredentialsServer interface {
	// Create new Credentials for rest backend
	Create(context.Context, *CreateRestRequest) (*CreateRestResponse, error)
}

// UnimplementedRestAuthCredentialsServer can be embedded to have forward compatible implementations.
type UnimplementedRestAuthCredentialsServer struct {
}

func (*UnimplementedRestAuthCredentialsServer) Create(ctx context.Context, req *CreateRestRequest) (*CreateRestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterRestAuthCredentialsServer(s *grpc.Server, srv RestAuthCredentialsServer) {
	s.RegisterService(&_RestAuthCredentials_serviceDesc, srv)
}

func _RestAuthCredentials_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestAuthCredentialsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rest_credentials.RestAuthCredentials/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestAuthCredentialsServer).Create(ctx, req.(*CreateRestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RestAuthCredentials_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rest_credentials.RestAuthCredentials",
	HandlerType: (*RestAuthCredentialsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RestAuthCredentials_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rest_credentials.proto",
}
