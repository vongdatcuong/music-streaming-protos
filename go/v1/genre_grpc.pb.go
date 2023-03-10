// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: protos/v1/genre.proto

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

// GenreServiceClient is the client API for GenreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GenreServiceClient interface {
	GetGenreOptionsList(ctx context.Context, in *GetGenreOptionsListRequest, opts ...grpc.CallOption) (*GetGenreOptionsListResponse, error)
}

type genreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGenreServiceClient(cc grpc.ClientConnInterface) GenreServiceClient {
	return &genreServiceClient{cc}
}

func (c *genreServiceClient) GetGenreOptionsList(ctx context.Context, in *GetGenreOptionsListRequest, opts ...grpc.CallOption) (*GetGenreOptionsListResponse, error) {
	out := new(GetGenreOptionsListResponse)
	err := c.cc.Invoke(ctx, "/music_streaming.music.genre.GenreService/GetGenreOptionsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenreServiceServer is the server API for GenreService service.
// All implementations must embed UnimplementedGenreServiceServer
// for forward compatibility
type GenreServiceServer interface {
	GetGenreOptionsList(context.Context, *GetGenreOptionsListRequest) (*GetGenreOptionsListResponse, error)
	mustEmbedUnimplementedGenreServiceServer()
}

// UnimplementedGenreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGenreServiceServer struct {
}

func (UnimplementedGenreServiceServer) GetGenreOptionsList(context.Context, *GetGenreOptionsListRequest) (*GetGenreOptionsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGenreOptionsList not implemented")
}
func (UnimplementedGenreServiceServer) mustEmbedUnimplementedGenreServiceServer() {}

// UnsafeGenreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GenreServiceServer will
// result in compilation errors.
type UnsafeGenreServiceServer interface {
	mustEmbedUnimplementedGenreServiceServer()
}

func RegisterGenreServiceServer(s grpc.ServiceRegistrar, srv GenreServiceServer) {
	s.RegisterService(&GenreService_ServiceDesc, srv)
}

func _GenreService_GetGenreOptionsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGenreOptionsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenreServiceServer).GetGenreOptionsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/music_streaming.music.genre.GenreService/GetGenreOptionsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenreServiceServer).GetGenreOptionsList(ctx, req.(*GetGenreOptionsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GenreService_ServiceDesc is the grpc.ServiceDesc for GenreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GenreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "music_streaming.music.genre.GenreService",
	HandlerType: (*GenreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGenreOptionsList",
			Handler:    _GenreService_GetGenreOptionsList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/v1/genre.proto",
}
