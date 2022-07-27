// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: proto/courier.proto

package courier_proto

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

// CourierClient is the client API for Courier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourierClient interface {
	// Sends a greeting
	InsertNewCourier(ctx context.Context, in *NewCourierRequest, opts ...grpc.CallOption) (*CourierResponse, error)
	GetAllCourier(ctx context.Context, in *SearchParamCourier, opts ...grpc.CallOption) (*CourierListResponse, error)
	DeleteCourier(ctx context.Context, in *CourierID, opts ...grpc.CallOption) (*CourierDeleteResponse, error)
	UpdateCourier(ctx context.Context, in *UpdateCourierRequest, opts ...grpc.CallOption) (*CourierResponse, error)
	GetCourier(ctx context.Context, in *CourierID, opts ...grpc.CallOption) (*CourierResponse, error)
	InsertNewLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Location, error)
	UpdateLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Location, error)
	GetLocation(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Location, error)
	GetLocationList(ctx context.Context, in *SearchParamLocation, opts ...grpc.CallOption) (*LocationList, error)
}

type courierClient struct {
	cc grpc.ClientConnInterface
}

func NewCourierClient(cc grpc.ClientConnInterface) CourierClient {
	return &courierClient{cc}
}

func (c *courierClient) InsertNewCourier(ctx context.Context, in *NewCourierRequest, opts ...grpc.CallOption) (*CourierResponse, error) {
	out := new(CourierResponse)
	err := c.cc.Invoke(ctx, "/courier.Courier/insertNewCourier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierClient) GetAllCourier(ctx context.Context, in *SearchParamCourier, opts ...grpc.CallOption) (*CourierListResponse, error) {
	out := new(CourierListResponse)
	err := c.cc.Invoke(ctx, "/courier.Courier/getAllCourier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierClient) DeleteCourier(ctx context.Context, in *CourierID, opts ...grpc.CallOption) (*CourierDeleteResponse, error) {
	out := new(CourierDeleteResponse)
	err := c.cc.Invoke(ctx, "/courier.Courier/deleteCourier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierClient) UpdateCourier(ctx context.Context, in *UpdateCourierRequest, opts ...grpc.CallOption) (*CourierResponse, error) {
	out := new(CourierResponse)
	err := c.cc.Invoke(ctx, "/courier.Courier/updateCourier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierClient) GetCourier(ctx context.Context, in *CourierID, opts ...grpc.CallOption) (*CourierResponse, error) {
	out := new(CourierResponse)
	err := c.cc.Invoke(ctx, "/courier.Courier/getCourier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierClient) InsertNewLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Location, error) {
	out := new(Location)
	err := c.cc.Invoke(ctx, "/courier.Courier/insertNewLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierClient) UpdateLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*Location, error) {
	out := new(Location)
	err := c.cc.Invoke(ctx, "/courier.Courier/updateLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierClient) GetLocation(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Location, error) {
	out := new(Location)
	err := c.cc.Invoke(ctx, "/courier.Courier/getLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierClient) GetLocationList(ctx context.Context, in *SearchParamLocation, opts ...grpc.CallOption) (*LocationList, error) {
	out := new(LocationList)
	err := c.cc.Invoke(ctx, "/courier.Courier/getLocationList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourierServer is the server API for Courier service.
// All implementations must embed UnimplementedCourierServer
// for forward compatibility
type CourierServer interface {
	// Sends a greeting
	InsertNewCourier(context.Context, *NewCourierRequest) (*CourierResponse, error)
	GetAllCourier(context.Context, *SearchParamCourier) (*CourierListResponse, error)
	DeleteCourier(context.Context, *CourierID) (*CourierDeleteResponse, error)
	UpdateCourier(context.Context, *UpdateCourierRequest) (*CourierResponse, error)
	GetCourier(context.Context, *CourierID) (*CourierResponse, error)
	InsertNewLocation(context.Context, *Location) (*Location, error)
	UpdateLocation(context.Context, *Location) (*Location, error)
	GetLocation(context.Context, *UserID) (*Location, error)
	GetLocationList(context.Context, *SearchParamLocation) (*LocationList, error)
	mustEmbedUnimplementedCourierServer()
}

// UnimplementedCourierServer must be embedded to have forward compatible implementations.
type UnimplementedCourierServer struct {
}

func (UnimplementedCourierServer) InsertNewCourier(context.Context, *NewCourierRequest) (*CourierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertNewCourier not implemented")
}
func (UnimplementedCourierServer) GetAllCourier(context.Context, *SearchParamCourier) (*CourierListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCourier not implemented")
}
func (UnimplementedCourierServer) DeleteCourier(context.Context, *CourierID) (*CourierDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourier not implemented")
}
func (UnimplementedCourierServer) UpdateCourier(context.Context, *UpdateCourierRequest) (*CourierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCourier not implemented")
}
func (UnimplementedCourierServer) GetCourier(context.Context, *CourierID) (*CourierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourier not implemented")
}
func (UnimplementedCourierServer) InsertNewLocation(context.Context, *Location) (*Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertNewLocation not implemented")
}
func (UnimplementedCourierServer) UpdateLocation(context.Context, *Location) (*Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLocation not implemented")
}
func (UnimplementedCourierServer) GetLocation(context.Context, *UserID) (*Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocation not implemented")
}
func (UnimplementedCourierServer) GetLocationList(context.Context, *SearchParamLocation) (*LocationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocationList not implemented")
}
func (UnimplementedCourierServer) mustEmbedUnimplementedCourierServer() {}

// UnsafeCourierServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourierServer will
// result in compilation errors.
type UnsafeCourierServer interface {
	mustEmbedUnimplementedCourierServer()
}

func RegisterCourierServer(s grpc.ServiceRegistrar, srv CourierServer) {
	s.RegisterService(&Courier_ServiceDesc, srv)
}

func _Courier_InsertNewCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCourierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServer).InsertNewCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/courier.Courier/insertNewCourier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServer).InsertNewCourier(ctx, req.(*NewCourierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courier_GetAllCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchParamCourier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServer).GetAllCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/courier.Courier/getAllCourier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServer).GetAllCourier(ctx, req.(*SearchParamCourier))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courier_DeleteCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourierID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServer).DeleteCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/courier.Courier/deleteCourier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServer).DeleteCourier(ctx, req.(*CourierID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courier_UpdateCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCourierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServer).UpdateCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/courier.Courier/updateCourier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServer).UpdateCourier(ctx, req.(*UpdateCourierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courier_GetCourier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourierID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServer).GetCourier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/courier.Courier/getCourier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServer).GetCourier(ctx, req.(*CourierID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courier_InsertNewLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServer).InsertNewLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/courier.Courier/insertNewLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServer).InsertNewLocation(ctx, req.(*Location))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courier_UpdateLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServer).UpdateLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/courier.Courier/updateLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServer).UpdateLocation(ctx, req.(*Location))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courier_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/courier.Courier/getLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServer).GetLocation(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courier_GetLocationList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchParamLocation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServer).GetLocationList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/courier.Courier/getLocationList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServer).GetLocationList(ctx, req.(*SearchParamLocation))
	}
	return interceptor(ctx, in, info, handler)
}

// Courier_ServiceDesc is the grpc.ServiceDesc for Courier service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Courier_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "courier.Courier",
	HandlerType: (*CourierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "insertNewCourier",
			Handler:    _Courier_InsertNewCourier_Handler,
		},
		{
			MethodName: "getAllCourier",
			Handler:    _Courier_GetAllCourier_Handler,
		},
		{
			MethodName: "deleteCourier",
			Handler:    _Courier_DeleteCourier_Handler,
		},
		{
			MethodName: "updateCourier",
			Handler:    _Courier_UpdateCourier_Handler,
		},
		{
			MethodName: "getCourier",
			Handler:    _Courier_GetCourier_Handler,
		},
		{
			MethodName: "insertNewLocation",
			Handler:    _Courier_InsertNewLocation_Handler,
		},
		{
			MethodName: "updateLocation",
			Handler:    _Courier_UpdateLocation_Handler,
		},
		{
			MethodName: "getLocation",
			Handler:    _Courier_GetLocation_Handler,
		},
		{
			MethodName: "getLocationList",
			Handler:    _Courier_GetLocationList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/courier.proto",
}
