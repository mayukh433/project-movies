// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: moviesapp.proto

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

const (
	MovieService_GetMovies_FullMethodName    = "/moviesappv1.MovieService/GetMovies"
	MovieService_GetMovieById_FullMethodName = "/moviesappv1.MovieService/GetMovieById"
	MovieService_CreateMovie_FullMethodName  = "/moviesappv1.MovieService/CreateMovie"
	MovieService_UpdateMovie_FullMethodName  = "/moviesappv1.MovieService/UpdateMovie"
	MovieService_DeleteMovie_FullMethodName  = "/moviesappv1.MovieService/DeleteMovie"
)

// MovieServiceClient is the client API for MovieService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieServiceClient interface {
	GetMovies(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MovieService_GetMoviesClient, error)
	GetMovieById(ctx context.Context, in *GetMovieRequest, opts ...grpc.CallOption) (*Movie, error)
	CreateMovie(ctx context.Context, in *CreateMovieRequest, opts ...grpc.CallOption) (*Movie, error)
	UpdateMovie(ctx context.Context, in *UpdateMovieRequest, opts ...grpc.CallOption) (*Movie, error)
	DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...grpc.CallOption) (*Status, error)
}

type movieServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieServiceClient(cc grpc.ClientConnInterface) MovieServiceClient {
	return &movieServiceClient{cc}
}

func (c *movieServiceClient) GetMovies(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MovieService_GetMoviesClient, error) {
	stream, err := c.cc.NewStream(ctx, &MovieService_ServiceDesc.Streams[0], MovieService_GetMovies_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &movieServiceGetMoviesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MovieService_GetMoviesClient interface {
	Recv() (*Movie, error)
	grpc.ClientStream
}

type movieServiceGetMoviesClient struct {
	grpc.ClientStream
}

func (x *movieServiceGetMoviesClient) Recv() (*Movie, error) {
	m := new(Movie)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *movieServiceClient) GetMovieById(ctx context.Context, in *GetMovieRequest, opts ...grpc.CallOption) (*Movie, error) {
	out := new(Movie)
	err := c.cc.Invoke(ctx, MovieService_GetMovieById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) CreateMovie(ctx context.Context, in *CreateMovieRequest, opts ...grpc.CallOption) (*Movie, error) {
	out := new(Movie)
	err := c.cc.Invoke(ctx, MovieService_CreateMovie_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) UpdateMovie(ctx context.Context, in *UpdateMovieRequest, opts ...grpc.CallOption) (*Movie, error) {
	out := new(Movie)
	err := c.cc.Invoke(ctx, MovieService_UpdateMovie_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, MovieService_DeleteMovie_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieServiceServer is the server API for MovieService service.
// All implementations must embed UnimplementedMovieServiceServer
// for forward compatibility
type MovieServiceServer interface {
	GetMovies(*Empty, MovieService_GetMoviesServer) error
	GetMovieById(context.Context, *GetMovieRequest) (*Movie, error)
	CreateMovie(context.Context, *CreateMovieRequest) (*Movie, error)
	UpdateMovie(context.Context, *UpdateMovieRequest) (*Movie, error)
	DeleteMovie(context.Context, *DeleteMovieRequest) (*Status, error)
	mustEmbedUnimplementedMovieServiceServer()
}

// UnimplementedMovieServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMovieServiceServer struct {
}

func (UnimplementedMovieServiceServer) GetMovies(*Empty, MovieService_GetMoviesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMovies not implemented")
}
func (UnimplementedMovieServiceServer) GetMovieById(context.Context, *GetMovieRequest) (*Movie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieById not implemented")
}
func (UnimplementedMovieServiceServer) CreateMovie(context.Context, *CreateMovieRequest) (*Movie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMovie not implemented")
}
func (UnimplementedMovieServiceServer) UpdateMovie(context.Context, *UpdateMovieRequest) (*Movie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMovie not implemented")
}
func (UnimplementedMovieServiceServer) DeleteMovie(context.Context, *DeleteMovieRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMovie not implemented")
}
func (UnimplementedMovieServiceServer) mustEmbedUnimplementedMovieServiceServer() {}

// UnsafeMovieServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieServiceServer will
// result in compilation errors.
type UnsafeMovieServiceServer interface {
	mustEmbedUnimplementedMovieServiceServer()
}

func RegisterMovieServiceServer(s grpc.ServiceRegistrar, srv MovieServiceServer) {
	s.RegisterService(&MovieService_ServiceDesc, srv)
}

func _MovieService_GetMovies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MovieServiceServer).GetMovies(m, &movieServiceGetMoviesServer{stream})
}

type MovieService_GetMoviesServer interface {
	Send(*Movie) error
	grpc.ServerStream
}

type movieServiceGetMoviesServer struct {
	grpc.ServerStream
}

func (x *movieServiceGetMoviesServer) Send(m *Movie) error {
	return x.ServerStream.SendMsg(m)
}

func _MovieService_GetMovieById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).GetMovieById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieService_GetMovieById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).GetMovieById(ctx, req.(*GetMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_CreateMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).CreateMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieService_CreateMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).CreateMovie(ctx, req.(*CreateMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_UpdateMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).UpdateMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieService_UpdateMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).UpdateMovie(ctx, req.(*UpdateMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_DeleteMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).DeleteMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieService_DeleteMovie_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).DeleteMovie(ctx, req.(*DeleteMovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MovieService_ServiceDesc is the grpc.ServiceDesc for MovieService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "moviesappv1.MovieService",
	HandlerType: (*MovieServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovieById",
			Handler:    _MovieService_GetMovieById_Handler,
		},
		{
			MethodName: "CreateMovie",
			Handler:    _MovieService_CreateMovie_Handler,
		},
		{
			MethodName: "UpdateMovie",
			Handler:    _MovieService_UpdateMovie_Handler,
		},
		{
			MethodName: "DeleteMovie",
			Handler:    _MovieService_DeleteMovie_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMovies",
			Handler:       _MovieService_GetMovies_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "moviesapp.proto",
}
