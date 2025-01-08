// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: swapmeet.proto

package swapmeet_grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SwapmeetService_GetCategories_FullMethodName                    = "/api.SwapmeetService/GetCategories"
	SwapmeetService_CreateCategory_FullMethodName                   = "/api.SwapmeetService/CreateCategory"
	SwapmeetService_GetPublishedAdvertisements_FullMethodName       = "/api.SwapmeetService/GetPublishedAdvertisements"
	SwapmeetService_GetPublishedAdvertisementByID_FullMethodName    = "/api.SwapmeetService/GetPublishedAdvertisementByID"
	SwapmeetService_GetUserAdvertisements_FullMethodName            = "/api.SwapmeetService/GetUserAdvertisements"
	SwapmeetService_CreateAdvertisement_FullMethodName              = "/api.SwapmeetService/CreateAdvertisement"
	SwapmeetService_UpdateAdvertisement_FullMethodName              = "/api.SwapmeetService/UpdateAdvertisement"
	SwapmeetService_SubmitAdvertisementForModeration_FullMethodName = "/api.SwapmeetService/SubmitAdvertisementForModeration"
	SwapmeetService_GetModerationAdvertisements_FullMethodName      = "/api.SwapmeetService/GetModerationAdvertisements"
	SwapmeetService_PublishAdvertisement_FullMethodName             = "/api.SwapmeetService/PublishAdvertisement"
	SwapmeetService_ReturnAdvertisementToDraft_FullMethodName       = "/api.SwapmeetService/ReturnAdvertisementToDraft"
)

// SwapmeetServiceClient is the client API for SwapmeetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SwapmeetServiceClient interface {
	GetCategories(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*GetCategoriesResponse, error)
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error)
	GetPublishedAdvertisements(ctx context.Context, in *GetPublishedAdvertisementsRequest, opts ...grpc.CallOption) (*GetPublishedAdvertisementsResponse, error)
	GetPublishedAdvertisementByID(ctx context.Context, in *GetPublishedAdvertisementByIDRequest, opts ...grpc.CallOption) (*GetPublishedAdvertisementByIDResponse, error)
	GetUserAdvertisements(ctx context.Context, in *GetUserAdvertisementsRequest, opts ...grpc.CallOption) (*GetUserAdvertisementsResponse, error)
	CreateAdvertisement(ctx context.Context, in *CreateAdvertisementRequest, opts ...grpc.CallOption) (*CreateAdvertisementResponse, error)
	UpdateAdvertisement(ctx context.Context, in *UpdateAdvertisementRequest, opts ...grpc.CallOption) (*UpdateAdvertisementResponse, error)
	SubmitAdvertisementForModeration(ctx context.Context, in *SubmitAdvertisementForModerationRequest, opts ...grpc.CallOption) (*SubmitAdvertisementForModerationResponse, error)
	GetModerationAdvertisements(ctx context.Context, in *GetModerationAdvertisementsRequest, opts ...grpc.CallOption) (*GetModerationAdvertisementsResponse, error)
	PublishAdvertisement(ctx context.Context, in *PublishAdvertisementRequest, opts ...grpc.CallOption) (*PublishAdvertisementResponse, error)
	ReturnAdvertisementToDraft(ctx context.Context, in *ReturnAdvertisementToDraftRequest, opts ...grpc.CallOption) (*ReturnAdvertisementToDraftResponse, error)
}

type swapmeetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSwapmeetServiceClient(cc grpc.ClientConnInterface) SwapmeetServiceClient {
	return &swapmeetServiceClient{cc}
}

