package pkg

import (
	"context"
	"github.com/alpha-omega-corp/docker-svc/pkg/handlers"
	"github.com/alpha-omega-corp/docker-svc/proto"
	st "github.com/alpha-omega-corp/services/types"
	"github.com/docker/docker/client"
	"github.com/uptrace/bun"
)

type Server struct {
	proto.UnimplementedDockerServiceServer
	imageService handlers.ImageService
}

func NewServer(config st.Config, client *client.Client, db *bun.DB) *Server {
	return &Server{
		imageService: handlers.NewImageService(config, client, db),
	}
}

func (s *Server) GetImage(ctx context.Context, req *proto.GetImageRequest) (*proto.GetImageResponse, error) {
	return s.imageService.GetImage(ctx, req)
}

func (s *Server) StoreImage(ctx context.Context, req *proto.StoreImageRequest) (*proto.StoreImageResponse, error) {
	return s.imageService.StoreImage(ctx, req)
}

func (s *Server) BuildImage(ctx context.Context, req *proto.BuildImageRequest) (*proto.BuildImageResponse, error) {
	return s.imageService.BuildImage(ctx, req)
}
