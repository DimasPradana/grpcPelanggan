// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: proto/pelanggan.proto

package grpcPelanggan

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

// GetPelangganClient is the client API for GetPelanggan service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetPelangganClient interface {
	// rpc getPelanggan (PelangganRequest) returns (PelangganResponse) {};
	GetPelangganApi(ctx context.Context, in *PelangganRequest, opts ...grpc.CallOption) (*PelangganResponse, error)
	GetAllPelangganApi(ctx context.Context, in *PelangganRequest, opts ...grpc.CallOption) (*AllPelangganResponse, error)
}

type getPelangganClient struct {
	cc grpc.ClientConnInterface
}

func NewGetPelangganClient(cc grpc.ClientConnInterface) GetPelangganClient {
	return &getPelangganClient{cc}
}

func (c *getPelangganClient) GetPelangganApi(ctx context.Context, in *PelangganRequest, opts ...grpc.CallOption) (*PelangganResponse, error) {
	out := new(PelangganResponse)
	err := c.cc.Invoke(ctx, "/grpcPelanggan.getPelanggan/getPelangganApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getPelangganClient) GetAllPelangganApi(ctx context.Context, in *PelangganRequest, opts ...grpc.CallOption) (*AllPelangganResponse, error) {
	out := new(AllPelangganResponse)
	err := c.cc.Invoke(ctx, "/grpcPelanggan.getPelanggan/getAllPelangganApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetPelangganServer is the server API for GetPelanggan service.
// All implementations should embed UnimplementedGetPelangganServer
// for forward compatibility
type GetPelangganServer interface {
	// rpc getPelanggan (PelangganRequest) returns (PelangganResponse) {};
	GetPelangganApi(context.Context, *PelangganRequest) (*PelangganResponse, error)
	GetAllPelangganApi(context.Context, *PelangganRequest) (*AllPelangganResponse, error)
}

// UnimplementedGetPelangganServer should be embedded to have forward compatible implementations.
type UnimplementedGetPelangganServer struct {
}

func (UnimplementedGetPelangganServer) GetPelangganApi(context.Context, *PelangganRequest) (*PelangganResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPelangganApi not implemented")
}
func (UnimplementedGetPelangganServer) GetAllPelangganApi(context.Context, *PelangganRequest) (*AllPelangganResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPelangganApi not implemented")
}

// UnsafeGetPelangganServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetPelangganServer will
// result in compilation errors.
type UnsafeGetPelangganServer interface {
	mustEmbedUnimplementedGetPelangganServer()
}

func RegisterGetPelangganServer(s grpc.ServiceRegistrar, srv GetPelangganServer) {
	s.RegisterService(&GetPelanggan_ServiceDesc, srv)
}

func _GetPelanggan_GetPelangganApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PelangganRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetPelangganServer).GetPelangganApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcPelanggan.getPelanggan/getPelangganApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetPelangganServer).GetPelangganApi(ctx, req.(*PelangganRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetPelanggan_GetAllPelangganApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PelangganRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetPelangganServer).GetAllPelangganApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcPelanggan.getPelanggan/getAllPelangganApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetPelangganServer).GetAllPelangganApi(ctx, req.(*PelangganRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GetPelanggan_ServiceDesc is the grpc.ServiceDesc for GetPelanggan service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetPelanggan_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcPelanggan.getPelanggan",
	HandlerType: (*GetPelangganServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getPelangganApi",
			Handler:    _GetPelanggan_GetPelangganApi_Handler,
		},
		{
			MethodName: "getAllPelangganApi",
			Handler:    _GetPelanggan_GetAllPelangganApi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/pelanggan.proto",
}
