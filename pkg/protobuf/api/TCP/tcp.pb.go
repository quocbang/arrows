// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/TCP/tcp.proto

package TCP

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	authentication "github.com/quocbang/arrows/pkg/protobuf/authentication"
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

func init() { proto.RegisterFile("api/TCP/tcp.proto", fileDescriptor_680e5dbf1d3ddac1) }

var fileDescriptor_680e5dbf1d3ddac1 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8e, 0xb1, 0x0a, 0xc2, 0x30,
	0x14, 0x45, 0x5d, 0x2c, 0xd8, 0xcd, 0xe0, 0x62, 0x47, 0x3f, 0x20, 0x0f, 0xed, 0x0f, 0xa8, 0x5d,
	0x1d, 0x8a, 0x74, 0x72, 0x4b, 0x42, 0x4c, 0x83, 0x9a, 0x97, 0xa6, 0xef, 0x21, 0xfe, 0xbd, 0x90,
	0x38, 0x39, 0xde, 0xc3, 0xb9, 0xdc, 0x5b, 0xaf, 0x55, 0xf4, 0x30, 0x74, 0x3d, 0x90, 0x89, 0x32,
	0x26, 0x24, 0x14, 0x95, 0x4a, 0x09, 0xdf, 0x73, 0xb3, 0x55, 0x4c, 0xa3, 0x0d, 0xe4, 0x8d, 0x22,
	0x8f, 0x01, 0x78, 0xb6, 0xa9, 0x28, 0x87, 0x63, 0xbd, 0x1a, 0xba, 0xfe, 0x94, 0x3d, 0xd1, 0xd6,
	0xcb, 0x0b, 0x3a, 0x1f, 0xc4, 0x46, 0x96, 0xa6, 0xcc, 0xf1, 0x6a, 0x27, 0xb6, 0x33, 0x35, 0xe2,
	0x8f, 0xc6, 0xe7, 0x67, 0xb7, 0x38, 0xef, 0x6f, 0xe0, 0x3c, 0x8d, 0xac, 0xa5, 0xc1, 0x17, 0x4c,
	0x8c, 0x46, 0xab, 0xe0, 0xa0, 0xa8, 0x10, 0x1f, 0x0e, 0xf2, 0x92, 0xe6, 0x3b, 0xfc, 0x1e, 0xea,
	0x2a, 0x93, 0xf6, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x11, 0x2e, 0x1c, 0xe9, 0xb3, 0x00, 0x00, 0x00,
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
	Metadata: "api/TCP/tcp.proto",
}
