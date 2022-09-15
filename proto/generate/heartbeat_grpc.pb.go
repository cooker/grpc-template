// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: heartbeat.proto

package proto

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

// HeartBeatServiceClient is the client API for HeartBeatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeartBeatServiceClient interface {
	Send(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*MessageResponse, error)
}

type heartBeatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHeartBeatServiceClient(cc grpc.ClientConnInterface) HeartBeatServiceClient {
	return &heartBeatServiceClient{cc}
}

func (c *heartBeatServiceClient) Send(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/proto.HeartBeatService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartBeatServiceServer is the server API for HeartBeatService service.
// All implementations must embed UnimplementedHeartBeatServiceServer
// for forward compatibility
type HeartBeatServiceServer interface {
	Send(context.Context, *HeartBeatRequest) (*MessageResponse, error)
	mustEmbedUnimplementedHeartBeatServiceServer()
}

// UnimplementedHeartBeatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHeartBeatServiceServer struct {
}

func (UnimplementedHeartBeatServiceServer) Send(context.Context, *HeartBeatRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedHeartBeatServiceServer) mustEmbedUnimplementedHeartBeatServiceServer() {}

// UnsafeHeartBeatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeartBeatServiceServer will
// result in compilation errors.
type UnsafeHeartBeatServiceServer interface {
	mustEmbedUnimplementedHeartBeatServiceServer()
}

func RegisterHeartBeatServiceServer(s grpc.ServiceRegistrar, srv HeartBeatServiceServer) {
	s.RegisterService(&HeartBeatService_ServiceDesc, srv)
}

func _HeartBeatService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartBeatServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HeartBeatService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartBeatServiceServer).Send(ctx, req.(*HeartBeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HeartBeatService_ServiceDesc is the grpc.ServiceDesc for HeartBeatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeartBeatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HeartBeatService",
	HandlerType: (*HeartBeatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _HeartBeatService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heartbeat.proto",
}