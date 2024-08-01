package handlers

import (
	"context"
	"github.com/alpha-omega-corp/docker-svc/pkg/models"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/docker/docker/client"
	"github.com/uptrace/bun"
	"net/http"
)

type ImageService interface {
	CreateImage(ctx context.Context, req *proto.CreateImageRequest) (*proto.CreateImageResponse, error)
}

type imageService struct {
	ImageService
	db       *bun.DB
	client   *client.Client
	template TemplateHandler
}

func NewImageService(client *client.Client, db *bun.DB) ImageService {
	return &imageService{
		db:       db,
		client:   client,
		template: NewTemplateHandler(),
	}
}

func (s *imageService) CreateImage(ctx context.Context, req *proto.CreateImageRequest) (*proto.CreateImageResponse, error) {
	buffer, err := s.template.CreateDockerfile(req.Content)
	if err != nil {
		return nil, err
	}

	_, err = s.db.NewInsert().Model(&models.Dockerfile{
		Content: buffer.Bytes(),
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &proto.CreateImageResponse{
		Status: http.StatusCreated,
	}, nil
}
