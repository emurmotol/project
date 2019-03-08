// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

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

type GetByUsernameRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByUsernameRequest) Reset()         { *m = GetByUsernameRequest{} }
func (m *GetByUsernameRequest) String() string { return proto.CompactTextString(m) }
func (*GetByUsernameRequest) ProtoMessage()    {}
func (*GetByUsernameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_37757815a5c2e68a, []int{0}
}

func (m *GetByUsernameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByUsernameRequest.Unmarshal(m, b)
}
func (m *GetByUsernameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByUsernameRequest.Marshal(b, m, deterministic)
}
func (m *GetByUsernameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByUsernameRequest.Merge(m, src)
}
func (m *GetByUsernameRequest) XXX_Size() int {
	return xxx_messageInfo_GetByUsernameRequest.Size(m)
}
func (m *GetByUsernameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByUsernameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetByUsernameRequest proto.InternalMessageInfo

func (m *GetByUsernameRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type GetByUsernameReply struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByUsernameReply) Reset()         { *m = GetByUsernameReply{} }
func (m *GetByUsernameReply) String() string { return proto.CompactTextString(m) }
func (*GetByUsernameReply) ProtoMessage()    {}
func (*GetByUsernameReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_37757815a5c2e68a, []int{1}
}

func (m *GetByUsernameReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByUsernameReply.Unmarshal(m, b)
}
func (m *GetByUsernameReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByUsernameReply.Marshal(b, m, deterministic)
}
func (m *GetByUsernameReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByUsernameReply.Merge(m, src)
}
func (m *GetByUsernameReply) XXX_Size() int {
	return xxx_messageInfo_GetByUsernameReply.Size(m)
}
func (m *GetByUsernameReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByUsernameReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetByUsernameReply proto.InternalMessageInfo

func (m *GetByUsernameReply) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GetByUsernameReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type CreateUserRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Role                 string   `protobuf:"bytes,4,opt,name=Role,proto3" json:"Role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_37757815a5c2e68a, []int{2}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateUserRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type User struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
	Role                 string   `protobuf:"bytes,5,opt,name=Role,proto3" json:"Role,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt            string   `protobuf:"bytes,8,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_37757815a5c2e68a, []int{3}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *User) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

type CreateUserReply struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserReply) Reset()         { *m = CreateUserReply{} }
func (m *CreateUserReply) String() string { return proto.CompactTextString(m) }
func (*CreateUserReply) ProtoMessage()    {}
func (*CreateUserReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_37757815a5c2e68a, []int{4}
}

func (m *CreateUserReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserReply.Unmarshal(m, b)
}
func (m *CreateUserReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserReply.Marshal(b, m, deterministic)
}
func (m *CreateUserReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserReply.Merge(m, src)
}
func (m *CreateUserReply) XXX_Size() int {
	return xxx_messageInfo_CreateUserReply.Size(m)
}
func (m *CreateUserReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserReply.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserReply proto.InternalMessageInfo

func (m *CreateUserReply) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CreateUserReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type GetUserForAuthRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserForAuthRequest) Reset()         { *m = GetUserForAuthRequest{} }
func (m *GetUserForAuthRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserForAuthRequest) ProtoMessage()    {}
func (*GetUserForAuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_37757815a5c2e68a, []int{5}
}

