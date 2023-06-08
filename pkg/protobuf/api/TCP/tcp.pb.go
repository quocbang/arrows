// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/tcp/tcp.proto

package tcp

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	authentication "github.com/quocbang/arrows/pkg/protobuf/authentication"
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

func init() { proto.RegisterFile("api/tcp/tcp.proto", fileDescriptor_f4f7d2875ee403d5) }

var fileDescriptor_f4f7d2875ee403d5 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2f, 0x49, 0x2e, 0x00, 0x61, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0xb6, 0xc4, 0xa2, 0xa2,
	0xfc, 0xf2, 0x62, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0x7d, 0x90, 0x8a, 0xc4, 0xbc,
	0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0x2a, 0x29, 0x69, 0xa8, 0x2c, 0x98,
	0x97, 0x54, 0x9a, 0xa6, 0x9f, 0x9a, 0x5b, 0x50, 0x52, 0x09, 0x95, 0x94, 0x4c, 0x2c, 0x2d, 0xc9,
	0x48, 0xcd, 0x2b, 0xc9, 0x4c, 0x06, 0xeb, 0xd1, 0x2f, 0x2d, 0x4e, 0x2d, 0x82, 0x48, 0x19, 0x05,
	0x72, 0x71, 0x86, 0x38, 0x07, 0x38, 0x82, 0xad, 0x10, 0x72, 0xe1, 0x62, 0xf5, 0xc9, 0x4f, 0xcf,
	0xcc, 0x13, 0x12, 0xd1, 0x83, 0x58, 0xaa, 0x07, 0xe6, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x48, 0x09, 0xa1, 0x89, 0x16, 0xe4, 0x54, 0x2a, 0x09, 0x36, 0x5d, 0x7e, 0x32, 0x99, 0x89, 0x5b,
	0x89, 0x4d, 0x3f, 0x07, 0x24, 0x68, 0xc5, 0xa8, 0xe5, 0x64, 0x18, 0xa5, 0x9f, 0x9e, 0x59, 0x92,
	0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x5f, 0x58, 0x9a, 0x9f, 0x9c, 0x94, 0x98, 0x97, 0xae,
	0x0f, 0xd1, 0xab, 0x5f, 0x90, 0x9d, 0x8e, 0x70, 0x24, 0xd4, 0xb7, 0x49, 0x6c, 0x60, 0x11, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0xdb, 0x4f, 0xaa, 0xff, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TCPArrowsClient is the client API for TCPArrows service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TCPArrowsClient interface {
	// Login to access user agant.
	//
	// Required:
	// - UserName
	// - Password
	Login(ctx context.Context, in *authentication.LoginRequest, opts ...grpc.CallOption) (*authentication.LoginReply, error)
}

type tCPArrowsClient struct {
	cc *grpc.ClientConn
}

func NewTCPArrowsClient(cc *grpc.ClientConn) TCPArrowsClient {
	return &tCPArrowsClient{cc}
}

func (c *tCPArrowsClient) Login(ctx context.Context, in *authentication.LoginRequest, opts ...grpc.CallOption) (*authentication.LoginReply, error) {
	out := new(authentication.LoginReply)
	err := c.cc.Invoke(ctx, "/arrows.TCPArrows/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TCPArrowsServer is the server API for TCPArrows service.
type TCPArrowsServer interface {
	// Login to access user agant.
	//
	// Required:
	// - UserName
	// - Password
	Login(context.Context, *authentication.LoginRequest) (*authentication.LoginReply, error)
}

func RegisterTCPArrowsServer(s *grpc.Server, srv TCPArrowsServer) {
	s.RegisterService(&_TCPArrows_serviceDesc, srv)
}

func _TCPArrows_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(authentication.LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TCPArrowsServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arrows.TCPArrows/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TCPArrowsServer).Login(ctx, req.(*authentication.LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TCPArrows_serviceDesc = grpc.ServiceDesc{
	ServiceName: "arrows.TCPArrows",
	HandlerType: (*TCPArrowsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _TCPArrows_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/tcp/tcp.proto",
}
