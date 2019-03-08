// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth_api.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5490162f5f4e8e93, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginReply struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	TokenType            string   `protobuf:"bytes,2,opt,name=TokenType,proto3" json:"TokenType,omitempty"`
	ExpiresAt            int64    `protobuf:"varint,3,opt,name=ExpiresAt,proto3" json:"ExpiresAt,omitempty"`
	Error                string   `protobuf:"bytes,4,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReply) Reset()         { *m = LoginReply{} }
func (m *LoginReply) String() string { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()    {}
func (*LoginReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5490162f5f4e8e93, []int{1}
}

func (m *LoginReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReply.Unmarshal(m, b)
}
func (m *LoginReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReply.Marshal(b, m, deterministic)
}
func (m *LoginReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReply.Merge(m, src)
}
func (m *LoginReply) XXX_Size() int {
	return xxx_messageInfo_LoginReply.Size(m)
}
func (m *LoginReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReply.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReply proto.InternalMessageInfo

func (m *LoginReply) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *LoginReply) GetTokenType() string {
	if m != nil {
		return m.TokenType
	}
	return ""
}

func (m *LoginReply) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func (m *LoginReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type RestrictedRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestrictedRequest) Reset()         { *m = RestrictedRequest{} }
func (m *RestrictedRequest) String() string { return proto.CompactTextString(m) }
func (*RestrictedRequest) ProtoMessage()    {}
func (*RestrictedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5490162f5f4e8e93, []int{2}
}

func (m *RestrictedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestrictedRequest.Unmarshal(m, b)
}
func (m *RestrictedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestrictedRequest.Marshal(b, m, deterministic)
}
func (m *RestrictedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestrictedRequest.Merge(m, src)
}
func (m *RestrictedRequest) XXX_Size() int {
	return xxx_messageInfo_RestrictedRequest.Size(m)
}
func (m *RestrictedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RestrictedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RestrictedRequest proto.InternalMessageInfo

type Claims struct {
	Audience             string   `protobuf:"bytes,1,opt,name=Audience,proto3" json:"Audience,omitempty"`
	ExpiresAt            int64    `protobuf:"varint,2,opt,name=ExpiresAt,proto3" json:"ExpiresAt,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=Id,proto3" json:"Id,omitempty"`
	IssuedAt             int64    `protobuf:"varint,4,opt,name=IssuedAt,proto3" json:"IssuedAt,omitempty"`
	Issuer               string   `protobuf:"bytes,5,opt,name=Issuer,proto3" json:"Issuer,omitempty"`
	NotBefore            int64    `protobuf:"varint,6,opt,name=NotBefore,proto3" json:"NotBefore,omitempty"`
	Subject              string   `protobuf:"bytes,7,opt,name=Subject,proto3" json:"Subject,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Claims) Reset()         { *m = Claims{} }
func (m *Claims) String() string { return proto.CompactTextString(m) }
func (*Claims) ProtoMessage()    {}
func (*Claims) Descriptor() ([]byte, []int) {
	return fileDescriptor_5490162f5f4e8e93, []int{3}
}

func (m *Claims) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Claims.Unmarshal(m, b)
}
func (m *Claims) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Claims.Marshal(b, m, deterministic)
}
func (m *Claims) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Claims.Merge(m, src)
}
func (m *Claims) XXX_Size() int {
	return xxx_messageInfo_Claims.Size(m)
}
func (m *Claims) XXX_DiscardUnknown() {
	xxx_messageInfo_Claims.DiscardUnknown(m)
}

var xxx_messageInfo_Claims proto.InternalMessageInfo

func (m *Claims) GetAudience() string {
	if m != nil {
		return m.Audience
	}
	return ""
}

func (m *Claims) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func (m *Claims) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Claims) GetIssuedAt() int64 {
	if m != nil {
		return m.IssuedAt
	}
	return 0
}

func (m *Claims) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *Claims) GetNotBefore() int64 {
	if m != nil {
		return m.NotBefore
	}
	return 0
}

func (m *Claims) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

