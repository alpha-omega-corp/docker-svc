package handlers

import (
	"bytes"
	"context"
	"fmt"
	"github.com/alpha-omega-corp/docker-svc/pkg/models"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/alpha-omega-corp/services/core"
	st "github.com/alpha-omega-corp/services/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/uptrace/bun"
	"net/http"
	"os"
)

type ImageService interface {
	GetImage(ctx context.Context, req *proto.GetImageRequest) (*proto.GetImageResponse, error)
	StoreImage(ctx context.Context, req *proto.StoreImageRequest) (*proto.StoreImageResponse, error)
	BuildImage(ctx context.Context, req *proto.BuildImageRequest) (*proto.BuildImageResponse, error)
}

type imageService struct {
	ImageService
	config   st.Config
	client   *client.Client
	template TemplateHandler
	db       *bun.DB
}

func NewImageService(config st.Config, client *client.Client, db *bun.DB) ImageService {
	return &imageService{
		db:       db,
		client:   client,
		config:   config,
		template: NewTemplateHandler(config),
	}
}

func (s *imageService) GetImage(ctx context.Context, req *proto.GetImageRequest) (*proto.GetImageResponse, error) {
	dockerfile := new(models.Dockerfile)
	err := s.db.NewSelect().Model(dockerfile).Where("name = ?", req.Name).Scan(ctx)
	if err != nil {
		return nil, err
	}

	images, err := s.client.ImageList(ctx, image.ListOptions{})
	if err != nil {
		return nil, err
	}

	imgSlice := make([]*models.Image, len(images))
	for index, item := range images {
		imgSlice[index] = &models.Image{
			ID:         item.ID,
			Containers: item.Containers,
			Created:    item.Created,
			Size:       item.Size,
		}

		fmt.Print(imgSlice[index])
	}

	return &proto.GetImageResponse{
		Status:  http.StatusOK,
		Content: dockerfile.Content,
	}, nil
}

func (s *imageService) StoreImage(ctx context.Context, req *proto.StoreImageRequest) (*proto.StoreImageResponse, error) {
	dockerfile := new(models.Dockerfile)
	err := s.db.NewSelect().Model(dockerfile).Where("name = ?", req.Name).Scan(ctx)

	content := bytes.Trim(req.Content, "\x00")

	if err != nil {
		_, err = s.db.NewInsert().Model(&models.Dockerfile{
			Name:    req.Name,
			Content: content,
		}).Exec(ctx)

		if err != nil {
			return nil, err
		}
	} else {
		_, err = s.db.
			NewUpdate().
			Model(dockerfile).
			Set("content = ?", content).
			Where("name = ?", req.Name).
			Exec(ctx)
		if err != nil {
			return nil, err
		}
	}

	return &proto.StoreImageResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *imageService) BuildImage(ctx context.Context, req *proto.BuildImageRequest) (*proto.BuildImageResponse, error) {
	path := s.config.Viper.GetString("build_path") + req.Name
	tag := "latest"

	makeFile, err := s.template.CreateDockerBuild(req.Name, tag)
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return nil, err
	}

	dockerFile := new(models.Dockerfile)
	if err := s.db.NewSelect().Model(dockerFile).Where("name = ?", req.Name).Scan(ctx); err != nil {
		return nil, err
	}

	if err = os.WriteFile(path+"/Dockerfile", dockerFile.Content, 0644); err != nil {
		return nil, err
	}

	if err = os.WriteFile(path+"/Makefile", makeFile.Bytes(), 0644); err != nil {
		return nil, err
	}

	if err = core.Command("make", path, "build"); err != nil {
		return nil, err
	}

	return nil, nil
}