func (c *swapmeetServiceClient) GetCategories(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*GetCategoriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCategoriesResponse)
	err := c.cc.Invoke(ctx, SwapmeetService_GetCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapmeetServiceClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCategoryResponse)
	err := c.cc.Invoke(ctx, SwapmeetService_CreateCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapmeetServiceClient) GetPublishedAdvertisements(ctx context.Context, in *GetPublishedAdvertisementsRequest, opts ...grpc.CallOption) (*GetPublishedAdvertisementsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPublishedAdvertisementsResponse)
	err := c.cc.Invoke(ctx, SwapmeetService_GetPublishedAdvertisements_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapmeetServiceClient) GetPublishedAdvertisementByID(ctx context.Context, in *GetPublishedAdvertisementByIDRequest, opts ...grpc.CallOption) (*GetPublishedAdvertisementByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPublishedAdvertisementByIDResponse)
	err := c.cc.Invoke(ctx, SwapmeetService_GetPublishedAdvertisementByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapmeetServiceClient) GetUserAdvertisements(ctx context.Context, in *GetUserAdvertisementsRequest, opts ...grpc.CallOption) (*GetUserAdvertisementsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserAdvertisementsResponse)
	err := c.cc.Invoke(ctx, SwapmeetService_GetUserAdvertisements_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapmeetServiceClient) CreateAdvertisement(ctx context.Context, in *CreateAdvertisementRequest, opts ...grpc.CallOption) (*CreateAdvertisementResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAdvertisementResponse)
	err := c.cc.Invoke(ctx, SwapmeetService_CreateAdvertisement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapmeetServiceClient) UpdateAdvertisement(ctx context.Context, in *UpdateAdvertisementRequest, opts ...grpc.CallOption) (*UpdateAdvertisementResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateAdvertisementResponse)
	err := c.cc.Invoke(ctx, SwapmeetService_UpdateAdvertisement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapmeetServiceClient) SubmitAdvertisementForModeration(ctx context.Context, in *SubmitAdvertisementForModerationRequest, opts ...grpc.CallOption) (*SubmitAdvertisementForModerationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitAdvertisementForModerationResponse)
	err := c.cc.Invoke(ctx, SwapmeetService_SubmitAdvertisementForModeration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapmeetServiceClient) GetModerationAdvertisements(ctx context.Context, in *GetModerationAdvertisementsRequest, opts ...grpc.CallOption) (*GetModerationAdvertisementsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetModerationAdvertisementsResponse)
	err := c.cc.Invoke(ctx, SwapmeetService_GetModerationAdvertisements_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapmeetServiceClient) PublishAdvertisement(ctx context.Context, in *PublishAdvertisementRequest, opts ...grpc.CallOption) (*PublishAdvertisementResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishAdvertisementResponse)
	err := c.cc.Invoke(ctx, SwapmeetService_PublishAdvertisement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapmeetServiceClient) ReturnAdvertisementToDraft(ctx context.Context, in *ReturnAdvertisementToDraftRequest, opts ...grpc.CallOption) (*ReturnAdvertisementToDraftResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReturnAdvertisementToDraftResponse)
	err := c.cc.Invoke(ctx, SwapmeetService_ReturnAdvertisementToDraft_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SwapmeetServiceServer is the server API for SwapmeetService service.
// All implementations must embed UnimplementedSwapmeetServiceServer
// for forward compatibility.
type SwapmeetServiceServer interface {
	GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error)
	CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error)
	GetPublishedAdvertisements(context.Context, *GetPublishedAdvertisementsRequest) (*GetPublishedAdvertisementsResponse, error)
	GetPublishedAdvertisementByID(context.Context, *GetPublishedAdvertisementByIDRequest) (*GetPublishedAdvertisementByIDResponse, error)
	GetUserAdvertisements(context.Context, *GetUserAdvertisementsRequest) (*GetUserAdvertisementsResponse, error)
	CreateAdvertisement(context.Context, *CreateAdvertisementRequest) (*CreateAdvertisementResponse, error)
	UpdateAdvertisement(context.Context, *UpdateAdvertisementRequest) (*UpdateAdvertisementResponse, error)
	SubmitAdvertisementForModeration(context.Context, *SubmitAdvertisementForModerationRequest) (*SubmitAdvertisementForModerationResponse, error)
	GetModerationAdvertisements(context.Context, *GetModerationAdvertisementsRequest) (*GetModerationAdvertisementsResponse, error)
	PublishAdvertisement(context.Context, *PublishAdvertisementRequest) (*PublishAdvertisementResponse, error)
	ReturnAdvertisementToDraft(context.Context, *ReturnAdvertisementToDraftRequest) (*ReturnAdvertisementToDraftResponse, error)
	mustEmbedUnimplementedSwapmeetServiceServer()
}

// UnimplementedSwapmeetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSwapmeetServiceServer struct{}

func (UnimplementedSwapmeetServiceServer) GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (UnimplementedSwapmeetServiceServer) CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedSwapmeetServiceServer) GetPublishedAdvertisements(context.Context, *GetPublishedAdvertisementsRequest) (*GetPublishedAdvertisementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublishedAdvertisements not implemented")
}
func (UnimplementedSwapmeetServiceServer) GetPublishedAdvertisementByID(context.Context, *GetPublishedAdvertisementByIDRequest) (*GetPublishedAdvertisementByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublishedAdvertisementByID not implemented")
}
func (UnimplementedSwapmeetServiceServer) GetUserAdvertisements(context.Context, *GetUserAdvertisementsRequest) (*GetUserAdvertisementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAdvertisements not implemented")
}
func (UnimplementedSwapmeetServiceServer) CreateAdvertisement(context.Context, *CreateAdvertisementRequest) (*CreateAdvertisementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdvertisement not implemented")
}
func (UnimplementedSwapmeetServiceServer) UpdateAdvertisement(context.Context, *UpdateAdvertisementRequest) (*UpdateAdvertisementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAdvertisement not implemented")
}
func (UnimplementedSwapmeetServiceServer) SubmitAdvertisementForModeration(context.Context, *SubmitAdvertisementForModerationRequest) (*SubmitAdvertisementForModerationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitAdvertisementForModeration not implemented")
}
func (UnimplementedSwapmeetServiceServer) GetModerationAdvertisements(context.Context, *GetModerationAdvertisementsRequest) (*GetModerationAdvertisementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModerationAdvertisements not implemented")
}
func (UnimplementedSwapmeetServiceServer) PublishAdvertisement(context.Context, *PublishAdvertisementRequest) (*PublishAdvertisementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishAdvertisement not implemented")
}
func (UnimplementedSwapmeetServiceServer) ReturnAdvertisementToDraft(context.Context, *ReturnAdvertisementToDraftRequest) (*ReturnAdvertisementToDraftResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnAdvertisementToDraft not implemented")
}
func (UnimplementedSwapmeetServiceServer) mustEmbedUnimplementedSwapmeetServiceServer() {}
func (UnimplementedSwapmeetServiceServer) testEmbeddedByValue()                         {}

// UnsafeSwapmeetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SwapmeetServiceServer will
// result in compilation errors.
type UnsafeSwapmeetServiceServer interface {
	mustEmbedUnimplementedSwapmeetServiceServer()
}

func RegisterSwapmeetServiceServer(s grpc.ServiceRegistrar, srv SwapmeetServiceServer) {
	// If the following call pancis, it indicates UnimplementedSwapmeetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SwapmeetService_ServiceDesc, srv)
}

