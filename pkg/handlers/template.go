package handlers

import (
	"bytes"
	"embed"
	"github.com/alpha-omega-corp/docker-svc/pkg/types"
	"html/template"
	"io/fs"
	"sync"
)

var (
	//go:embed templates
	embedFS      embed.FS
	unwrapFSOnce sync.Once
	unwrappedFS  fs.FS
)

type TemplateHandler interface {
	CreateDockerfile(content []byte) (*bytes.Buffer, error)
}

type templateHandler struct {
	TemplateHandler
	template *template.Template
}

func NewTemplateHandler() TemplateHandler {
	fileSys := getFS()
	tmpl, err := template.ParseFS(fileSys, "*.template")
	if err != nil {
		panic(err)
	}

	return &templateHandler{
		template: tmpl,
	}
}

func (h *templateHandler) CreateDockerfile(content []byte) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}

	if err := h.template.ExecuteTemplate(buf, "dockerfile.template", types.CreateDockerfileDto{
		Content: string(bytes.Trim(content, "\x00")),
	}); err != nil {
		return nil, err
	}

	return buf, nil
}

func getFS() fs.FS {
	unwrapFSOnce.Do(func() {
		fileSys, err := fs.Sub(embedFS, "templates")
		if err != nil {
			panic(err)
		}
		unwrappedFS = fileSys
	})
	return unwrappedFS
}
