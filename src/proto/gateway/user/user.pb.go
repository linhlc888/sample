// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package userpb is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	LoginRequest
	RegisterRequest
	Response
*/
package userpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LoginRequest struct {
	Email    string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterRequest struct {
	Email    string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Fullname string `protobuf:"bytes,2,opt,name=fullname" json:"fullname,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterRequest) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Response struct {
	Error   int32  `protobuf:"varint,1,opt,name=error" json:"error,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetError() int32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "user.LoginRequest")
	proto.RegisterType((*RegisterRequest)(nil), "user.RegisterRequest")
	proto.RegisterType((*Response)(nil), "user.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserService service

type UserServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Response, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/user.UserService/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/user.UserService/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	Login(context.Context, *LoginRequest) (*Response, error)
	Register(context.Context, *RegisterRequest) (*Response, error)
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
		FullMethod: "/user.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x95, 0x42, 0xa1, 0x18, 0x54, 0x84, 0xf9, 0xa3, 0x52, 0x75, 0x40, 0x99, 0x50, 0x87,
	0x5a, 0xc0, 0x52, 0xda, 0x05, 0x31, 0x30, 0x21, 0x21, 0x05, 0xb1, 0xb0, 0x20, 0x37, 0x1c, 0xae,
	0x25, 0xc7, 0x36, 0x3e, 0xa7, 0x15, 0x2b, 0xaf, 0xc0, 0x3b, 0xf0, 0x42, 0xbc, 0x02, 0x0f, 0x82,
	0xe2, 0xe0, 0x8a, 0xc2, 0xc0, 0x12, 0xe5, 0xf3, 0xdd, 0xfd, 0x3e, 0x7f, 0x67, 0x42, 0x4a, 0x04,
	0x37, 0xb0, 0xce, 0x78, 0x43, 0x57, 0xab, 0xff, 0x6e, 0x4f, 0x18, 0x23, 0x14, 0x30, 0x6e, 0x25,
	0xe3, 0x5a, 0x1b, 0xcf, 0xbd, 0x34, 0x1a, 0xeb, 0x9e, 0xf4, 0x82, 0x6c, 0x5d, 0x1b, 0x21, 0x75,
	0x06, 0xcf, 0x25, 0xa0, 0xa7, 0x7b, 0xa4, 0x09, 0x05, 0x97, 0xaa, 0x93, 0x1c, 0x25, 0xc7, 0x1b,
	0x59, 0x2d, 0x68, 0x97, 0xb4, 0x2c, 0x47, 0x9c, 0x1b, 0xf7, 0xd8, 0x69, 0x84, 0xc2, 0x42, 0xa7,
	0x0f, 0x64, 0x3b, 0x03, 0x21, 0xd1, 0x83, 0xfb, 0x17, 0xf2, 0x54, 0x2a, 0xa5, 0x79, 0x01, 0x11,
	0x12, 0xf5, 0x92, 0xc1, 0xca, 0x2f, 0x83, 0x11, 0x69, 0x65, 0x80, 0xd6, 0x68, 0x84, 0x40, 0x76,
	0xce, 0xb8, 0x40, 0x6e, 0x66, 0xb5, 0xa0, 0x1d, 0xb2, 0x5e, 0x00, 0x22, 0x17, 0x11, 0x1c, 0xe5,
	0xe9, 0x7b, 0x42, 0x36, 0xef, 0x10, 0xdc, 0x2d, 0xb8, 0x99, 0xcc, 0x81, 0x5e, 0x91, 0x66, 0x88,
	0x4b, 0xe9, 0x20, 0x2c, 0xea, 0x67, 0xf6, 0x6e, 0xbb, 0x3e, 0x8b, 0x66, 0xe9, 0xe1, 0xeb, 0xc7,
	0xe7, 0x5b, 0x63, 0x37, 0x6d, 0xb3, 0xd9, 0x09, 0xab, 0x4a, 0x4c, 0x55, 0xed, 0xa3, 0xa4, 0x4f,
	0x6f, 0xaa, 0x3b, 0xd5, 0xa1, 0xe9, 0x7e, 0x1c, 0x5b, 0x5a, 0xc2, 0x1f, 0x5a, 0x2f, 0xd0, 0x0e,
	0xd2, 0x9d, 0x05, 0xcd, 0x7d, 0x4f, 0x8c, 0x92, 0xfe, 0xe5, 0xf8, 0xfe, 0x5c, 0x48, 0x3f, 0x2d,
	0x27, 0x83, 0xdc, 0x14, 0x4c, 0x49, 0x3d, 0x55, 0xf9, 0x70, 0x38, 0x64, 0xc8, 0x0b, 0xab, 0x80,
	0xa1, 0xcb, 0x59, 0x78, 0x30, 0x26, 0xb8, 0x87, 0x39, 0x7f, 0x09, 0x84, 0x71, 0xf5, 0xb1, 0x93,
	0xc9, 0x5a, 0x28, 0x9d, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0x99, 0x0e, 0xf3, 0xa3, 0xfd, 0x01,
	0x00, 0x00,
}
