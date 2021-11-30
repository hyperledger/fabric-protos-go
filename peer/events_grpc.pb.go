// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package peer

import (
	context "context"
	common "github.com/hyperledger/fabric-protos-go/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DeliverClient is the client API for Deliver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeliverClient interface {
	// Deliver first requires an Envelope of type ab.DELIVER_SEEK_INFO with
	// Payload data as a marshaled orderer.SeekInfo message,
	// then a stream of block replies is received
	Deliver(ctx context.Context, opts ...grpc.CallOption) (Deliver_DeliverClient, error)
	// DeliverFiltered first requires an Envelope of type ab.DELIVER_SEEK_INFO with
	// Payload data as a marshaled orderer.SeekInfo message,
	// then a stream of **filtered** block replies is received
	DeliverFiltered(ctx context.Context, opts ...grpc.CallOption) (Deliver_DeliverFilteredClient, error)
	// DeliverWithPrivateData first requires an Envelope of type ab.DELIVER_SEEK_INFO with
	// Payload data as a marshaled orderer.SeekInfo message,
	// then a stream of block and private data replies is received
	DeliverWithPrivateData(ctx context.Context, opts ...grpc.CallOption) (Deliver_DeliverWithPrivateDataClient, error)
}

type deliverClient struct {
	cc grpc.ClientConnInterface
}

func NewDeliverClient(cc grpc.ClientConnInterface) DeliverClient {
	return &deliverClient{cc}
}

func (c *deliverClient) Deliver(ctx context.Context, opts ...grpc.CallOption) (Deliver_DeliverClient, error) {
	stream, err := c.cc.NewStream(ctx, &Deliver_ServiceDesc.Streams[0], "/protos.Deliver/Deliver", opts...)
	if err != nil {
		return nil, err
	}
	x := &deliverDeliverClient{stream}
	return x, nil
}

type Deliver_DeliverClient interface {
	Send(*common.Envelope) error
	Recv() (*DeliverResponse, error)
	grpc.ClientStream
}

type deliverDeliverClient struct {
	grpc.ClientStream
}

