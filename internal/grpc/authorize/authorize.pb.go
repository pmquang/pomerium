// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authorize.proto

package authorize

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type IsAuthorizedRequest struct {
	// User Context
	//
	UserToken string `protobuf:"bytes,1,opt,name=user_token,json=userToken,proto3" json:"user_token,omitempty"`
	// Request Context
	//
	// Method specifies the HTTP method (GET, POST, PUT, etc.).
	RequestMethod string `protobuf:"bytes,2,opt,name=request_method,json=requestMethod,proto3" json:"request_method,omitempty"`
	// URL specifies either the URI being requested
	RequestUrl string `protobuf:"bytes,3,opt,name=request_url,json=requestUrl,proto3" json:"request_url,omitempty"`
	// host specifies the host on which the URL per RFC 7230, section 5.4
	RequestHost string `protobuf:"bytes,4,opt,name=request_host,json=requestHost,proto3" json:"request_host,omitempty"`
	// request_uri is the unmodified request-target of the
	// Request-Line (RFC 7230, Section 3.1.1) as sent by the client
	RequestRequestUri string `protobuf:"bytes,5,opt,name=request_request_uri,json=requestRequestUri,proto3" json:"request_request_uri,omitempty"`
	// RemoteAddr allows HTTP servers and other software to record
	// the network address that sent the request, usually for
	RequestRemoteAddr    string                                  `protobuf:"bytes,6,opt,name=request_remote_addr,json=requestRemoteAddr,proto3" json:"request_remote_addr,omitempty"`
	RequestHeaders       map[string]*IsAuthorizedRequest_Headers `protobuf:"bytes,7,rep,name=request_headers,json=requestHeaders,proto3" json:"request_headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *IsAuthorizedRequest) Reset()         { *m = IsAuthorizedRequest{} }
func (m *IsAuthorizedRequest) String() string { return proto.CompactTextString(m) }
func (*IsAuthorizedRequest) ProtoMessage()    {}
func (*IsAuthorizedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffbc3c71370bee9a, []int{0}
}

func (m *IsAuthorizedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthorizedRequest.Unmarshal(m, b)
}
func (m *IsAuthorizedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthorizedRequest.Marshal(b, m, deterministic)
}
func (m *IsAuthorizedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthorizedRequest.Merge(m, src)
}
func (m *IsAuthorizedRequest) XXX_Size() int {
	return xxx_messageInfo_IsAuthorizedRequest.Size(m)
}
func (m *IsAuthorizedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthorizedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthorizedRequest proto.InternalMessageInfo

func (m *IsAuthorizedRequest) GetUserToken() string {
	if m != nil {
		return m.UserToken
	}
	return ""
}

func (m *IsAuthorizedRequest) GetRequestMethod() string {
	if m != nil {
		return m.RequestMethod
	}
	return ""
}

func (m *IsAuthorizedRequest) GetRequestUrl() string {
	if m != nil {
		return m.RequestUrl
	}
	return ""
}

func (m *IsAuthorizedRequest) GetRequestHost() string {
	if m != nil {
		return m.RequestHost
	}
	return ""
}

func (m *IsAuthorizedRequest) GetRequestRequestUri() string {
	if m != nil {
		return m.RequestRequestUri
	}
	return ""
}

func (m *IsAuthorizedRequest) GetRequestRemoteAddr() string {
	if m != nil {
		return m.RequestRemoteAddr
	}
	return ""
}

func (m *IsAuthorizedRequest) GetRequestHeaders() map[string]*IsAuthorizedRequest_Headers {
	if m != nil {
		return m.RequestHeaders
	}
	return nil
}

// headers represents key-value pairs in an HTTP header; map[string][]string
type IsAuthorizedRequest_Headers struct {
	Value                []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAuthorizedRequest_Headers) Reset()         { *m = IsAuthorizedRequest_Headers{} }
func (m *IsAuthorizedRequest_Headers) String() string { return proto.CompactTextString(m) }
func (*IsAuthorizedRequest_Headers) ProtoMessage()    {}
func (*IsAuthorizedRequest_Headers) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffbc3c71370bee9a, []int{0, 0}
}

func (m *IsAuthorizedRequest_Headers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthorizedRequest_Headers.Unmarshal(m, b)
}
func (m *IsAuthorizedRequest_Headers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthorizedRequest_Headers.Marshal(b, m, deterministic)
}
func (m *IsAuthorizedRequest_Headers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthorizedRequest_Headers.Merge(m, src)
}
func (m *IsAuthorizedRequest_Headers) XXX_Size() int {
	return xxx_messageInfo_IsAuthorizedRequest_Headers.Size(m)
}
func (m *IsAuthorizedRequest_Headers) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthorizedRequest_Headers.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthorizedRequest_Headers proto.InternalMessageInfo

func (m *IsAuthorizedRequest_Headers) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type IsAuthorizedReply struct {
	IsValid              bool     `protobuf:"varint,1,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAuthorizedReply) Reset()         { *m = IsAuthorizedReply{} }
func (m *IsAuthorizedReply) String() string { return proto.CompactTextString(m) }
func (*IsAuthorizedReply) ProtoMessage()    {}
func (*IsAuthorizedReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffbc3c71370bee9a, []int{1}
}

func (m *IsAuthorizedReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthorizedReply.Unmarshal(m, b)
}
func (m *IsAuthorizedReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthorizedReply.Marshal(b, m, deterministic)
}
func (m *IsAuthorizedReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthorizedReply.Merge(m, src)
}
func (m *IsAuthorizedReply) XXX_Size() int {
	return xxx_messageInfo_IsAuthorizedReply.Size(m)
}
func (m *IsAuthorizedReply) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthorizedReply.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthorizedReply proto.InternalMessageInfo

func (m *IsAuthorizedReply) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

type IsAdminRequest struct {
	UserToken            string   `protobuf:"bytes,1,opt,name=user_token,json=userToken,proto3" json:"user_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAdminRequest) Reset()         { *m = IsAdminRequest{} }
func (m *IsAdminRequest) String() string { return proto.CompactTextString(m) }
func (*IsAdminRequest) ProtoMessage()    {}
func (*IsAdminRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffbc3c71370bee9a, []int{2}
}

func (m *IsAdminRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAdminRequest.Unmarshal(m, b)
}
func (m *IsAdminRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAdminRequest.Marshal(b, m, deterministic)
}
func (m *IsAdminRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAdminRequest.Merge(m, src)
}
func (m *IsAdminRequest) XXX_Size() int {
	return xxx_messageInfo_IsAdminRequest.Size(m)
}
func (m *IsAdminRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAdminRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsAdminRequest proto.InternalMessageInfo

func (m *IsAdminRequest) GetUserToken() string {
	if m != nil {
		return m.UserToken
	}
	return ""
}

type IsAdminReply struct {
	IsValid              bool     `protobuf:"varint,1,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAdminReply) Reset()         { *m = IsAdminReply{} }
func (m *IsAdminReply) String() string { return proto.CompactTextString(m) }
func (*IsAdminReply) ProtoMessage()    {}
func (*IsAdminReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ffbc3c71370bee9a, []int{3}
}

func (m *IsAdminReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAdminReply.Unmarshal(m, b)
}
func (m *IsAdminReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAdminReply.Marshal(b, m, deterministic)
}
func (m *IsAdminReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAdminReply.Merge(m, src)
}
func (m *IsAdminReply) XXX_Size() int {
	return xxx_messageInfo_IsAdminReply.Size(m)
}
func (m *IsAdminReply) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAdminReply.DiscardUnknown(m)
}

var xxx_messageInfo_IsAdminReply proto.InternalMessageInfo

func (m *IsAdminReply) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func init() {
	proto.RegisterType((*IsAuthorizedRequest)(nil), "authorize.IsAuthorizedRequest")
	proto.RegisterMapType((map[string]*IsAuthorizedRequest_Headers)(nil), "authorize.IsAuthorizedRequest.RequestHeadersEntry")
	proto.RegisterType((*IsAuthorizedRequest_Headers)(nil), "authorize.IsAuthorizedRequest.Headers")
	proto.RegisterType((*IsAuthorizedReply)(nil), "authorize.IsAuthorizedReply")
	proto.RegisterType((*IsAdminRequest)(nil), "authorize.IsAdminRequest")
	proto.RegisterType((*IsAdminReply)(nil), "authorize.IsAdminReply")
}

func init() { proto.RegisterFile("authorize.proto", fileDescriptor_ffbc3c71370bee9a) }

var fileDescriptor_ffbc3c71370bee9a = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x8e, 0xda, 0x30,
	0x10, 0x86, 0x1b, 0x52, 0x08, 0x19, 0x28, 0x14, 0x53, 0xa9, 0x26, 0x6a, 0x0b, 0x8d, 0xd4, 0x8a,
	0x5e, 0x52, 0x29, 0xbd, 0x54, 0x55, 0xa5, 0x8a, 0x43, 0x25, 0x38, 0xb4, 0x87, 0xa8, 0xed, 0xa5,
	0x87, 0x28, 0x2b, 0x5b, 0xc2, 0x22, 0x60, 0xd6, 0x76, 0x90, 0xb2, 0xef, 0xb2, 0xef, 0xb7, 0x8f,
	0xb1, 0x8a, 0xb1, 0xc3, 0xb2, 0x62, 0xd9, 0x3d, 0xe1, 0xf9, 0xe7, 0x9b, 0xdf, 0xe3, 0x5f, 0x04,
	0xfa, 0x59, 0xa1, 0x96, 0x5c, 0xb0, 0x2b, 0x1a, 0x6d, 0x05, 0x57, 0x1c, 0xf9, 0xb5, 0x10, 0xde,
	0xb8, 0x30, 0x5c, 0xc8, 0x99, 0xad, 0x49, 0x42, 0x2f, 0x0b, 0x2a, 0x15, 0x7a, 0x0b, 0x50, 0x48,
	0x2a, 0x52, 0xc5, 0x57, 0x74, 0x83, 0x9d, 0x89, 0x33, 0xf5, 0x13, 0xbf, 0x52, 0xfe, 0x54, 0x02,
	0xfa, 0x00, 0x3d, 0xb1, 0x27, 0xd3, 0x35, 0x55, 0x4b, 0x4e, 0x70, 0x43, 0x23, 0x2f, 0x8c, 0xfa,
	0x4b, 0x8b, 0x68, 0x0c, 0x1d, 0x8b, 0x15, 0x22, 0xc7, 0xae, 0x66, 0xc0, 0x48, 0x7f, 0x45, 0x8e,
	0xde, 0x43, 0xd7, 0x02, 0x4b, 0x2e, 0x15, 0x7e, 0xae, 0x09, 0x3b, 0x34, 0xe7, 0x52, 0xa1, 0x08,
	0x86, 0x16, 0x39, 0x78, 0x31, 0xdc, 0xd4, 0xe4, 0xc0, 0x48, 0x89, 0xb5, 0x64, 0xc7, 0xfc, 0x9a,
	0x2b, 0x9a, 0x66, 0x84, 0x08, 0xdc, 0xba, 0xc7, 0x57, 0x9d, 0x19, 0x21, 0x02, 0xfd, 0x87, 0x7e,
	0xbd, 0x02, 0xcd, 0x08, 0x15, 0x12, 0x7b, 0x13, 0x77, 0xda, 0x89, 0xe3, 0xe8, 0x90, 0xdb, 0x89,
	0x88, 0x22, 0xf3, 0x3b, 0xdf, 0x0f, 0xfd, 0xdc, 0x28, 0x51, 0x26, 0x36, 0x15, 0x23, 0x06, 0x63,
	0xf0, 0xcc, 0x11, 0xbd, 0x82, 0xe6, 0x2e, 0xcb, 0x0b, 0x8a, 0x9d, 0x89, 0x3b, 0xf5, 0x93, 0x7d,
	0x11, 0x30, 0x18, 0x9e, 0xf0, 0x41, 0x2f, 0xc1, 0x5d, 0xd1, 0xd2, 0xe4, 0x5e, 0x1d, 0xd1, 0x77,
	0x3b, 0x5e, 0x05, 0xdd, 0x89, 0x3f, 0x3e, 0xb2, 0x9c, 0x71, 0x33, 0xd7, 0x7c, 0x6b, 0x7c, 0x75,
	0xc2, 0x08, 0x06, 0xc7, 0xe4, 0x36, 0x2f, 0xd1, 0x08, 0xda, 0x4c, 0xa6, 0xbb, 0x2c, 0x67, 0x44,
	0xdf, 0xd6, 0x4e, 0x3c, 0x26, 0xff, 0x55, 0x65, 0xf8, 0x19, 0x7a, 0x0b, 0x39, 0x23, 0x6b, 0xb6,
	0x79, 0xda, 0x9f, 0x22, 0xfc, 0x04, 0xdd, 0x7a, 0xe0, 0xbc, 0x77, 0x7c, 0xed, 0x00, 0xd4, 0xab,
	0x08, 0xf4, 0x5b, 0x4f, 0xd6, 0xab, 0xa1, 0x77, 0xe7, 0x5f, 0x17, 0xbc, 0x79, 0xb0, 0xbf, 0xcd,
	0xcb, 0xf0, 0x19, 0xfa, 0x01, 0x9e, 0xd9, 0x04, 0x8d, 0x8e, 0xd1, 0x3b, 0xcf, 0x09, 0x5e, 0x9f,
	0x6a, 0x69, 0x83, 0x8b, 0x96, 0xfe, 0x50, 0xbe, 0xdc, 0x06, 0x00, 0x00, 0xff, 0xff, 0x72, 0x3a,
	0xa3, 0xe0, 0x3b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorizerClient is the client API for Authorizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizerClient interface {
	IsAuthorized(ctx context.Context, in *IsAuthorizedRequest, opts ...grpc.CallOption) (*IsAuthorizedReply, error)
	IsAdmin(ctx context.Context, in *IsAdminRequest, opts ...grpc.CallOption) (*IsAdminReply, error)
}

type authorizerClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizerClient(cc *grpc.ClientConn) AuthorizerClient {
	return &authorizerClient{cc}
}

func (c *authorizerClient) IsAuthorized(ctx context.Context, in *IsAuthorizedRequest, opts ...grpc.CallOption) (*IsAuthorizedReply, error) {
	out := new(IsAuthorizedReply)
	err := c.cc.Invoke(ctx, "/authorize.Authorizer/IsAuthorized", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizerClient) IsAdmin(ctx context.Context, in *IsAdminRequest, opts ...grpc.CallOption) (*IsAdminReply, error) {
	out := new(IsAdminReply)
	err := c.cc.Invoke(ctx, "/authorize.Authorizer/IsAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizerServer is the server API for Authorizer service.
type AuthorizerServer interface {
	IsAuthorized(context.Context, *IsAuthorizedRequest) (*IsAuthorizedReply, error)
	IsAdmin(context.Context, *IsAdminRequest) (*IsAdminReply, error)
}

// UnimplementedAuthorizerServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorizerServer struct {
}

func (*UnimplementedAuthorizerServer) IsAuthorized(ctx context.Context, req *IsAuthorizedRequest) (*IsAuthorizedReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAuthorized not implemented")
}
func (*UnimplementedAuthorizerServer) IsAdmin(ctx context.Context, req *IsAdminRequest) (*IsAdminReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAdmin not implemented")
}

func RegisterAuthorizerServer(s *grpc.Server, srv AuthorizerServer) {
	s.RegisterService(&_Authorizer_serviceDesc, srv)
}

func _Authorizer_IsAuthorized_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAuthorizedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizerServer).IsAuthorized(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorize.Authorizer/IsAuthorized",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizerServer).IsAuthorized(ctx, req.(*IsAuthorizedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorizer_IsAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizerServer).IsAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorize.Authorizer/IsAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizerServer).IsAdmin(ctx, req.(*IsAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorizer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authorize.Authorizer",
	HandlerType: (*AuthorizerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAuthorized",
			Handler:    _Authorizer_IsAuthorized_Handler,
		},
		{
			MethodName: "IsAdmin",
			Handler:    _Authorizer_IsAdmin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authorize.proto",
}
