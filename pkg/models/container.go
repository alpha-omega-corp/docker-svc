package models

import (
	"github.com/alpha-omega-corp/github-svc/pkg/types"
	dockerTypes "github.com/docker/docker/api/types"
	"github.com/uptrace/bun"
	"os"
)

type ContainerPackage struct {
	bun.BaseModel `bun:"table:packages,alias:pkg"`

	ID         int64                   `json:"id" bun:"id,pk,autoincrement"`
	Name       string                  `json:"name" bun:"name"`
	Tag        string                  `json:"tag" bun:"tag"`
	Pushed     bool                    `bun:"pushed"`
	Dockerfile []byte                  `bun:"-"`
	Makefile   []byte                  `bun:"-"`
	Git        *types.GitPackage       `bun:"-"`
	Containers []dockerTypes.Container `bun:"-"`
}

func (h *ContainerPackage) GetFile(t string) []byte {
	file, err := os.ReadFile("template/" + h.Name + "/" + t)
	if err != nil {
		panic(err)
	}

	return file
}
