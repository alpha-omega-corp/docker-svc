// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: proto/docker.proto

package proto

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

// DockerServiceClient is the client API for DockerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DockerServiceClient interface {
	PushPackage(ctx context.Context, in *PushPackageRequest, opts ...grpc.CallOption) (*PushPackageResponse, error)
	ContainerPackage(ctx context.Context, in *ContainerPackageRequest, opts ...grpc.CallOption) (*ContainerPackageResponse, error)
	DeleteContainer(ctx context.Context, in *DeleteContainerRequest, opts ...grpc.CallOption) (*DeleteContainerResponse, error)
	GetContainers(ctx context.Context, in *GetContainersRequest, opts ...grpc.CallOption) (*GetContainersResponse, error)
	GetContainerLogs(ctx context.Context, in *GetContainerLogsRequest, opts ...grpc.CallOption) (*GetContainerLogsResponse, error)
	GetPackages(ctx context.Context, in *GetPackagesRequest, opts ...grpc.CallOption) (*GetPackagesResponse, error)
	GetPackage(ctx context.Context, in *GetPackageRequest, opts ...grpc.CallOption) (*GetPackageResponse, error)
	GetPackageFile(ctx context.Context, in *GetPackageFileRequest, opts ...grpc.CallOption) (*GetPackageFileResponse, error)
	CreatePackage(ctx context.Context, in *CreatePackageRequest, opts ...grpc.CallOption) (*CreatePackageResponse, error)
	CreatePackageVersion(ctx context.Context, in *CreatePackageVersionRequest, opts ...grpc.CallOption) (*CreatePackageVersionResponse, error)
	DeletePackage(ctx context.Context, in *DeletePackageRequest, opts ...grpc.CallOption) (*DeletePackageResponse, error)
}

type dockerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDockerServiceClient(cc grpc.ClientConnInterface) DockerServiceClient {
	return &dockerServiceClient{cc}
}