func (m *GetUserForAuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserForAuthRequest.Unmarshal(m, b)
}
func (m *GetUserForAuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserForAuthRequest.Marshal(b, m, deterministic)
}
func (m *GetUserForAuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserForAuthRequest.Merge(m, src)
}
func (m *GetUserForAuthRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserForAuthRequest.Size(m)
}
func (m *GetUserForAuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserForAuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserForAuthRequest proto.InternalMessageInfo

func (m *GetUserForAuthRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type GetUserForAuthReply struct {
	User                 *User    `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserForAuthReply) Reset()         { *m = GetUserForAuthReply{} }
func (m *GetUserForAuthReply) String() string { return proto.CompactTextString(m) }
func (*GetUserForAuthReply) ProtoMessage()    {}
func (*GetUserForAuthReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_37757815a5c2e68a, []int{6}
}

func (m *GetUserForAuthReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserForAuthReply.Unmarshal(m, b)
}
func (m *GetUserForAuthReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserForAuthReply.Marshal(b, m, deterministic)
}
func (m *GetUserForAuthReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserForAuthReply.Merge(m, src)
}
func (m *GetUserForAuthReply) XXX_Size() int {
	return xxx_messageInfo_GetUserForAuthReply.Size(m)
}
func (m *GetUserForAuthReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserForAuthReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserForAuthReply proto.InternalMessageInfo

func (m *GetUserForAuthReply) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GetUserForAuthReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*GetByUsernameRequest)(nil), "pb.GetByUsernameRequest")
	proto.RegisterType((*GetByUsernameReply)(nil), "pb.GetByUsernameReply")
	proto.RegisterType((*CreateUserRequest)(nil), "pb.CreateUserRequest")
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*CreateUserReply)(nil), "pb.CreateUserReply")
	proto.RegisterType((*GetUserForAuthRequest)(nil), "pb.GetUserForAuthRequest")
	proto.RegisterType((*GetUserForAuthReply)(nil), "pb.GetUserForAuthReply")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_37757815a5c2e68a) }

var fileDescriptor_37757815a5c2e68a = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0x86, 0xd3, 0x52, 0xbe, 0xc6, 0x88, 0x71, 0x00, 0xad, 0x0d, 0x07, 0xd3, 0x93, 0x27, 0x0e,
	0x70, 0xf1, 0x5a, 0x2d, 0x22, 0x37, 0xd3, 0x84, 0xb3, 0x29, 0x61, 0x12, 0x49, 0x8a, 0x5d, 0xb7,
	0xdb, 0x18, 0xfe, 0xa3, 0x27, 0x7f, 0x91, 0x99, 0x5d, 0x28, 0xf2, 0xd1, 0x68, 0xb8, 0x75, 0xde,
	0x67, 0x66, 0xde, 0xdd, 0xe9, 0x2c, 0xb4, 0xf2, 0x8c, 0xe4, 0x6b, 0x2c, 0x16, 0x7d, 0x21, 0x53,
	0x95, 0xa2, 0x2d, 0x66, 0xfe, 0x00, 0x3a, 0x63, 0x52, 0x0f, 0xab, 0x69, 0x46, 0xf2, 0x3d, 0x5e,
	0x52, 0x44, 0x1f, 0x39, 0x65, 0x0a, 0x3d, 0x68, 0x6c, 0x24, 0xd7, 0xba, 0xb5, 0xee, 0x9a, 0x51,
	0x11, 0xfb, 0xcf, 0x80, 0x7b, 0x35, 0x22, 0x59, 0x61, 0x0f, 0x1c, 0x16, 0x74, 0xf6, 0xd9, 0xa0,
	0xd1, 0x17, 0xb3, 0x3e, 0xc7, 0x91, 0x56, 0xb1, 0x03, 0xd5, 0x91, 0x94, 0xa9, 0x74, 0x6d, 0xdd,
	0xcc, 0x04, 0x7e, 0x0e, 0x97, 0x8f, 0x92, 0x62, 0x45, 0x3a, 0xf3, 0x6f, 0x6b, 0xdd, 0x66, 0x19,
	0x2f, 0x92, 0xa2, 0x0d, 0x07, 0x5c, 0xf1, 0x12, 0x67, 0xd9, 0x67, 0x2a, 0xe7, 0x6e, 0xc5, 0x54,
	0x6c, 0x62, 0x44, 0x70, 0xa2, 0x34, 0x21, 0xd7, 0xd1, 0xba, 0xfe, 0xf6, 0xbf, 0x2d, 0x73, 0x56,
	0x6c, 0x81, 0x3d, 0x09, 0xd7, 0x26, 0xf6, 0x24, 0xdc, 0xb1, 0xb6, 0xcb, 0xac, 0x2b, 0x65, 0xd6,
	0x4e, 0x89, 0x75, 0x75, 0x6b, 0x8d, 0x3d, 0x68, 0x9a, 0x1b, 0xcf, 0x03, 0xe5, 0xd6, 0x34, 0xd8,
	0x0a, 0x4c, 0xa7, 0x62, 0xbe, 0xa6, 0x75, 0x43, 0x0b, 0x81, 0x69, 0x48, 0x09, 0x19, 0xda, 0x30,
	0xb4, 0x10, 0xfc, 0x11, 0x5c, 0xfc, 0x9e, 0xe5, 0xa9, 0xbf, 0x64, 0x08, 0xdd, 0x31, 0x29, 0x4e,
	0x78, 0x4a, 0x65, 0x90, 0xab, 0xb7, 0xff, 0x6c, 0xc4, 0x04, 0xda, 0xfb, 0x45, 0x27, 0xfa, 0x0f,
	0xbe, 0x2c, 0xa8, 0x33, 0x0e, 0xc4, 0x02, 0x03, 0x38, 0xdf, 0x59, 0x34, 0x74, 0xb9, 0xc5, 0xb1,
	0x7d, 0xf5, 0xae, 0x8e, 0x10, 0x3e, 0xc2, 0x3d, 0xc0, 0x76, 0x2a, 0xd8, 0xe5, 0xac, 0x83, 0x8d,
	0xf3, 0xda, 0xfb, 0x32, 0x57, 0x86, 0xd0, 0xda, 0xbd, 0x13, 0xde, 0xac, 0x3d, 0x0e, 0x87, 0xe3,
	0x5d, 0x1f, 0x43, 0x22, 0x59, 0xcd, 0x6a, 0xfa, 0xa9, 0x0d, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xfc, 0x14, 0xdb, 0x96, 0x7c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	GetByUsername(ctx context.Context, in *GetByUsernameRequest, opts ...grpc.CallOption) (*GetByUsernameReply, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserReply, error)
	GetUserForAuth(ctx context.Context, in *GetUserForAuthRequest, opts ...grpc.CallOption) (*GetUserForAuthReply, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetByUsername(ctx context.Context, in *GetByUsernameRequest, opts ...grpc.CallOption) (*GetByUsernameReply, error) {
	out := new(GetByUsernameReply)
	err := c.cc.Invoke(ctx, "/pb.User/GetByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserReply, error) {
	out := new(CreateUserReply)
	err := c.cc.Invoke(ctx, "/pb.User/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserForAuth(ctx context.Context, in *GetUserForAuthRequest, opts ...grpc.CallOption) (*GetUserForAuthReply, error) {
	out := new(GetUserForAuthReply)
	err := c.cc.Invoke(ctx, "/pb.User/GetUserForAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	GetByUsername(context.Context, *GetByUsernameRequest) (*GetByUsernameReply, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error)
	GetUserForAuth(context.Context, *GetUserForAuthRequest) (*GetUserForAuthReply, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/GetByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetByUsername(ctx, req.(*GetByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserForAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserForAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserForAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/GetUserForAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserForAuth(ctx, req.(*GetUserForAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByUsername",
			Handler:    _User_GetByUsername_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "GetUserForAuth",
			Handler:    _User_GetUserForAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}