type RestrictedReply struct {
	Claims               *Claims  `protobuf:"bytes,1,opt,name=Claims,proto3" json:"Claims,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestrictedReply) Reset()         { *m = RestrictedReply{} }
func (m *RestrictedReply) String() string { return proto.CompactTextString(m) }
func (*RestrictedReply) ProtoMessage()    {}
func (*RestrictedReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5490162f5f4e8e93, []int{4}
}

func (m *RestrictedReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestrictedReply.Unmarshal(m, b)
}
func (m *RestrictedReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestrictedReply.Marshal(b, m, deterministic)
}
func (m *RestrictedReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestrictedReply.Merge(m, src)
}
func (m *RestrictedReply) XXX_Size() int {
	return xxx_messageInfo_RestrictedReply.Size(m)
}
func (m *RestrictedReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RestrictedReply.DiscardUnknown(m)
}

var xxx_messageInfo_RestrictedReply proto.InternalMessageInfo

func (m *RestrictedReply) GetClaims() *Claims {
	if m != nil {
		return m.Claims
	}
	return nil
}

func (m *RestrictedReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type HealthCheckRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5490162f5f4e8e93, []int{5}
}

func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (m *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(m, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

type HealthCheckReply struct {
	Status               string   `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckReply) Reset()         { *m = HealthCheckReply{} }
func (m *HealthCheckReply) String() string { return proto.CompactTextString(m) }
func (*HealthCheckReply) ProtoMessage()    {}
func (*HealthCheckReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5490162f5f4e8e93, []int{6}
}

func (m *HealthCheckReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckReply.Unmarshal(m, b)
}
func (m *HealthCheckReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckReply.Marshal(b, m, deterministic)
}
func (m *HealthCheckReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckReply.Merge(m, src)
}
func (m *HealthCheckReply) XXX_Size() int {
	return xxx_messageInfo_HealthCheckReply.Size(m)
}
func (m *HealthCheckReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckReply.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckReply proto.InternalMessageInfo

func (m *HealthCheckReply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *HealthCheckReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "pb.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "pb.LoginReply")
	proto.RegisterType((*RestrictedRequest)(nil), "pb.RestrictedRequest")
	proto.RegisterType((*Claims)(nil), "pb.Claims")
	proto.RegisterType((*RestrictedReply)(nil), "pb.RestrictedReply")
	proto.RegisterType((*HealthCheckRequest)(nil), "pb.HealthCheckRequest")
	proto.RegisterType((*HealthCheckReply)(nil), "pb.HealthCheckReply")
}

func init() { proto.RegisterFile("auth_api.proto", fileDescriptor_5490162f5f4e8e93) }

var fileDescriptor_5490162f5f4e8e93 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x0f, 0xd2, 0x40,
	0x10, 0x4d, 0x8b, 0x14, 0x19, 0x0c, 0xe2, 0x82, 0x64, 0xd3, 0x78, 0x20, 0x3d, 0xe1, 0x85, 0x03,
	0x5e, 0x4c, 0xbc, 0x58, 0x09, 0x46, 0xa2, 0x31, 0xa6, 0xe0, 0xd9, 0xf4, 0x63, 0xb4, 0x95, 0xd2,
	0x5d, 0x77, 0xb7, 0x51, 0x6e, 0xfe, 0x1f, 0xff, 0x83, 0xbf, 0xcd, 0xec, 0xf6, 0x83, 0x02, 0xde,
	0xfa, 0xde, 0xec, 0xcc, 0x9b, 0xf7, 0x3a, 0x30, 0x0e, 0x4b, 0x95, 0x7e, 0x09, 0x79, 0xb6, 0xe2,
	0x82, 0x29, 0x46, 0x6c, 0x1e, 0x79, 0x6f, 0xe1, 0xd1, 0x07, 0xf6, 0x2d, 0x2b, 0x02, 0xfc, 0x51,
	0xa2, 0x54, 0xc4, 0x85, 0x87, 0x9f, 0x25, 0x8a, 0x22, 0x3c, 0x21, 0xb5, 0x16, 0xd6, 0x72, 0x18,
	0xb4, 0x58, 0xd7, 0x3e, 0x85, 0x52, 0xfe, 0x64, 0x22, 0xa1, 0x76, 0x55, 0x6b, 0xb0, 0xf7, 0xdb,
	0x02, 0xa8, 0x07, 0xf1, 0xfc, 0x4c, 0x16, 0x30, 0xf2, 0xe3, 0x18, 0xa5, 0x3c, 0xb0, 0x23, 0x16,
	0xf5, 0xa4, 0x2e, 0x45, 0x9e, 0xc1, 0xd0, 0x7c, 0x1c, 0xce, 0x1c, 0xeb, 0x69, 0x17, 0x42, 0x57,
	0xb7, 0xbf, 0x78, 0x26, 0x50, 0xfa, 0x8a, 0xf6, 0x16, 0xd6, 0xb2, 0x17, 0x5c, 0x08, 0x32, 0x83,
	0xfe, 0x56, 0x08, 0x26, 0xe8, 0x03, 0xd3, 0x57, 0x01, 0x6f, 0x0a, 0x4f, 0x02, 0x94, 0x4a, 0x64,
	0xb1, 0xc2, 0xa4, 0xf6, 0xe3, 0xfd, 0xb5, 0xc0, 0xd9, 0xe4, 0x61, 0x76, 0x92, 0x7a, 0x7d, 0xbf,
	0x4c, 0x32, 0x2c, 0xe2, 0xd6, 0x5a, 0x83, 0xaf, 0xf5, 0xec, 0x5b, 0xbd, 0x31, 0xd8, 0xbb, 0xc4,
	0xac, 0x31, 0x0c, 0xec, 0x5d, 0xa2, 0x27, 0xed, 0xa4, 0x2c, 0x31, 0xf1, 0x95, 0x59, 0xa1, 0x17,
	0xb4, 0x98, 0xcc, 0xc1, 0x31, 0xdf, 0x82, 0xf6, 0xcd, 0xfb, 0x1a, 0x69, 0x85, 0x8f, 0x4c, 0xbd,
	0xc1, 0xaf, 0x4c, 0x20, 0x75, 0x2a, 0x85, 0x96, 0x20, 0x14, 0x06, 0xfb, 0x32, 0xfa, 0x8e, 0xb1,
	0xa2, 0x03, 0xd3, 0xd6, 0x40, 0xef, 0x3d, 0x3c, 0xee, 0xba, 0xd2, 0xe1, 0x7a, 0x8d, 0x25, 0x63,
	0x63, 0xb4, 0x86, 0x15, 0x8f, 0x56, 0x15, 0x13, 0x34, 0x66, 0xdb, 0x88, 0xec, 0x6e, 0x44, 0x33,
	0x20, 0xef, 0x30, 0xcc, 0x55, 0xba, 0x49, 0x31, 0x3e, 0x36, 0x19, 0xbd, 0x86, 0xc9, 0x15, 0xab,
	0x35, 0xe6, 0xe0, 0xec, 0x55, 0xa8, 0x4a, 0x59, 0x47, 0x55, 0xa3, 0xff, 0xcf, 0x5d, 0xff, 0xb1,
	0x60, 0xe0, 0x97, 0x2a, 0xf5, 0x79, 0x46, 0x9e, 0x43, 0xdf, 0x1c, 0x02, 0x99, 0xe8, 0xb5, 0xba,
	0xc7, 0xe5, 0x8e, 0x3b, 0x8c, 0x16, 0x79, 0x09, 0x70, 0xf1, 0x46, 0x9e, 0xea, 0xea, 0xdd, 0x1f,
	0x74, 0xa7, 0xb7, 0xb4, 0xee, 0x7c, 0x05, 0xa3, 0xce, 0xca, 0x64, 0xae, 0xdf, 0xdc, 0x3b, 0x73,
	0x67, 0x77, 0x3c, 0xcf, 0xcf, 0x91, 0x63, 0xce, 0xff, 0xc5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xe8, 0xff, 0x54, 0x12, 0x10, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthApiClient is the client API for AuthApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthApiClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	Restricted(ctx context.Context, in *RestrictedRequest, opts ...grpc.CallOption) (*RestrictedReply, error)
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckReply, error)
}

