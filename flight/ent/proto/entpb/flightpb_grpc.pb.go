// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: entpb/flightpb.proto

package entpb

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
	FlightService_CreateFlight_FullMethodName  = "/entpb.FlightService/CreateFlight"
	FlightService_UpdateFlight_FullMethodName  = "/entpb.FlightService/UpdateFlight"
	FlightService_SearchFlights_FullMethodName = "/entpb.FlightService/SearchFlights"
)

// FlightServiceClient is the client API for FlightService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlightServiceClient interface {
	CreateFlight(ctx context.Context, in *CreateFlightInput, opts ...grpc.CallOption) (*Flight, error)
	UpdateFlight(ctx context.Context, in *UpdateFlightInput, opts ...grpc.CallOption) (*Flight, error)
	SearchFlights(ctx context.Context, in *QueryFlightInput, opts ...grpc.CallOption) (*SearchFlightResponse, error)
}

type flightServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlightServiceClient(cc grpc.ClientConnInterface) FlightServiceClient {
	return &flightServiceClient{cc}
}

func (c *flightServiceClient) CreateFlight(ctx context.Context, in *CreateFlightInput, opts ...grpc.CallOption) (*Flight, error) {
	out := new(Flight)
	err := c.cc.Invoke(ctx, FlightService_CreateFlight_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceClient) UpdateFlight(ctx context.Context, in *UpdateFlightInput, opts ...grpc.CallOption) (*Flight, error) {
	out := new(Flight)
	err := c.cc.Invoke(ctx, FlightService_UpdateFlight_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceClient) SearchFlights(ctx context.Context, in *QueryFlightInput, opts ...grpc.CallOption) (*SearchFlightResponse, error) {
	out := new(SearchFlightResponse)
	err := c.cc.Invoke(ctx, FlightService_SearchFlights_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlightServiceServer is the server API for FlightService service.
// All implementations must embed UnimplementedFlightServiceServer
// for forward compatibility
type FlightServiceServer interface {
	CreateFlight(context.Context, *CreateFlightInput) (*Flight, error)
	UpdateFlight(context.Context, *UpdateFlightInput) (*Flight, error)
	SearchFlights(context.Context, *QueryFlightInput) (*SearchFlightResponse, error)
	mustEmbedUnimplementedFlightServiceServer()
}

// UnimplementedFlightServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFlightServiceServer struct {
}

func (UnimplementedFlightServiceServer) CreateFlight(context.Context, *CreateFlightInput) (*Flight, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFlight not implemented")
}
func (UnimplementedFlightServiceServer) UpdateFlight(context.Context, *UpdateFlightInput) (*Flight, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFlight not implemented")
}
func (UnimplementedFlightServiceServer) SearchFlights(context.Context, *QueryFlightInput) (*SearchFlightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchFlights not implemented")
}
func (UnimplementedFlightServiceServer) mustEmbedUnimplementedFlightServiceServer() {}

// UnsafeFlightServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlightServiceServer will
// result in compilation errors.
type UnsafeFlightServiceServer interface {
	mustEmbedUnimplementedFlightServiceServer()
}

func RegisterFlightServiceServer(s grpc.ServiceRegistrar, srv FlightServiceServer) {
	s.RegisterService(&FlightService_ServiceDesc, srv)
}

func _FlightService_CreateFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFlightInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).CreateFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightService_CreateFlight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).CreateFlight(ctx, req.(*CreateFlightInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightService_UpdateFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFlightInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).UpdateFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightService_UpdateFlight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).UpdateFlight(ctx, req.(*UpdateFlightInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightService_SearchFlights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFlightInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).SearchFlights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightService_SearchFlights_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).SearchFlights(ctx, req.(*QueryFlightInput))
	}
	return interceptor(ctx, in, info, handler)
}

// FlightService_ServiceDesc is the grpc.ServiceDesc for FlightService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlightService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entpb.FlightService",
	HandlerType: (*FlightServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFlight",
			Handler:    _FlightService_CreateFlight_Handler,
		},
		{
			MethodName: "UpdateFlight",
			Handler:    _FlightService_UpdateFlight_Handler,
		},
		{
			MethodName: "SearchFlights",
			Handler:    _FlightService_SearchFlights_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entpb/flightpb.proto",
}
