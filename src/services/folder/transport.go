package folder

import (
	"context"
	folderpb "ggclass_resource_service/src/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type transport struct {
	folderpb.UnimplementedFolderServiceServer
}

func NewTransport() *transport {
	return &transport{}
}

type customError struct {
	*status.Status
	Code int `protobuf:"bytes,3,rep,name=code,proto3" json:"code,omitempty"`
}

func (t *transport) Create(ctx context.Context, request *folderpb.CreateFolderRequest) (*folderpb.CreateFolderResponse, error) {
	return nil, status.Error(codes.NotFound, "abcv")
	return &folderpb.CreateFolderResponse{
		Message: "done 123",
	}, nil
}
