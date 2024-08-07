package main

import (
	"github.com/alpha-omega-corp/docker-svc/pkg"
	"github.com/alpha-omega-corp/docker-svc/pkg/models"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/alpha-omega-corp/services/config"
	"github.com/alpha-omega-corp/services/core"
	"github.com/alpha-omega-corp/services/database"
	"github.com/docker/docker/client"
	"github.com/uptrace/bun"
	"google.golang.org/grpc"
)

func main() {
	env, err := config.NewHandler().Environment("docker")
	if err != nil {
		panic(err)
	}

	dbHandler := database.NewHandler(env.Host.Dsn)
	dbHandler.Database().RegisterModel(
		(*models.Dockerfile)(nil),
	)

	if err := core.NewServer(env.Host.Url, dbHandler, func(db *bun.DB, grpc *grpc.Server) {
		dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}

		defer func(cli *client.Client) {
			err := cli.Close()
			if err != nil {
				panic(err)
			}
		}(dockerClient)

		proto.RegisterDockerServiceServer(grpc, pkg.NewServer(env.Config, dockerClient, db))
	}); err != nil {
		panic(err)
	}
}
