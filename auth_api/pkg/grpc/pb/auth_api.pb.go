// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth_api/pkg/grpc/pb/auth_api.proto

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
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ae2df39e8e406b7, []int{0}
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

type LoginReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReply) Reset()         { *m = LoginReply{} }
func (m *LoginReply) String() string { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()    {}
func (*LoginReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ae2df39e8e406b7, []int{1}
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

func init() {
	proto.RegisterType((*LoginRequest)(nil), "pb.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "pb.LoginReply")
}

func init() {
	proto.RegisterFile("auth_api/pkg/grpc/pb/auth_api.proto", fileDescriptor_5ae2df39e8e406b7)
}

var fileDescriptor_5ae2df39e8e406b7 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x2c, 0x2d, 0xc9,
	0x88, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0xc8, 0x4e, 0xd7, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x48,
	0xd2, 0x87, 0x09, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x31, 0x15, 0x24, 0x29, 0xf1, 0x71,
	0xf1, 0xf8, 0xe4, 0xa7, 0x67, 0xe6, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x28, 0xf1, 0x70,
	0x71, 0x41, 0xf9, 0x05, 0x39, 0x95, 0x46, 0x26, 0x5c, 0xec, 0x8e, 0xa5, 0x25, 0x19, 0x8e, 0x05,
	0x99, 0x42, 0x9a, 0x5c, 0xac, 0x60, 0x09, 0x21, 0x01, 0xbd, 0x82, 0x24, 0x3d, 0x64, 0x3d, 0x52,
	0x7c, 0x48, 0x22, 0x05, 0x39, 0x95, 0x49, 0x6c, 0x60, 0xe3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x10, 0x47, 0x6c, 0x38, 0x85, 0x00, 0x00, 0x00,
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

// AuthApiServer is the server API for AuthApi service.
type AuthApiServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
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

var _AuthApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AuthApi",
	HandlerType: (*AuthApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthApi_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_api/pkg/grpc/pb/auth_api.proto",
}
