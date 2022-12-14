// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: folder.proto

package folderpb

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

// FolderServiceClient is the client API for FolderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FolderServiceClient interface {
	Create(ctx context.Context, in *CreateFolderRequest, opts ...grpc.CallOption) (*CreateFolderResponse, error)
}

type folderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFolderServiceClient(cc grpc.ClientConnInterface) FolderServiceClient {
	return &folderServiceClient{cc}
}

func (c *folderServiceClient) Create(ctx context.Context, in *CreateFolderRequest, opts ...grpc.CallOption) (*CreateFolderResponse, error) {
	out := new(CreateFolderResponse)
	err := c.cc.Invoke(ctx, "/pb.FolderService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FolderServiceServer is the server API for FolderService service.
// All implementations must embed UnimplementedFolderServiceServer
// for forward compatibility
type FolderServiceServer interface {
	Create(context.Context, *CreateFolderRequest) (*CreateFolderResponse, error)
	mustEmbedUnimplementedFolderServiceServer()
}

// UnimplementedFolderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFolderServiceServer struct {
}

func (UnimplementedFolderServiceServer) Create(context.Context, *CreateFolderRequest) (*CreateFolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFolderServiceServer) mustEmbedUnimplementedFolderServiceServer() {}

// UnsafeFolderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FolderServiceServer will
// result in compilation errors.
type UnsafeFolderServiceServer interface {
	mustEmbedUnimplementedFolderServiceServer()
}

func RegisterFolderServiceServer(s grpc.ServiceRegistrar, srv FolderServiceServer) {
	s.RegisterService(&FolderService_ServiceDesc, srv)
}

func _FolderService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FolderService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).Create(ctx, req.(*CreateFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FolderService_ServiceDesc is the grpc.ServiceDesc for FolderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FolderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.FolderService",
	HandlerType: (*FolderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _FolderService_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "folder.proto",
}
