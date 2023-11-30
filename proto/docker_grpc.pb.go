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
	CreateContainer(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (*CreateContainerResponse, error)
	GetContainers(ctx context.Context, in *GetContainersRequest, opts ...grpc.CallOption) (*GetContainersResponse, error)
}

type dockerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDockerServiceClient(cc grpc.ClientConnInterface) DockerServiceClient {
	return &dockerServiceClient{cc}
}

func (c *dockerServiceClient) CreateContainer(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (*CreateContainerResponse, error) {
	out := new(CreateContainerResponse)
	err := c.cc.Invoke(ctx, "/docker.DockerService/CreateContainer", in, out, opts...)
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

// DockerServiceServer is the server API for DockerService service.
// All implementations must embed UnimplementedDockerServiceServer
// for forward compatibility
type DockerServiceServer interface {
	CreateContainer(context.Context, *CreateContainerRequest) (*CreateContainerResponse, error)
	GetContainers(context.Context, *GetContainersRequest) (*GetContainersResponse, error)
	mustEmbedUnimplementedDockerServiceServer()
}

// UnimplementedDockerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDockerServiceServer struct {
}

func (UnimplementedDockerServiceServer) CreateContainer(context.Context, *CreateContainerRequest) (*CreateContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContainer not implemented")
}
func (UnimplementedDockerServiceServer) GetContainers(context.Context, *GetContainersRequest) (*GetContainersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContainers not implemented")
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

func _DockerService_CreateContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServiceServer).CreateContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.DockerService/CreateContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServiceServer).CreateContainer(ctx, req.(*CreateContainerRequest))
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

// DockerService_ServiceDesc is the grpc.ServiceDesc for DockerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DockerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "docker.DockerService",
	HandlerType: (*DockerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContainer",
			Handler:    _DockerService_CreateContainer_Handler,
		},
		{
			MethodName: "GetContainers",
			Handler:    _DockerService_GetContainers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/docker.proto",
}