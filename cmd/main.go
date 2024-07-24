package main

import (
	"github.com/alpha-omega-corp/docker-svc/pkg"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/alpha-omega-corp/services/config"
	svc "github.com/alpha-omega-corp/services/server"
	"github.com/docker/docker/client"
	"github.com/uptrace/bun"
	"google.golang.org/grpc"
)

func main() {
	env, err := config.NewHandler().Environment("docker")
	if err != nil {
		panic(err)
	}

	if err := svc.NewGRPC(env.Host.Url, nil, func(_ *bun.DB, grpc *grpc.Server) {
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

		proto.RegisterDockerServiceServer(grpc, pkg.NewServer(dockerClient))
	}); err != nil {
		panic(err)
	}
}