func (x *deliverDeliverClient) Send(m *common.Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *deliverDeliverClient) Recv() (*DeliverResponse, error) {
	m := new(DeliverResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *deliverClient) DeliverFiltered(ctx context.Context, opts ...grpc.CallOption) (Deliver_DeliverFilteredClient, error) {
	stream, err := c.cc.NewStream(ctx, &Deliver_ServiceDesc.Streams[1], "/protos.Deliver/DeliverFiltered", opts...)
	if err != nil {
		return nil, err
	}
	x := &deliverDeliverFilteredClient{stream}
	return x, nil
}

type Deliver_DeliverFilteredClient interface {
	Send(*common.Envelope) error
	Recv() (*DeliverResponse, error)
	grpc.ClientStream
}

type deliverDeliverFilteredClient struct {
	grpc.ClientStream
}

func (x *deliverDeliverFilteredClient) Send(m *common.Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *deliverDeliverFilteredClient) Recv() (*DeliverResponse, error) {
	m := new(DeliverResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *deliverClient) DeliverWithPrivateData(ctx context.Context, opts ...grpc.CallOption) (Deliver_DeliverWithPrivateDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &Deliver_ServiceDesc.Streams[2], "/protos.Deliver/DeliverWithPrivateData", opts...)
	if err != nil {
		return nil, err
	}
	x := &deliverDeliverWithPrivateDataClient{stream}
	return x, nil
}

type Deliver_DeliverWithPrivateDataClient interface {
	Send(*common.Envelope) error
	Recv() (*DeliverResponse, error)
	grpc.ClientStream
}

type deliverDeliverWithPrivateDataClient struct {
	grpc.ClientStream
}

func (x *deliverDeliverWithPrivateDataClient) Send(m *common.Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *deliverDeliverWithPrivateDataClient) Recv() (*DeliverResponse, error) {
	m := new(DeliverResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DeliverServer is the server API for Deliver service.
// All implementations must embed UnimplementedDeliverServer
// for forward compatibility
type DeliverServer interface {
	// Deliver first requires an Envelope of type ab.DELIVER_SEEK_INFO with
	// Payload data as a marshaled orderer.SeekInfo message,
	// then a stream of block replies is received
	Deliver(Deliver_DeliverServer) error
	// DeliverFiltered first requires an Envelope of type ab.DELIVER_SEEK_INFO with
	// Payload data as a marshaled orderer.SeekInfo message,
	// then a stream of **filtered** block replies is received
	DeliverFiltered(Deliver_DeliverFilteredServer) error
	// DeliverWithPrivateData first requires an Envelope of type ab.DELIVER_SEEK_INFO with
	// Payload data as a marshaled orderer.SeekInfo message,
	// then a stream of block and private data replies is received
	DeliverWithPrivateData(Deliver_DeliverWithPrivateDataServer) error
	mustEmbedUnimplementedDeliverServer()
}

// UnimplementedDeliverServer must be embedded to have forward compatible implementations.
type UnimplementedDeliverServer struct {
}

func (UnimplementedDeliverServer) Deliver(Deliver_DeliverServer) error {
	return status.Errorf(codes.Unimplemented, "method Deliver not implemented")
}
func (UnimplementedDeliverServer) DeliverFiltered(Deliver_DeliverFilteredServer) error {
	return status.Errorf(codes.Unimplemented, "method DeliverFiltered not implemented")
}
func (UnimplementedDeliverServer) DeliverWithPrivateData(Deliver_DeliverWithPrivateDataServer) error {
	return status.Errorf(codes.Unimplemented, "method DeliverWithPrivateData not implemented")
}
func (UnimplementedDeliverServer) mustEmbedUnimplementedDeliverServer() {}

// UnsafeDeliverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeliverServer will
// result in compilation errors.
type UnsafeDeliverServer interface {
	mustEmbedUnimplementedDeliverServer()
}

func RegisterDeliverServer(s grpc.ServiceRegistrar, srv DeliverServer) {
	s.RegisterService(&Deliver_ServiceDesc, srv)
}

func _Deliver_Deliver_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DeliverServer).Deliver(&deliverDeliverServer{stream})
}

type Deliver_DeliverServer interface {
	Send(*DeliverResponse) error
	Recv() (*common.Envelope, error)
	grpc.ServerStream
}

type deliverDeliverServer struct {
	grpc.ServerStream
}

func (x *deliverDeliverServer) Send(m *DeliverResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *deliverDeliverServer) Recv() (*common.Envelope, error) {
	m := new(common.Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Deliver_DeliverFiltered_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DeliverServer).DeliverFiltered(&deliverDeliverFilteredServer{stream})
}

type Deliver_DeliverFilteredServer interface {
	Send(*DeliverResponse) error
	Recv() (*common.Envelope, error)
	grpc.ServerStream
}

type deliverDeliverFilteredServer struct {
	grpc.ServerStream
}

func (x *deliverDeliverFilteredServer) Send(m *DeliverResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *deliverDeliverFilteredServer) Recv() (*common.Envelope, error) {
	m := new(common.Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Deliver_DeliverWithPrivateData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DeliverServer).DeliverWithPrivateData(&deliverDeliverWithPrivateDataServer{stream})
}

type Deliver_DeliverWithPrivateDataServer interface {
	Send(*DeliverResponse) error
	Recv() (*common.Envelope, error)
	grpc.ServerStream
}

type deliverDeliverWithPrivateDataServer struct {
	grpc.ServerStream
}

func (x *deliverDeliverWithPrivateDataServer) Send(m *DeliverResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *deliverDeliverWithPrivateDataServer) Recv() (*common.Envelope, error) {
	m := new(common.Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Deliver_ServiceDesc is the grpc.ServiceDesc for Deliver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Deliver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Deliver",
	HandlerType: (*DeliverServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Deliver",
			Handler:       _Deliver_Deliver_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeliverFiltered",
			Handler:       _Deliver_DeliverFiltered_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeliverWithPrivateData",
			Handler:       _Deliver_DeliverWithPrivateData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "peer/events.proto",
}
