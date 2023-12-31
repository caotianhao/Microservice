// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServClient is the client API for Serv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServClient interface {
	Exe(ctx context.Context, in *Cth, opts ...grpc.CallOption) (*Cth, error)
}

type servClient struct {
	cc grpc.ClientConnInterface
}

func NewServClient(cc grpc.ClientConnInterface) ServClient {
	return &servClient{cc}
}

func (c *servClient) Exe(ctx context.Context, in *Cth, opts ...grpc.CallOption) (*Cth, error) {
	out := new(Cth)
	err := c.cc.Invoke(ctx, "/pb.Serv/Exe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServServer is the server API for Serv service.
// All implementations must embed UnimplementedServServer
// for forward compatibility
type ServServer interface {
	Exe(context.Context, *Cth) (*Cth, error)
	mustEmbedUnimplementedServServer()
}

// UnimplementedServServer must be embedded to have forward compatible implementations.
type UnimplementedServServer struct {
}

func (UnimplementedServServer) Exe(context.Context, *Cth) (*Cth, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exe not implemented")
}
func (UnimplementedServServer) mustEmbedUnimplementedServServer() {}

// UnsafeServServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServServer will
// result in compilation errors.
type UnsafeServServer interface {
	mustEmbedUnimplementedServServer()
}

func RegisterServServer(s grpc.ServiceRegistrar, srv ServServer) {
	s.RegisterService(&Serv_ServiceDesc, srv)
}

func _Serv_Exe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServServer).Exe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Serv/Exe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServServer).Exe(ctx, req.(*Cth))
	}
	return interceptor(ctx, in, info, handler)
}

// Serv_ServiceDesc is the grpc.ServiceDesc for Serv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Serv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Serv",
	HandlerType: (*ServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exe",
			Handler:    _Serv_Exe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}
