package pkg

import (
	"context"
	"github.com/alpha-omega-corp/docker-svc/pkg/handlers"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/docker/docker/client"
	"github.com/uptrace/bun"
)

type Server struct {
	proto.UnimplementedDockerServiceServer
	imageService handlers.ImageService
}

func NewServer(db *bun.DB, client *client.Client) *Server {
	return &Server{
		imageService: handlers.NewImageService(client, db),
	}
}

func (s *Server) CreateImage(ctx context.Context, req *proto.CreateImageRequest) (*proto.CreateImageResponse, error) {
	return s.imageService.CreateImage(ctx, req)
}
