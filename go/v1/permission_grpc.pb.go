// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: protos/v1/permission.proto

package v1

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

// PermissionServiceClient is the client API for PermissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PermissionServiceClient interface {
	GetPermissionList(ctx context.Context, in *GetPermissionListRequest, opts ...grpc.CallOption) (*GetPermissionListResponse, error)
	CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*CreatePermissionResponse, error)
	PutPermission(ctx context.Context, in *PutPermissionRequest, opts ...grpc.CallOption) (*PutPermissionResponse, error)
	CheckUserPermission(ctx context.Context, in *CheckUserPermissionRequest, opts ...grpc.CallOption) (*CheckUserPermissionResponse, error)
}

type permissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPermissionServiceClient(cc grpc.ClientConnInterface) PermissionServiceClient {
	return &permissionServiceClient{cc}
}

func (c *permissionServiceClient) GetPermissionList(ctx context.Context, in *GetPermissionListRequest, opts ...grpc.CallOption) (*GetPermissionListResponse, error) {
	out := new(GetPermissionListResponse)
	err := c.cc.Invoke(ctx, "/music_streaming.authentication.permission.PermissionService/GetPermissionList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*CreatePermissionResponse, error) {
	out := new(CreatePermissionResponse)
	err := c.cc.Invoke(ctx, "/music_streaming.authentication.permission.PermissionService/CreatePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) PutPermission(ctx context.Context, in *PutPermissionRequest, opts ...grpc.CallOption) (*PutPermissionResponse, error) {
	out := new(PutPermissionResponse)
	err := c.cc.Invoke(ctx, "/music_streaming.authentication.permission.PermissionService/PutPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) CheckUserPermission(ctx context.Context, in *CheckUserPermissionRequest, opts ...grpc.CallOption) (*CheckUserPermissionResponse, error) {
	out := new(CheckUserPermissionResponse)
	err := c.cc.Invoke(ctx, "/music_streaming.authentication.permission.PermissionService/CheckUserPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PermissionServiceServer is the server API for PermissionService service.
// All implementations must embed UnimplementedPermissionServiceServer
// for forward compatibility
type PermissionServiceServer interface {
	GetPermissionList(context.Context, *GetPermissionListRequest) (*GetPermissionListResponse, error)
	CreatePermission(context.Context, *CreatePermissionRequest) (*CreatePermissionResponse, error)
	PutPermission(context.Context, *PutPermissionRequest) (*PutPermissionResponse, error)
	CheckUserPermission(context.Context, *CheckUserPermissionRequest) (*CheckUserPermissionResponse, error)
	mustEmbedUnimplementedPermissionServiceServer()
}

// UnimplementedPermissionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPermissionServiceServer struct {
}

func (UnimplementedPermissionServiceServer) GetPermissionList(context.Context, *GetPermissionListRequest) (*GetPermissionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPermissionList not implemented")
}
func (UnimplementedPermissionServiceServer) CreatePermission(context.Context, *CreatePermissionRequest) (*CreatePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePermission not implemented")
}
func (UnimplementedPermissionServiceServer) PutPermission(context.Context, *PutPermissionRequest) (*PutPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutPermission not implemented")
}
func (UnimplementedPermissionServiceServer) CheckUserPermission(context.Context, *CheckUserPermissionRequest) (*CheckUserPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserPermission not implemented")
}
func (UnimplementedPermissionServiceServer) mustEmbedUnimplementedPermissionServiceServer() {}

// UnsafePermissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PermissionServiceServer will
// result in compilation errors.
type UnsafePermissionServiceServer interface {
	mustEmbedUnimplementedPermissionServiceServer()
}

func RegisterPermissionServiceServer(s grpc.ServiceRegistrar, srv PermissionServiceServer) {
	s.RegisterService(&PermissionService_ServiceDesc, srv)
}

func _PermissionService_GetPermissionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPermissionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetPermissionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_streaming.authentication.permission.PermissionService/GetPermissionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetPermissionList(ctx, req.(*GetPermissionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_CreatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).CreatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_streaming.authentication.permission.PermissionService/CreatePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).CreatePermission(ctx, req.(*CreatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_PutPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).PutPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_streaming.authentication.permission.PermissionService/PutPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).PutPermission(ctx, req.(*PutPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_CheckUserPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).CheckUserPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_streaming.authentication.permission.PermissionService/CheckUserPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).CheckUserPermission(ctx, req.(*CheckUserPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PermissionService_ServiceDesc is the grpc.ServiceDesc for PermissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PermissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "music_streaming.authentication.permission.PermissionService",
	HandlerType: (*PermissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPermissionList",
			Handler:    _PermissionService_GetPermissionList_Handler,
		},
		{
			MethodName: "CreatePermission",
			Handler:    _PermissionService_CreatePermission_Handler,
		},
		{
			MethodName: "PutPermission",
			Handler:    _PermissionService_PutPermission_Handler,
		},
		{
			MethodName: "CheckUserPermission",
			Handler:    _PermissionService_CheckUserPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/v1/permission.proto",
}
