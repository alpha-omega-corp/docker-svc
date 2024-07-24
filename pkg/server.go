package pkg

import (
	"context"
	"github.com/alpha-omega-corp/docker-svc/pkg/handlers"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/docker/docker/client"
)

type Server struct {
	proto.UnimplementedDockerServiceServer

	imageService handlers.ImageService
}

func NewServer(client *client.Client) *Server {

	return &Server{
		imageService: handlers.NewImageService(client),
	}
}

func (s *Server) CreateImage(ctx context.Context, req *proto.CreateImageRequest) (*proto.CreateImageResponse, error) {
	return &proto.CreateImageResponse{}, nil
}
