// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.2
// source: resource.proto

package proto_gen

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

// ResourceServiceClient is the client API for ResourceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceServiceClient interface {
	// Методы для управления рабочими местами
	GetWorkplaces(ctx context.Context, in *GetWorkplacesRequest, opts ...grpc.CallOption) (*GetWorkplacesResponse, error)
	GetWorkplaceById(ctx context.Context, in *GetWorkplaceByIdRequest, opts ...grpc.CallOption) (*Workplace, error)
	CreateWorkplace(ctx context.Context, in *CreateWorkplaceRequest, opts ...grpc.CallOption) (*Workplace, error)
	UpdateWorkplace(ctx context.Context, in *UpdateWorkplaceRequest, opts ...grpc.CallOption) (*Workplace, error)
	DeleteWorkplace(ctx context.Context, in *DeleteWorkplaceRequest, opts ...grpc.CallOption) (*DeleteWorkplaceResponse, error)
	CheckWorkplaceAvailability(ctx context.Context, in *CheckWorkplaceAvailabilityRequest, opts ...grpc.CallOption) (*WorkplaceAvailabilityResponse, error)
	// Методы для управления парковочными местами
	GetParkingSpaces(ctx context.Context, in *GetParkingSpacesRequest, opts ...grpc.CallOption) (*GetParkingSpacesResponse, error)
	GetParkingSpaceById(ctx context.Context, in *GetParkingSpaceByIdRequest, opts ...grpc.CallOption) (*ParkingSpace, error)
	CreateParkingSpace(ctx context.Context, in *CreateParkingSpaceRequest, opts ...grpc.CallOption) (*ParkingSpace, error)
	UpdateParkingSpace(ctx context.Context, in *UpdateParkingSpaceRequest, opts ...grpc.CallOption) (*ParkingSpace, error)
	DeleteParkingSpace(ctx context.Context, in *DeleteParkingSpaceRequest, opts ...grpc.CallOption) (*DeleteParkingSpaceResponse, error)
	CheckParkingSpaceAvailability(ctx context.Context, in *CheckParkingSpaceAvailabilityRequest, opts ...grpc.CallOption) (*ParkingSpaceAvailabilityResponse, error)
	// Методы для управления предметами
	GetItems(ctx context.Context, in *GetItemsRequest, opts ...grpc.CallOption) (*GetItemsResponse, error)
	GetItemById(ctx context.Context, in *GetItemByIdRequest, opts ...grpc.CallOption) (*Item, error)
	CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*Item, error)
	UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*Item, error)
	DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemResponse, error)
	AttachItemToWorkplace(ctx context.Context, in *AttachItemToWorkplaceRequest, opts ...grpc.CallOption) (*Item, error)
}

type resourceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceServiceClient(cc grpc.ClientConnInterface) ResourceServiceClient {
	return &resourceServiceClient{cc}
}

