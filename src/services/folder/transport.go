package folder

import (
	"context"
	"ggclass_resource_service/src/logger"
	folderpb "ggclass_resource_service/src/pb"
)

type transport struct {
	folderpb.UnimplementedFolderServiceServer
}

func NewTransport() *transport {
	return &transport{}
}

func (t *transport) Create(ctx context.Context, request *folderpb.CreateFolderRequest) (*folderpb.CreateFolderResponse, error) {
	logger.Sugar().Info(request)
	return &folderpb.CreateFolderResponse{
		Message: "done 123",
	}, nil
}
