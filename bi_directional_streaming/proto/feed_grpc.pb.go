// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// FeedsClient is the client API for Feeds service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedsClient interface {
	// bi-directional rpc
	Broadcast(ctx context.Context, opts ...grpc.CallOption) (Feeds_BroadcastClient, error)
}

type feedsClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedsClient(cc grpc.ClientConnInterface) FeedsClient {
	return &feedsClient{cc}
}

func (c *feedsClient) Broadcast(ctx context.Context, opts ...grpc.CallOption) (Feeds_BroadcastClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Feeds_serviceDesc.Streams[0], "/proto.Feeds/Broadcast", opts...)
	if err != nil {
		return nil, err
	}
	x := &feedsBroadcastClient{stream}
	return x, nil
}

type Feeds_BroadcastClient interface {
	Send(*FeedRequest) error
	Recv() (*FeedResponse, error)
	grpc.ClientStream
}

type feedsBroadcastClient struct {
	grpc.ClientStream
}

func (x *feedsBroadcastClient) Send(m *FeedRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *feedsBroadcastClient) Recv() (*FeedResponse, error) {
	m := new(FeedResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FeedsServer is the server API for Feeds service.
// All implementations must embed UnimplementedFeedsServer
// for forward compatibility
type FeedsServer interface {
	// bi-directional rpc
	Broadcast(Feeds_BroadcastServer) error
	mustEmbedUnimplementedFeedsServer()
}

// UnimplementedFeedsServer must be embedded to have forward compatible implementations.
type UnimplementedFeedsServer struct {
}

func (UnimplementedFeedsServer) Broadcast(Feeds_BroadcastServer) error {
	return status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedFeedsServer) mustEmbedUnimplementedFeedsServer() {}

// UnsafeFeedsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedsServer will
// result in compilation errors.
type UnsafeFeedsServer interface {
	mustEmbedUnimplementedFeedsServer()
}

func RegisterFeedsServer(s grpc.ServiceRegistrar, srv FeedsServer) {
	s.RegisterService(&_Feeds_serviceDesc, srv)
}

func _Feeds_Broadcast_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FeedsServer).Broadcast(&feedsBroadcastServer{stream})
}

type Feeds_BroadcastServer interface {
	Send(*FeedResponse) error
	Recv() (*FeedRequest, error)
	grpc.ServerStream
}

type feedsBroadcastServer struct {
	grpc.ServerStream
}

func (x *feedsBroadcastServer) Send(m *FeedResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *feedsBroadcastServer) Recv() (*FeedRequest, error) {
	m := new(FeedRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Feeds_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Feeds",
	HandlerType: (*FeedsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Broadcast",
			Handler:       _Feeds_Broadcast_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/feed.proto",
}
