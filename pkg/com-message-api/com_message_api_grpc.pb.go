// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package com_message_api

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

// ComMessageApiServiceClient is the client API for ComMessageApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComMessageApiServiceClient interface {
	// CreateMessageV1 - Create a message
	CreateMessageV1(ctx context.Context, in *CreateMessageV1Request, opts ...grpc.CallOption) (*CreateMessageV1Response, error)
	// DescribeMessageV1 - Describe a message
	DescribeMessageV1(ctx context.Context, in *DescribeMessageV1Request, opts ...grpc.CallOption) (*DescribeMessageV1Response, error)
	// ListMessageV1 - List a messages
	ListMessageV1(ctx context.Context, in *ListMessageV1Request, opts ...grpc.CallOption) (*ListMessageV1Response, error)
	// RemoveMessageV1 - Describe a message
	RemoveMessageV1(ctx context.Context, in *RemoveMessageV1Request, opts ...grpc.CallOption) (*RemoveMessageV1Response, error)
	// UpdateMessageV1 - Create a message
	UpdateMessageV1(ctx context.Context, in *UpdateMessageV1Request, opts ...grpc.CallOption) (*UpdateMessageV1Response, error)
}

type comMessageApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComMessageApiServiceClient(cc grpc.ClientConnInterface) ComMessageApiServiceClient {
	return &comMessageApiServiceClient{cc}
}

func (c *comMessageApiServiceClient) CreateMessageV1(ctx context.Context, in *CreateMessageV1Request, opts ...grpc.CallOption) (*CreateMessageV1Response, error) {
	out := new(CreateMessageV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.com_message_api.v1.ComMessageApiService/CreateMessageV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comMessageApiServiceClient) DescribeMessageV1(ctx context.Context, in *DescribeMessageV1Request, opts ...grpc.CallOption) (*DescribeMessageV1Response, error) {
	out := new(DescribeMessageV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.com_message_api.v1.ComMessageApiService/DescribeMessageV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comMessageApiServiceClient) ListMessageV1(ctx context.Context, in *ListMessageV1Request, opts ...grpc.CallOption) (*ListMessageV1Response, error) {
	out := new(ListMessageV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.com_message_api.v1.ComMessageApiService/ListMessageV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comMessageApiServiceClient) RemoveMessageV1(ctx context.Context, in *RemoveMessageV1Request, opts ...grpc.CallOption) (*RemoveMessageV1Response, error) {
	out := new(RemoveMessageV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.com_message_api.v1.ComMessageApiService/RemoveMessageV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comMessageApiServiceClient) UpdateMessageV1(ctx context.Context, in *UpdateMessageV1Request, opts ...grpc.CallOption) (*UpdateMessageV1Response, error) {
	out := new(UpdateMessageV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.com_message_api.v1.ComMessageApiService/UpdateMessageV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComMessageApiServiceServer is the server API for ComMessageApiService service.
// All implementations must embed UnimplementedComMessageApiServiceServer
// for forward compatibility
type ComMessageApiServiceServer interface {
	// CreateMessageV1 - Create a message
	CreateMessageV1(context.Context, *CreateMessageV1Request) (*CreateMessageV1Response, error)
	// DescribeMessageV1 - Describe a message
	DescribeMessageV1(context.Context, *DescribeMessageV1Request) (*DescribeMessageV1Response, error)
	// ListMessageV1 - List a messages
	ListMessageV1(context.Context, *ListMessageV1Request) (*ListMessageV1Response, error)
	// RemoveMessageV1 - Describe a message
	RemoveMessageV1(context.Context, *RemoveMessageV1Request) (*RemoveMessageV1Response, error)
	// UpdateMessageV1 - Create a message
	UpdateMessageV1(context.Context, *UpdateMessageV1Request) (*UpdateMessageV1Response, error)
	mustEmbedUnimplementedComMessageApiServiceServer()
}

// UnimplementedComMessageApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedComMessageApiServiceServer struct {
}

func (UnimplementedComMessageApiServiceServer) CreateMessageV1(context.Context, *CreateMessageV1Request) (*CreateMessageV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessageV1 not implemented")
}
func (UnimplementedComMessageApiServiceServer) DescribeMessageV1(context.Context, *DescribeMessageV1Request) (*DescribeMessageV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeMessageV1 not implemented")
}
func (UnimplementedComMessageApiServiceServer) ListMessageV1(context.Context, *ListMessageV1Request) (*ListMessageV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMessageV1 not implemented")
}
func (UnimplementedComMessageApiServiceServer) RemoveMessageV1(context.Context, *RemoveMessageV1Request) (*RemoveMessageV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMessageV1 not implemented")
}
func (UnimplementedComMessageApiServiceServer) UpdateMessageV1(context.Context, *UpdateMessageV1Request) (*UpdateMessageV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessageV1 not implemented")
}
func (UnimplementedComMessageApiServiceServer) mustEmbedUnimplementedComMessageApiServiceServer() {}

// UnsafeComMessageApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComMessageApiServiceServer will
// result in compilation errors.
type UnsafeComMessageApiServiceServer interface {
	mustEmbedUnimplementedComMessageApiServiceServer()
}

func RegisterComMessageApiServiceServer(s grpc.ServiceRegistrar, srv ComMessageApiServiceServer) {
	s.RegisterService(&ComMessageApiService_ServiceDesc, srv)
}

func _ComMessageApiService_CreateMessageV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComMessageApiServiceServer).CreateMessageV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.com_message_api.v1.ComMessageApiService/CreateMessageV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComMessageApiServiceServer).CreateMessageV1(ctx, req.(*CreateMessageV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComMessageApiService_DescribeMessageV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeMessageV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComMessageApiServiceServer).DescribeMessageV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.com_message_api.v1.ComMessageApiService/DescribeMessageV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComMessageApiServiceServer).DescribeMessageV1(ctx, req.(*DescribeMessageV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComMessageApiService_ListMessageV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMessageV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComMessageApiServiceServer).ListMessageV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.com_message_api.v1.ComMessageApiService/ListMessageV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComMessageApiServiceServer).ListMessageV1(ctx, req.(*ListMessageV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComMessageApiService_RemoveMessageV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveMessageV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComMessageApiServiceServer).RemoveMessageV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.com_message_api.v1.ComMessageApiService/RemoveMessageV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComMessageApiServiceServer).RemoveMessageV1(ctx, req.(*RemoveMessageV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComMessageApiService_UpdateMessageV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessageV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComMessageApiServiceServer).UpdateMessageV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.com_message_api.v1.ComMessageApiService/UpdateMessageV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComMessageApiServiceServer).UpdateMessageV1(ctx, req.(*UpdateMessageV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ComMessageApiService_ServiceDesc is the grpc.ServiceDesc for ComMessageApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComMessageApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ozonmp.com_message_api.v1.ComMessageApiService",
	HandlerType: (*ComMessageApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMessageV1",
			Handler:    _ComMessageApiService_CreateMessageV1_Handler,
		},
		{
			MethodName: "DescribeMessageV1",
			Handler:    _ComMessageApiService_DescribeMessageV1_Handler,
		},
		{
			MethodName: "ListMessageV1",
			Handler:    _ComMessageApiService_ListMessageV1_Handler,
		},
		{
			MethodName: "RemoveMessageV1",
			Handler:    _ComMessageApiService_RemoveMessageV1_Handler,
		},
		{
			MethodName: "UpdateMessageV1",
			Handler:    _ComMessageApiService_UpdateMessageV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ozonmp/com_message_api/v1/com_message_api.proto",
}
