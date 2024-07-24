package handlers

import (
	"context"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/docker/docker/client"
)

type ImageService interface {
}

type imageService struct {
	ImageService
	client *client.Client
}

func NewImageService(client *client.Client) ImageService {

	return &imageService{
		client: client,
	}
}

func (s *imageService) CreateImage(ctx context.Context, req *proto.CreateImageRequest) (*proto.CreateImageResponse, error) {
	return &proto.CreateImageResponse{}, nil
}