func (c *resourceServiceClient) GetWorkplaces(ctx context.Context, in *GetWorkplacesRequest, opts ...grpc.CallOption) (*GetWorkplacesResponse, error) {
	out := new(GetWorkplacesResponse)
	err := c.cc.Invoke(ctx, "/ResourceService.ResourceService/GetWorkplaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) GetWorkplaceById(ctx context.Context, in *GetWorkplaceByIdRequest, opts ...grpc.CallOption) (*Workplace, error) {
	out := new(Workplace)
	err := c.cc.Invoke(ctx, "/ResourceService.ResourceService/GetWorkplaceById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) CreateWorkplace(ctx context.Context, in *CreateWorkplaceRequest, opts ...grpc.CallOption) (*Workplace, error) {
	out := new(Workplace)
	err := c.cc.Invoke(ctx, "/ResourceService.ResourceService/CreateWorkplace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) UpdateWorkplace(ctx context.Context, in *UpdateWorkplaceRequest, opts ...grpc.CallOption) (*Workplace, error) {
	out := new(Workplace)
	err := c.cc.Invoke(ctx, "/ResourceService.ResourceService/UpdateWorkplace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) DeleteWorkplace(ctx context.Context, in *DeleteWorkplaceRequest, opts ...grpc.CallOption) (*DeleteWorkplaceResponse, error) {
	out := new(DeleteWorkplaceResponse)
	err := c.cc.Invoke(ctx, "/ResourceService.ResourceService/DeleteWorkplace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) CheckWorkplaceAvailability(ctx context.Context, in *CheckWorkplaceAvailabilityRequest, opts ...grpc.CallOption) (*WorkplaceAvailabilityResponse, error) {
	out := new(WorkplaceAvailabilityResponse)
	err := c.cc.Invoke(ctx, "/ResourceService.ResourceService/CheckWorkplaceAvailability", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) GetParkingSpaces(ctx context.Context, in *GetParkingSpacesRequest, opts ...grpc.CallOption) (*GetParkingSpacesResponse, error) {
	out := new(GetParkingSpacesResponse)
	err := c.cc.Invoke(ctx, "/ResourceService.ResourceService/GetParkingSpaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) GetParkingSpaceById(ctx context.Context, in *GetParkingSpaceByIdRequest, opts ...grpc.CallOption) (*ParkingSpace, error) {
	out := new(ParkingSpace)
	err := c.cc.Invoke(ctx, "/ResourceService.ResourceService/GetParkingSpaceById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) CreateParkingSpace(ctx context.Context, in *CreateParkingSpaceRequest, opts ...grpc.CallOption) (*ParkingSpace, error) {
	out := new(ParkingSpace)
	err := c.cc.Invoke(ctx, "/ResourceService.ResourceService/CreateParkingSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) UpdateParkingSpace(ctx context.Context, in *UpdateParkingSpaceRequest, opts ...grpc.CallOption) (*ParkingSpace, error) {
	out := new(ParkingSpace)
	err := c.cc.Invoke(ctx, "/ResourceService.ResourceService/UpdateParkingSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) DeleteParkingSpace(ctx context.Context, in *DeleteParkingSpaceRequest, opts ...grpc.CallOption) (*DeleteParkingSpaceResponse, error) {
	out := new(DeleteParkingSpaceResponse)
	err := c.cc.Invoke(ctx, "/ResourceService.ResourceService/DeleteParkingSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) CheckParkingSpaceAvailability(ctx context.Context, in *CheckParkingSpaceAvailabilityRequest, opts ...grpc.CallOption) (*ParkingSpaceAvailabilityResponse, error) {
	out := new(ParkingSpaceAvailabilityResponse)
	err := c.cc.Invoke(ctx, "/ResourceService.ResourceService/CheckParkingSpaceAvailability", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) GetItems(ctx context.Context, in *GetItemsRequest, opts ...grpc.CallOption) (*GetItemsResponse, error) {
	out := new(GetItemsResponse)
	err := c.cc.Invoke(ctx, "/ResourceService.ResourceService/GetItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) GetItemById(ctx context.Context, in *GetItemByIdRequest, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/ResourceService.ResourceService/GetItemById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/ResourceService.ResourceService/CreateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/ResourceService.ResourceService/UpdateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemResponse, error) {
	out := new(DeleteItemResponse)
	err := c.cc.Invoke(ctx, "/ResourceService.ResourceService/DeleteItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceServiceClient) AttachItemToWorkplace(ctx context.Context, in *AttachItemToWorkplaceRequest, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/ResourceService.ResourceService/AttachItemToWorkplace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceServiceServer is the server API for ResourceService service.
// All implementations must embed UnimplementedResourceServiceServer
// for forward compatibility
type ResourceServiceServer interface {
	// Методы для управления рабочими местами
	GetWorkplaces(context.Context, *GetWorkplacesRequest) (*GetWorkplacesResponse, error)
	GetWorkplaceById(context.Context, *GetWorkplaceByIdRequest) (*Workplace, error)
	CreateWorkplace(context.Context, *CreateWorkplaceRequest) (*Workplace, error)
	UpdateWorkplace(context.Context, *UpdateWorkplaceRequest) (*Workplace, error)
	DeleteWorkplace(context.Context, *DeleteWorkplaceRequest) (*DeleteWorkplaceResponse, error)
	CheckWorkplaceAvailability(context.Context, *CheckWorkplaceAvailabilityRequest) (*WorkplaceAvailabilityResponse, error)
	// Методы для управления парковочными местами
	GetParkingSpaces(context.Context, *GetParkingSpacesRequest) (*GetParkingSpacesResponse, error)
	GetParkingSpaceById(context.Context, *GetParkingSpaceByIdRequest) (*ParkingSpace, error)
	CreateParkingSpace(context.Context, *CreateParkingSpaceRequest) (*ParkingSpace, error)
	UpdateParkingSpace(context.Context, *UpdateParkingSpaceRequest) (*ParkingSpace, error)
	DeleteParkingSpace(context.Context, *DeleteParkingSpaceRequest) (*DeleteParkingSpaceResponse, error)
	CheckParkingSpaceAvailability(context.Context, *CheckParkingSpaceAvailabilityRequest) (*ParkingSpaceAvailabilityResponse, error)
	// Методы для управления предметами
	GetItems(context.Context, *GetItemsRequest) (*GetItemsResponse, error)
	GetItemById(context.Context, *GetItemByIdRequest) (*Item, error)
	CreateItem(context.Context, *CreateItemRequest) (*Item, error)
	UpdateItem(context.Context, *UpdateItemRequest) (*Item, error)
	DeleteItem(context.Context, *DeleteItemRequest) (*DeleteItemResponse, error)
	AttachItemToWorkplace(context.Context, *AttachItemToWorkplaceRequest) (*Item, error)
	mustEmbedUnimplementedResourceServiceServer()
}

// UnimplementedResourceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResourceServiceServer struct {
}

func (UnimplementedResourceServiceServer) GetWorkplaces(context.Context, *GetWorkplacesRequest) (*GetWorkplacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkplaces not implemented")
}
func (UnimplementedResourceServiceServer) GetWorkplaceById(context.Context, *GetWorkplaceByIdRequest) (*Workplace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkplaceById not implemented")
}
func (UnimplementedResourceServiceServer) CreateWorkplace(context.Context, *CreateWorkplaceRequest) (*Workplace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkplace not implemented")
}
func (UnimplementedResourceServiceServer) UpdateWorkplace(context.Context, *UpdateWorkplaceRequest) (*Workplace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWorkplace not implemented")
}
func (UnimplementedResourceServiceServer) DeleteWorkplace(context.Context, *DeleteWorkplaceRequest) (*DeleteWorkplaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkplace not implemented")
}
func (UnimplementedResourceServiceServer) CheckWorkplaceAvailability(context.Context, *CheckWorkplaceAvailabilityRequest) (*WorkplaceAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckWorkplaceAvailability not implemented")
}
func (UnimplementedResourceServiceServer) GetParkingSpaces(context.Context, *GetParkingSpacesRequest) (*GetParkingSpacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParkingSpaces not implemented")
}
func (UnimplementedResourceServiceServer) GetParkingSpaceById(context.Context, *GetParkingSpaceByIdRequest) (*ParkingSpace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParkingSpaceById not implemented")
}
func (UnimplementedResourceServiceServer) CreateParkingSpace(context.Context, *CreateParkingSpaceRequest) (*ParkingSpace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateParkingSpace not implemented")
}
func (UnimplementedResourceServiceServer) UpdateParkingSpace(context.Context, *UpdateParkingSpaceRequest) (*ParkingSpace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParkingSpace not implemented")
}
func (UnimplementedResourceServiceServer) DeleteParkingSpace(context.Context, *DeleteParkingSpaceRequest) (*DeleteParkingSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteParkingSpace not implemented")
}
func (UnimplementedResourceServiceServer) CheckParkingSpaceAvailability(context.Context, *CheckParkingSpaceAvailabilityRequest) (*ParkingSpaceAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckParkingSpaceAvailability not implemented")
}
func (UnimplementedResourceServiceServer) GetItems(context.Context, *GetItemsRequest) (*GetItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItems not implemented")
}
func (UnimplementedResourceServiceServer) GetItemById(context.Context, *GetItemByIdRequest) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemById not implemented")
}
func (UnimplementedResourceServiceServer) CreateItem(context.Context, *CreateItemRequest) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateItem not implemented")
}
func (UnimplementedResourceServiceServer) UpdateItem(context.Context, *UpdateItemRequest) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItem not implemented")
}
func (UnimplementedResourceServiceServer) DeleteItem(context.Context, *DeleteItemRequest) (*DeleteItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItem not implemented")
}
func (UnimplementedResourceServiceServer) AttachItemToWorkplace(context.Context, *AttachItemToWorkplaceRequest) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachItemToWorkplace not implemented")
}
func (UnimplementedResourceServiceServer) mustEmbedUnimplementedResourceServiceServer() {}

// UnsafeResourceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceServiceServer will
// result in compilation errors.
type UnsafeResourceServiceServer interface {
	mustEmbedUnimplementedResourceServiceServer()
}

func RegisterResourceServiceServer(s grpc.ServiceRegistrar, srv ResourceServiceServer) {
	s.RegisterService(&ResourceService_ServiceDesc, srv)
}

func _ResourceService_GetWorkplaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkplacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).GetWorkplaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService.ResourceService/GetWorkplaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).GetWorkplaces(ctx, req.(*GetWorkplacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_GetWorkplaceById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkplaceByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).GetWorkplaceById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService.ResourceService/GetWorkplaceById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).GetWorkplaceById(ctx, req.(*GetWorkplaceByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_CreateWorkplace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkplaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).CreateWorkplace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService.ResourceService/CreateWorkplace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).CreateWorkplace(ctx, req.(*CreateWorkplaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_UpdateWorkplace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkplaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).UpdateWorkplace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService.ResourceService/UpdateWorkplace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).UpdateWorkplace(ctx, req.(*UpdateWorkplaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_DeleteWorkplace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkplaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).DeleteWorkplace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService.ResourceService/DeleteWorkplace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).DeleteWorkplace(ctx, req.(*DeleteWorkplaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_CheckWorkplaceAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckWorkplaceAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).CheckWorkplaceAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService.ResourceService/CheckWorkplaceAvailability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).CheckWorkplaceAvailability(ctx, req.(*CheckWorkplaceAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_GetParkingSpaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParkingSpacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).GetParkingSpaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService.ResourceService/GetParkingSpaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).GetParkingSpaces(ctx, req.(*GetParkingSpacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_GetParkingSpaceById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParkingSpaceByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).GetParkingSpaceById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService.ResourceService/GetParkingSpaceById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).GetParkingSpaceById(ctx, req.(*GetParkingSpaceByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_CreateParkingSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateParkingSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).CreateParkingSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService.ResourceService/CreateParkingSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).CreateParkingSpace(ctx, req.(*CreateParkingSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_UpdateParkingSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateParkingSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).UpdateParkingSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService.ResourceService/UpdateParkingSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).UpdateParkingSpace(ctx, req.(*UpdateParkingSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_DeleteParkingSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteParkingSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).DeleteParkingSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService.ResourceService/DeleteParkingSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).DeleteParkingSpace(ctx, req.(*DeleteParkingSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_CheckParkingSpaceAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckParkingSpaceAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).CheckParkingSpaceAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService.ResourceService/CheckParkingSpaceAvailability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).CheckParkingSpaceAvailability(ctx, req.(*CheckParkingSpaceAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_GetItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).GetItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService.ResourceService/GetItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).GetItems(ctx, req.(*GetItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_GetItemById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).GetItemById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService.ResourceService/GetItemById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).GetItemById(ctx, req.(*GetItemByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_CreateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).CreateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService.ResourceService/CreateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).CreateItem(ctx, req.(*CreateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_UpdateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).UpdateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService.ResourceService/UpdateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).UpdateItem(ctx, req.(*UpdateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_DeleteItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).DeleteItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService.ResourceService/DeleteItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).DeleteItem(ctx, req.(*DeleteItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceService_AttachItemToWorkplace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachItemToWorkplaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceServiceServer).AttachItemToWorkplace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ResourceService.ResourceService/AttachItemToWorkplace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceServiceServer).AttachItemToWorkplace(ctx, req.(*AttachItemToWorkplaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceService_ServiceDesc is the grpc.ServiceDesc for ResourceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ResourceService.ResourceService",
	HandlerType: (*ResourceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWorkplaces",
			Handler:    _ResourceService_GetWorkplaces_Handler,
		},
		{
			MethodName: "GetWorkplaceById",
			Handler:    _ResourceService_GetWorkplaceById_Handler,
		},
		{
			MethodName: "CreateWorkplace",
			Handler:    _ResourceService_CreateWorkplace_Handler,
		},
		{
			MethodName: "UpdateWorkplace",
			Handler:    _ResourceService_UpdateWorkplace_Handler,
		},
		{
			MethodName: "DeleteWorkplace",
			Handler:    _ResourceService_DeleteWorkplace_Handler,
		},
		{
			MethodName: "CheckWorkplaceAvailability",
			Handler:    _ResourceService_CheckWorkplaceAvailability_Handler,
		},
		{
			MethodName: "GetParkingSpaces",
			Handler:    _ResourceService_GetParkingSpaces_Handler,
		},
		{
			MethodName: "GetParkingSpaceById",
			Handler:    _ResourceService_GetParkingSpaceById_Handler,
		},
		{
			MethodName: "CreateParkingSpace",
			Handler:    _ResourceService_CreateParkingSpace_Handler,
		},
		{
			MethodName: "UpdateParkingSpace",
			Handler:    _ResourceService_UpdateParkingSpace_Handler,
		},
		{
			MethodName: "DeleteParkingSpace",
			Handler:    _ResourceService_DeleteParkingSpace_Handler,
		},
		{
			MethodName: "CheckParkingSpaceAvailability",
			Handler:    _ResourceService_CheckParkingSpaceAvailability_Handler,
		},
		{
			MethodName: "GetItems",
			Handler:    _ResourceService_GetItems_Handler,
		},
		{
			MethodName: "GetItemById",
			Handler:    _ResourceService_GetItemById_Handler,
		},
		{
			MethodName: "CreateItem",
			Handler:    _ResourceService_CreateItem_Handler,
		},
		{
			MethodName: "UpdateItem",
			Handler:    _ResourceService_UpdateItem_Handler,
		},
		{
			MethodName: "DeleteItem",
			Handler:    _ResourceService_DeleteItem_Handler,
		},
		{
			MethodName: "AttachItemToWorkplace",
			Handler:    _ResourceService_AttachItemToWorkplace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resource.proto",
}