func _SwapmeetService_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapmeetServiceServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SwapmeetService_GetCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapmeetServiceServer).GetCategories(ctx, req.(*GetCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapmeetService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapmeetServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SwapmeetService_CreateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapmeetServiceServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapmeetService_GetPublishedAdvertisements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublishedAdvertisementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapmeetServiceServer).GetPublishedAdvertisements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SwapmeetService_GetPublishedAdvertisements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapmeetServiceServer).GetPublishedAdvertisements(ctx, req.(*GetPublishedAdvertisementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapmeetService_GetPublishedAdvertisementByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublishedAdvertisementByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapmeetServiceServer).GetPublishedAdvertisementByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SwapmeetService_GetPublishedAdvertisementByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapmeetServiceServer).GetPublishedAdvertisementByID(ctx, req.(*GetPublishedAdvertisementByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapmeetService_GetUserAdvertisements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAdvertisementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapmeetServiceServer).GetUserAdvertisements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SwapmeetService_GetUserAdvertisements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapmeetServiceServer).GetUserAdvertisements(ctx, req.(*GetUserAdvertisementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapmeetService_CreateAdvertisement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdvertisementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapmeetServiceServer).CreateAdvertisement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SwapmeetService_CreateAdvertisement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapmeetServiceServer).CreateAdvertisement(ctx, req.(*CreateAdvertisementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapmeetService_UpdateAdvertisement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAdvertisementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapmeetServiceServer).UpdateAdvertisement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SwapmeetService_UpdateAdvertisement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapmeetServiceServer).UpdateAdvertisement(ctx, req.(*UpdateAdvertisementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapmeetService_SubmitAdvertisementForModeration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitAdvertisementForModerationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapmeetServiceServer).SubmitAdvertisementForModeration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SwapmeetService_SubmitAdvertisementForModeration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapmeetServiceServer).SubmitAdvertisementForModeration(ctx, req.(*SubmitAdvertisementForModerationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapmeetService_GetModerationAdvertisements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModerationAdvertisementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapmeetServiceServer).GetModerationAdvertisements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SwapmeetService_GetModerationAdvertisements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapmeetServiceServer).GetModerationAdvertisements(ctx, req.(*GetModerationAdvertisementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapmeetService_PublishAdvertisement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishAdvertisementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapmeetServiceServer).PublishAdvertisement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SwapmeetService_PublishAdvertisement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapmeetServiceServer).PublishAdvertisement(ctx, req.(*PublishAdvertisementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapmeetService_ReturnAdvertisementToDraft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReturnAdvertisementToDraftRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapmeetServiceServer).ReturnAdvertisementToDraft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SwapmeetService_ReturnAdvertisementToDraft_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapmeetServiceServer).ReturnAdvertisementToDraft(ctx, req.(*ReturnAdvertisementToDraftRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SwapmeetService_ServiceDesc is the grpc.ServiceDesc for SwapmeetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SwapmeetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.SwapmeetService",
	HandlerType: (*SwapmeetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCategories",
			Handler:    _SwapmeetService_GetCategories_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _SwapmeetService_CreateCategory_Handler,
		},
		{
			MethodName: "GetPublishedAdvertisements",
			Handler:    _SwapmeetService_GetPublishedAdvertisements_Handler,
		},
		{
			MethodName: "GetPublishedAdvertisementByID",
			Handler:    _SwapmeetService_GetPublishedAdvertisementByID_Handler,
		},
		{
			MethodName: "GetUserAdvertisements",
			Handler:    _SwapmeetService_GetUserAdvertisements_Handler,
		},
		{
			MethodName: "CreateAdvertisement",
			Handler:    _SwapmeetService_CreateAdvertisement_Handler,
		},
		{
			MethodName: "UpdateAdvertisement",
			Handler:    _SwapmeetService_UpdateAdvertisement_Handler,
		},
		{
			MethodName: "SubmitAdvertisementForModeration",
			Handler:    _SwapmeetService_SubmitAdvertisementForModeration_Handler,
		},
		{
			MethodName: "GetModerationAdvertisements",
			Handler:    _SwapmeetService_GetModerationAdvertisements_Handler,
		},
		{
			MethodName: "PublishAdvertisement",
			Handler:    _SwapmeetService_PublishAdvertisement_Handler,
		},
		{
			MethodName: "ReturnAdvertisementToDraft",
			Handler:    _SwapmeetService_ReturnAdvertisementToDraft_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "swapmeet.proto",
}
