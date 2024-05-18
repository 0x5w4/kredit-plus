// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: reader.proto

package readerService

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

const (
	ReaderService_CreateKonsumen_FullMethodName  = "/readerService.readerService/CreateKonsumen"
	ReaderService_CreateLimit_FullMethodName     = "/readerService.readerService/CreateLimit"
	ReaderService_CreateTransaksi_FullMethodName = "/readerService.readerService/CreateTransaksi"
	ReaderService_GetLimit_FullMethodName        = "/readerService.readerService/GetLimit"
	ReaderService_GetTransaksi_FullMethodName    = "/readerService.readerService/GetTransaksi"
)

// ReaderServiceClient is the client API for ReaderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReaderServiceClient interface {
	CreateKonsumen(ctx context.Context, in *CreateKonsumenRequest, opts ...grpc.CallOption) (*CreateKonsumenResponse, error)
	CreateLimit(ctx context.Context, in *CreateLimitRequest, opts ...grpc.CallOption) (*CreateLimitResponse, error)
	CreateTransaksi(ctx context.Context, in *CreateTransaksiRequest, opts ...grpc.CallOption) (*CreateTransaksiResponse, error)
	GetLimit(ctx context.Context, in *GetLimitRequest, opts ...grpc.CallOption) (*GetLimitResponse, error)
	GetTransaksi(ctx context.Context, in *GetTransaksiRequest, opts ...grpc.CallOption) (*GetTransaksiResponse, error)
}

type readerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReaderServiceClient(cc grpc.ClientConnInterface) ReaderServiceClient {
	return &readerServiceClient{cc}
}

func (c *readerServiceClient) CreateKonsumen(ctx context.Context, in *CreateKonsumenRequest, opts ...grpc.CallOption) (*CreateKonsumenResponse, error) {
	out := new(CreateKonsumenResponse)
	err := c.cc.Invoke(ctx, ReaderService_CreateKonsumen_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readerServiceClient) CreateLimit(ctx context.Context, in *CreateLimitRequest, opts ...grpc.CallOption) (*CreateLimitResponse, error) {
	out := new(CreateLimitResponse)
	err := c.cc.Invoke(ctx, ReaderService_CreateLimit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readerServiceClient) CreateTransaksi(ctx context.Context, in *CreateTransaksiRequest, opts ...grpc.CallOption) (*CreateTransaksiResponse, error) {
	out := new(CreateTransaksiResponse)
	err := c.cc.Invoke(ctx, ReaderService_CreateTransaksi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readerServiceClient) GetLimit(ctx context.Context, in *GetLimitRequest, opts ...grpc.CallOption) (*GetLimitResponse, error) {
	out := new(GetLimitResponse)
	err := c.cc.Invoke(ctx, ReaderService_GetLimit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *readerServiceClient) GetTransaksi(ctx context.Context, in *GetTransaksiRequest, opts ...grpc.CallOption) (*GetTransaksiResponse, error) {
	out := new(GetTransaksiResponse)
	err := c.cc.Invoke(ctx, ReaderService_GetTransaksi_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReaderServiceServer is the server API for ReaderService service.
// All implementations should embed UnimplementedReaderServiceServer
// for forward compatibility
type ReaderServiceServer interface {
	CreateKonsumen(context.Context, *CreateKonsumenRequest) (*CreateKonsumenResponse, error)
	CreateLimit(context.Context, *CreateLimitRequest) (*CreateLimitResponse, error)
	CreateTransaksi(context.Context, *CreateTransaksiRequest) (*CreateTransaksiResponse, error)
	GetLimit(context.Context, *GetLimitRequest) (*GetLimitResponse, error)
	GetTransaksi(context.Context, *GetTransaksiRequest) (*GetTransaksiResponse, error)
}

// UnimplementedReaderServiceServer should be embedded to have forward compatible implementations.
type UnimplementedReaderServiceServer struct {
}

func (UnimplementedReaderServiceServer) CreateKonsumen(context.Context, *CreateKonsumenRequest) (*CreateKonsumenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKonsumen not implemented")
}
func (UnimplementedReaderServiceServer) CreateLimit(context.Context, *CreateLimitRequest) (*CreateLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLimit not implemented")
}
func (UnimplementedReaderServiceServer) CreateTransaksi(context.Context, *CreateTransaksiRequest) (*CreateTransaksiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaksi not implemented")
}
func (UnimplementedReaderServiceServer) GetLimit(context.Context, *GetLimitRequest) (*GetLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLimit not implemented")
}
func (UnimplementedReaderServiceServer) GetTransaksi(context.Context, *GetTransaksiRequest) (*GetTransaksiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaksi not implemented")
}

// UnsafeReaderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReaderServiceServer will
// result in compilation errors.
type UnsafeReaderServiceServer interface {
	mustEmbedUnimplementedReaderServiceServer()
}

func RegisterReaderServiceServer(s grpc.ServiceRegistrar, srv ReaderServiceServer) {
	s.RegisterService(&ReaderService_ServiceDesc, srv)
}

func _ReaderService_CreateKonsumen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKonsumenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReaderServiceServer).CreateKonsumen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReaderService_CreateKonsumen_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReaderServiceServer).CreateKonsumen(ctx, req.(*CreateKonsumenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReaderService_CreateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReaderServiceServer).CreateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReaderService_CreateLimit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReaderServiceServer).CreateLimit(ctx, req.(*CreateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReaderService_CreateTransaksi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransaksiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReaderServiceServer).CreateTransaksi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReaderService_CreateTransaksi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReaderServiceServer).CreateTransaksi(ctx, req.(*CreateTransaksiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReaderService_GetLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReaderServiceServer).GetLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReaderService_GetLimit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReaderServiceServer).GetLimit(ctx, req.(*GetLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReaderService_GetTransaksi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransaksiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReaderServiceServer).GetTransaksi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReaderService_GetTransaksi_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReaderServiceServer).GetTransaksi(ctx, req.(*GetTransaksiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReaderService_ServiceDesc is the grpc.ServiceDesc for ReaderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReaderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "readerService.readerService",
	HandlerType: (*ReaderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateKonsumen",
			Handler:    _ReaderService_CreateKonsumen_Handler,
		},
		{
			MethodName: "CreateLimit",
			Handler:    _ReaderService_CreateLimit_Handler,
		},
		{
			MethodName: "CreateTransaksi",
			Handler:    _ReaderService_CreateTransaksi_Handler,
		},
		{
			MethodName: "GetLimit",
			Handler:    _ReaderService_GetLimit_Handler,
		},
		{
			MethodName: "GetTransaksi",
			Handler:    _ReaderService_GetTransaksi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reader.proto",
}