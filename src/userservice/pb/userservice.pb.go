// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userservice.proto

package userservice

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type MutationReply struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutationReply) Reset()         { *m = MutationReply{} }
func (m *MutationReply) String() string { return proto.CompactTextString(m) }
func (*MutationReply) ProtoMessage()    {}
func (*MutationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{0}
}

func (m *MutationReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutationReply.Unmarshal(m, b)
}
func (m *MutationReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutationReply.Marshal(b, m, deterministic)
}
func (m *MutationReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutationReply.Merge(m, src)
}
func (m *MutationReply) XXX_Size() int {
	return xxx_messageInfo_MutationReply.Size(m)
}
func (m *MutationReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MutationReply.DiscardUnknown(m)
}

var xxx_messageInfo_MutationReply proto.InternalMessageInfo

func (m *MutationReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt            int64    `protobuf:"varint,15,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{1}
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

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

type Token struct {
	AccessToken          string               `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	ExpiresIn            *timestamp.Timestamp `protobuf:"bytes,15,opt,name=expiresIn,proto3" json:"expiresIn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{2}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Token) GetExpiresIn() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiresIn
	}
	return nil
}

type LoginRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{3}
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

func (m *LoginRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type LoginReply struct {
	Token                *Token   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReply) Reset()         { *m = LoginReply{} }
func (m *LoginReply) String() string { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()    {}
func (*LoginReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{4}
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

func (m *LoginReply) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type ProfileReply struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	PostCount            int32    `protobuf:"varint,2,opt,name=postCount,proto3" json:"postCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfileReply) Reset()         { *m = ProfileReply{} }
func (m *ProfileReply) String() string { return proto.CompactTextString(m) }
func (*ProfileReply) ProtoMessage()    {}
func (*ProfileReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_68a7ca558839fd2b, []int{5}
}

func (m *ProfileReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileReply.Unmarshal(m, b)
}
func (m *ProfileReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileReply.Marshal(b, m, deterministic)
}
func (m *ProfileReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileReply.Merge(m, src)
}
func (m *ProfileReply) XXX_Size() int {
	return xxx_messageInfo_ProfileReply.Size(m)
}
func (m *ProfileReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileReply.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileReply proto.InternalMessageInfo

func (m *ProfileReply) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ProfileReply) GetPostCount() int32 {
	if m != nil {
		return m.PostCount
	}
	return 0
}

func init() {
	proto.RegisterType((*MutationReply)(nil), "userservice.MutationReply")
	proto.RegisterType((*User)(nil), "userservice.User")
	proto.RegisterType((*Token)(nil), "userservice.Token")
	proto.RegisterType((*LoginRequest)(nil), "userservice.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "userservice.LoginReply")
	proto.RegisterType((*ProfileReply)(nil), "userservice.ProfileReply")
}

func init() { proto.RegisterFile("userservice.proto", fileDescriptor_68a7ca558839fd2b) }

var fileDescriptor_68a7ca558839fd2b = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x5d, 0xaa, 0xd3, 0x40,
	0x14, 0x26, 0x69, 0xd3, 0xda, 0x93, 0xaa, 0x74, 0x1e, 0x6a, 0x8d, 0x05, 0xcb, 0x80, 0x58, 0x7c,
	0x48, 0xa0, 0x82, 0x88, 0x6f, 0x22, 0x82, 0x05, 0x45, 0x99, 0xd6, 0x05, 0xa4, 0xe9, 0x69, 0x19,
	0x4c, 0x66, 0xc6, 0x99, 0x89, 0xd8, 0x57, 0xb7, 0xe0, 0x2a, 0x5c, 0x8f, 0x5b, 0x70, 0x21, 0x92,
	0x49, 0xda, 0xa4, 0xf7, 0xde, 0xb7, 0x9c, 0xef, 0xe7, 0xe4, 0xfb, 0x4e, 0x02, 0x93, 0xd2, 0xa0,
	0x36, 0xa8, 0x7f, 0xf0, 0x0c, 0x63, 0xa5, 0xa5, 0x95, 0x24, 0xec, 0x40, 0xd1, 0x93, 0xa3, 0x94,
	0xc7, 0x1c, 0x13, 0x47, 0xed, 0xca, 0x43, 0x82, 0x85, 0xb2, 0xa7, 0x5a, 0x19, 0x3d, 0xbd, 0x49,
	0x5a, 0x5e, 0xa0, 0xb1, 0x69, 0xa1, 0x1a, 0xc1, 0xbc, 0x11, 0xa4, 0x8a, 0x27, 0xa9, 0x10, 0xd2,
	0xa6, 0x96, 0x4b, 0x61, 0x6a, 0x96, 0x3e, 0x87, 0xfb, 0x9f, 0xca, 0x1a, 0x62, 0xa8, 0xf2, 0x13,
	0x99, 0xc2, 0x40, 0xa3, 0x29, 0x73, 0x3b, 0xf3, 0x16, 0xde, 0xf2, 0x1e, 0x6b, 0x26, 0xfa, 0x01,
	0xfa, 0x5f, 0x0d, 0x6a, 0xf2, 0x00, 0x7c, 0xbe, 0x77, 0xdc, 0x88, 0xf9, 0x7c, 0x4f, 0x08, 0xf4,
	0x45, 0x5a, 0xe0, 0xcc, 0x77, 0x88, 0x7b, 0x26, 0x73, 0x18, 0x65, 0x1a, 0x53, 0x8b, 0xfb, 0xb7,
	0x76, 0xf6, 0x70, 0xe1, 0x2d, 0x7b, 0xac, 0x05, 0x68, 0x06, 0xc1, 0x56, 0x7e, 0x43, 0x41, 0x16,
	0x10, 0xa6, 0x59, 0x86, 0xc6, 0xb8, 0xb1, 0xd9, 0xd9, 0x85, 0xc8, 0x6b, 0x18, 0xe1, 0x4f, 0xc5,
	0x35, 0x9a, 0xb5, 0x70, 0x8b, 0xc2, 0x55, 0x14, 0xd7, 0x7d, 0xe2, 0x73, 0xe1, 0x78, 0x7b, 0x2e,
	0xcc, 0x5a, 0x31, 0xa5, 0x30, 0xfe, 0x28, 0x8f, 0x5c, 0x30, 0xfc, 0x5e, 0xa2, 0xb1, 0x97, 0x98,
	0x5e, 0x1b, 0x93, 0xbe, 0x02, 0x68, 0x34, 0x55, 0xf1, 0x25, 0x04, 0xf6, 0x92, 0x23, 0x5c, 0x91,
	0xb8, 0xfb, 0x55, 0x5c, 0x1c, 0x56, 0x0b, 0xe8, 0x06, 0xc6, 0x5f, 0xb4, 0x3c, 0xf0, 0x1c, 0x6b,
	0xe7, 0x33, 0xe8, 0x57, 0xda, 0xc6, 0x38, 0xb9, 0x32, 0x56, 0x37, 0x63, 0x8e, 0xae, 0xae, 0xa2,
	0xa4, 0xb1, 0xef, 0x64, 0x29, 0xac, 0x3b, 0x57, 0xc0, 0x5a, 0x60, 0xf5, 0xc7, 0x83, 0xb0, 0x12,
	0x6f, 0x6a, 0x23, 0xf9, 0x0c, 0x81, 0x0b, 0x47, 0x1e, 0x5f, 0xed, 0xeb, 0x96, 0x8a, 0x1e, 0xdd,
	0x45, 0xa9, 0xfc, 0x44, 0xc9, 0xaf, 0xbf, 0xff, 0x7e, 0xfb, 0x63, 0x3a, 0x4c, 0x5c, 0x62, 0xf3,
	0xc6, 0x7b, 0x41, 0xd6, 0x30, 0x6c, 0x52, 0x93, 0xe9, 0xad, 0x1b, 0xbe, 0xaf, 0xfe, 0xa8, 0xe8,
	0xfa, 0x55, 0xdd, 0x8e, 0x34, 0x74, 0x1b, 0x03, 0xd2, 0x4b, 0x0a, 0xdc, 0x0d, 0x9c, 0xef, 0xe5,
	0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xaf, 0x14, 0xaf, 0xb9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	Profile(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProfileReply, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/userservice.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Profile(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, "/userservice.UserService/Profile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Profile(context.Context, *empty.Empty) (*ProfileReply, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Profile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Profile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userservice.UserService/Profile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Profile(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "userservice.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "Profile",
			Handler:    _UserService_Profile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userservice.proto",
}