type authApiClient struct {
	cc *grpc.ClientConn
}

func NewAuthApiClient(cc *grpc.ClientConn) AuthApiClient {
	return &authApiClient{cc}
}

func (c *authApiClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/pb.AuthApi/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authApiClient) Restricted(ctx context.Context, in *RestrictedRequest, opts ...grpc.CallOption) (*RestrictedReply, error) {
	out := new(RestrictedReply)
	err := c.cc.Invoke(ctx, "/pb.AuthApi/Restricted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authApiClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckReply, error) {
	out := new(HealthCheckReply)
	err := c.cc.Invoke(ctx, "/pb.AuthApi/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthApiServer is the server API for AuthApi service.
type AuthApiServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Restricted(context.Context, *RestrictedRequest) (*RestrictedReply, error)
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckReply, error)
}

func RegisterAuthApiServer(s *grpc.Server, srv AuthApiServer) {
	s.RegisterService(&_AuthApi_serviceDesc, srv)
}

func _AuthApi_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthApiServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthApi/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthApiServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthApi_Restricted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestrictedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthApiServer).Restricted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthApi/Restricted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthApiServer).Restricted(ctx, req.(*RestrictedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthApi_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthApiServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthApi/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthApiServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AuthApi",
	HandlerType: (*AuthApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthApi_Login_Handler,
		},
		{
			MethodName: "Restricted",
			Handler:    _AuthApi_Restricted_Handler,
		},
		{
			MethodName: "HealthCheck",
			Handler:    _AuthApi_HealthCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_api.proto",
}