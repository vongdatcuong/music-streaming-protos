// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: protos/v1/user.proto

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error)
	GetUserDetails(ctx context.Context, in *GetUserDetailsRequest, opts ...grpc.CallOption) (*GetUserDetailsResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	PutUser(ctx context.Context, in *PutUserRequest, opts ...grpc.CallOption) (*PutUserResponse, error)
	UpdateUserStatus(ctx context.Context, in *UpdateUserStatusRequest, opts ...grpc.CallOption) (*UpdateUserStatusResponse, error)
	UpdateUserPermissions(ctx context.Context, in *UpdateUserPermissionsRequest, opts ...grpc.CallOption) (*UpdateUserPermissionsResponse, error)
	LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error)
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error) {
	out := new(GetUserListResponse)
	err := c.cc.Invoke(ctx, "/music_streaming.authentication.user.UserService/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserDetails(ctx context.Context, in *GetUserDetailsRequest, opts ...grpc.CallOption) (*GetUserDetailsResponse, error) {
	out := new(GetUserDetailsResponse)
	err := c.cc.Invoke(ctx, "/music_streaming.authentication.user.UserService/GetUserDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/music_streaming.authentication.user.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) PutUser(ctx context.Context, in *PutUserRequest, opts ...grpc.CallOption) (*PutUserResponse, error) {
	out := new(PutUserResponse)
	err := c.cc.Invoke(ctx, "/music_streaming.authentication.user.UserService/PutUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserStatus(ctx context.Context, in *UpdateUserStatusRequest, opts ...grpc.CallOption) (*UpdateUserStatusResponse, error) {
	out := new(UpdateUserStatusResponse)
	err := c.cc.Invoke(ctx, "/music_streaming.authentication.user.UserService/UpdateUserStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserPermissions(ctx context.Context, in *UpdateUserPermissionsRequest, opts ...grpc.CallOption) (*UpdateUserPermissionsResponse, error) {
	out := new(UpdateUserPermissionsResponse)
	err := c.cc.Invoke(ctx, "/music_streaming.authentication.user.UserService/UpdateUserPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error) {
	out := new(LogInResponse)
	err := c.cc.Invoke(ctx, "/music_streaming.authentication.user.UserService/LogIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/music_streaming.authentication.user.UserService/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error)
	GetUserDetails(context.Context, *GetUserDetailsRequest) (*GetUserDetailsResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	PutUser(context.Context, *PutUserRequest) (*PutUserResponse, error)
	UpdateUserStatus(context.Context, *UpdateUserStatusRequest) (*UpdateUserStatusResponse, error)
	UpdateUserPermissions(context.Context, *UpdateUserPermissionsRequest) (*UpdateUserPermissionsResponse, error)
	LogIn(context.Context, *LogInRequest) (*LogInResponse, error)
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedUserServiceServer) GetUserDetails(context.Context, *GetUserDetailsRequest) (*GetUserDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDetails not implemented")
}
func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) PutUser(context.Context, *PutUserRequest) (*PutUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserStatus(context.Context, *UpdateUserStatusRequest) (*UpdateUserStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserStatus not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserPermissions(context.Context, *UpdateUserPermissionsRequest) (*UpdateUserPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserPermissions not implemented")
}
func (UnimplementedUserServiceServer) LogIn(context.Context, *LogInRequest) (*LogInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogIn not implemented")
}
func (UnimplementedUserServiceServer) Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_streaming.authentication.user.UserService/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserList(ctx, req.(*GetUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_streaming.authentication.user.UserService/GetUserDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserDetails(ctx, req.(*GetUserDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_streaming.authentication.user.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_PutUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).PutUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_streaming.authentication.user.UserService/PutUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).PutUser(ctx, req.(*PutUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_streaming.authentication.user.UserService/UpdateUserStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserStatus(ctx, req.(*UpdateUserStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_streaming.authentication.user.UserService/UpdateUserPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserPermissions(ctx, req.(*UpdateUserPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_streaming.authentication.user.UserService/LogIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LogIn(ctx, req.(*LogInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_streaming.authentication.user.UserService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "music_streaming.authentication.user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserList",
			Handler:    _UserService_GetUserList_Handler,
		},
		{
			MethodName: "GetUserDetails",
			Handler:    _UserService_GetUserDetails_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "PutUser",
			Handler:    _UserService_PutUser_Handler,
		},
		{
			MethodName: "UpdateUserStatus",
			Handler:    _UserService_UpdateUserStatus_Handler,
		},
		{
			MethodName: "UpdateUserPermissions",
			Handler:    _UserService_UpdateUserPermissions_Handler,
		},
		{
			MethodName: "LogIn",
			Handler:    _UserService_LogIn_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _UserService_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/v1/user.proto",
}