func (c *dockerServiceClient) PushPackage(ctx context.Context, in *PushPackageRequest, opts ...grpc.CallOption) (*PushPackageResponse, error) {
	out := new(PushPackageResponse)
	err := c.cc.Invoke(ctx, "/docker.DockerService/PushPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerServiceClient) ContainerPackage(ctx context.Context, in *ContainerPackageRequest, opts ...grpc.CallOption) (*ContainerPackageResponse, error) {
	out := new(ContainerPackageResponse)
	err := c.cc.Invoke(ctx, "/docker.DockerService/ContainerPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerServiceClient) DeleteContainer(ctx context.Context, in *DeleteContainerRequest, opts ...grpc.CallOption) (*DeleteContainerResponse, error) {
	out := new(DeleteContainerResponse)
	err := c.cc.Invoke(ctx, "/docker.DockerService/DeleteContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerServiceClient) GetContainers(ctx context.Context, in *GetContainersRequest, opts ...grpc.CallOption) (*GetContainersResponse, error) {
	out := new(GetContainersResponse)
	err := c.cc.Invoke(ctx, "/docker.DockerService/GetContainers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerServiceClient) GetContainerLogs(ctx context.Context, in *GetContainerLogsRequest, opts ...grpc.CallOption) (*GetContainerLogsResponse, error) {
	out := new(GetContainerLogsResponse)
	err := c.cc.Invoke(ctx, "/docker.DockerService/GetContainerLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerServiceClient) GetPackages(ctx context.Context, in *GetPackagesRequest, opts ...grpc.CallOption) (*GetPackagesResponse, error) {
	out := new(GetPackagesResponse)
	err := c.cc.Invoke(ctx, "/docker.DockerService/GetPackages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerServiceClient) GetPackage(ctx context.Context, in *GetPackageRequest, opts ...grpc.CallOption) (*GetPackageResponse, error) {
	out := new(GetPackageResponse)
	err := c.cc.Invoke(ctx, "/docker.DockerService/GetPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerServiceClient) GetPackageFile(ctx context.Context, in *GetPackageFileRequest, opts ...grpc.CallOption) (*GetPackageFileResponse, error) {
	out := new(GetPackageFileResponse)
	err := c.cc.Invoke(ctx, "/docker.DockerService/GetPackageFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerServiceClient) CreatePackage(ctx context.Context, in *CreatePackageRequest, opts ...grpc.CallOption) (*CreatePackageResponse, error) {
	out := new(CreatePackageResponse)
	err := c.cc.Invoke(ctx, "/docker.DockerService/CreatePackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerServiceClient) CreatePackageVersion(ctx context.Context, in *CreatePackageVersionRequest, opts ...grpc.CallOption) (*CreatePackageVersionResponse, error) {
	out := new(CreatePackageVersionResponse)
	err := c.cc.Invoke(ctx, "/docker.DockerService/CreatePackageVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerServiceClient) DeletePackage(ctx context.Context, in *DeletePackageRequest, opts ...grpc.CallOption) (*DeletePackageResponse, error) {
	out := new(DeletePackageResponse)
	err := c.cc.Invoke(ctx, "/docker.DockerService/DeletePackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DockerServiceServer is the server API for DockerService service.
// All implementations must embed UnimplementedDockerServiceServer
// for forward compatibility
type DockerServiceServer interface {
	PushPackage(context.Context, *PushPackageRequest) (*PushPackageResponse, error)
	ContainerPackage(context.Context, *ContainerPackageRequest) (*ContainerPackageResponse, error)
	DeleteContainer(context.Context, *DeleteContainerRequest) (*DeleteContainerResponse, error)
	GetContainers(context.Context, *GetContainersRequest) (*GetContainersResponse, error)
	GetContainerLogs(context.Context, *GetContainerLogsRequest) (*GetContainerLogsResponse, error)
	GetPackages(context.Context, *GetPackagesRequest) (*GetPackagesResponse, error)
	GetPackage(context.Context, *GetPackageRequest) (*GetPackageResponse, error)
	GetPackageFile(context.Context, *GetPackageFileRequest) (*GetPackageFileResponse, error)
	CreatePackage(context.Context, *CreatePackageRequest) (*CreatePackageResponse, error)
	CreatePackageVersion(context.Context, *CreatePackageVersionRequest) (*CreatePackageVersionResponse, error)
	DeletePackage(context.Context, *DeletePackageRequest) (*DeletePackageResponse, error)
	mustEmbedUnimplementedDockerServiceServer()
}

// UnimplementedDockerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDockerServiceServer struct {
}

func (UnimplementedDockerServiceServer) PushPackage(context.Context, *PushPackageRequest) (*PushPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushPackage not implemented")
}
func (UnimplementedDockerServiceServer) ContainerPackage(context.Context, *ContainerPackageRequest) (*ContainerPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContainerPackage not implemented")
}
func (UnimplementedDockerServiceServer) DeleteContainer(context.Context, *DeleteContainerRequest) (*DeleteContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContainer not implemented")
}
func (UnimplementedDockerServiceServer) GetContainers(context.Context, *GetContainersRequest) (*GetContainersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContainers not implemented")
}
func (UnimplementedDockerServiceServer) GetContainerLogs(context.Context, *GetContainerLogsRequest) (*GetContainerLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContainerLogs not implemented")
}
func (UnimplementedDockerServiceServer) GetPackages(context.Context, *GetPackagesRequest) (*GetPackagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackages not implemented")
}
func (UnimplementedDockerServiceServer) GetPackage(context.Context, *GetPackageRequest) (*GetPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackage not implemented")
}
func (UnimplementedDockerServiceServer) GetPackageFile(context.Context, *GetPackageFileRequest) (*GetPackageFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackageFile not implemented")
}
func (UnimplementedDockerServiceServer) CreatePackage(context.Context, *CreatePackageRequest) (*CreatePackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePackage not implemented")
}
func (UnimplementedDockerServiceServer) CreatePackageVersion(context.Context, *CreatePackageVersionRequest) (*CreatePackageVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePackageVersion not implemented")
}
func (UnimplementedDockerServiceServer) DeletePackage(context.Context, *DeletePackageRequest) (*DeletePackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePackage not implemented")
}
func (UnimplementedDockerServiceServer) mustEmbedUnimplementedDockerServiceServer() {}

// UnsafeDockerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DockerServiceServer will
// result in compilation errors.
type UnsafeDockerServiceServer interface {
	mustEmbedUnimplementedDockerServiceServer()
}

func RegisterDockerServiceServer(s grpc.ServiceRegistrar, srv DockerServiceServer) {
	s.RegisterService(&DockerService_ServiceDesc, srv)
}

func _DockerService_PushPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServiceServer).PushPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.DockerService/PushPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServiceServer).PushPackage(ctx, req.(*PushPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerService_ContainerPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServiceServer).ContainerPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.DockerService/ContainerPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServiceServer).ContainerPackage(ctx, req.(*ContainerPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerService_DeleteContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServiceServer).DeleteContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.DockerService/DeleteContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServiceServer).DeleteContainer(ctx, req.(*DeleteContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerService_GetContainers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContainersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServiceServer).GetContainers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.DockerService/GetContainers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServiceServer).GetContainers(ctx, req.(*GetContainersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerService_GetContainerLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContainerLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServiceServer).GetContainerLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.DockerService/GetContainerLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServiceServer).GetContainerLogs(ctx, req.(*GetContainerLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerService_GetPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPackagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServiceServer).GetPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.DockerService/GetPackages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServiceServer).GetPackages(ctx, req.(*GetPackagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerService_GetPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServiceServer).GetPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.DockerService/GetPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServiceServer).GetPackage(ctx, req.(*GetPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerService_GetPackageFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPackageFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServiceServer).GetPackageFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.DockerService/GetPackageFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServiceServer).GetPackageFile(ctx, req.(*GetPackageFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerService_CreatePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServiceServer).CreatePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.DockerService/CreatePackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServiceServer).CreatePackage(ctx, req.(*CreatePackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerService_CreatePackageVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePackageVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServiceServer).CreatePackageVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.DockerService/CreatePackageVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServiceServer).CreatePackageVersion(ctx, req.(*CreatePackageVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerService_DeletePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServiceServer).DeletePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.DockerService/DeletePackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServiceServer).DeletePackage(ctx, req.(*DeletePackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DockerService_ServiceDesc is the grpc.ServiceDesc for DockerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DockerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "docker.DockerService",
	HandlerType: (*DockerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushPackage",
			Handler:    _DockerService_PushPackage_Handler,
		},
		{
			MethodName: "ContainerPackage",
			Handler:    _DockerService_ContainerPackage_Handler,
		},
		{
			MethodName: "DeleteContainer",
			Handler:    _DockerService_DeleteContainer_Handler,
		},
		{
			MethodName: "GetContainers",
			Handler:    _DockerService_GetContainers_Handler,
		},
		{
			MethodName: "GetContainerLogs",
			Handler:    _DockerService_GetContainerLogs_Handler,
		},
		{
			MethodName: "GetPackages",
			Handler:    _DockerService_GetPackages_Handler,
		},
		{
			MethodName: "GetPackage",
			Handler:    _DockerService_GetPackage_Handler,
		},
		{
			MethodName: "GetPackageFile",
			Handler:    _DockerService_GetPackageFile_Handler,
		},
		{
			MethodName: "CreatePackage",
			Handler:    _DockerService_CreatePackage_Handler,
		},
		{
			MethodName: "CreatePackageVersion",
			Handler:    _DockerService_CreatePackageVersion_Handler,
		},
		{
			MethodName: "DeletePackage",
			Handler:    _DockerService_DeletePackage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/docker.proto",
}
