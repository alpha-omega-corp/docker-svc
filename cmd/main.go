package main

import (
	"github.com/alpha-omega-corp/docker-svc/pkg/config"
	"github.com/alpha-omega-corp/docker-svc/pkg/services"
	"github.com/alpha-omega-corp/docker-svc/proto"
	"github.com/alpha-omega-corp/services/database"
	"github.com/alpha-omega-corp/services/server"
	"github.com/uptrace/bun"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	dbHandler := database.NewHandler(c.DSN)

	if err := server.NewGRPC(c.HOST, dbHandler, func(db *bun.DB, grpc *grpc.Server) {
		s := services.NewServer(db)
		proto.RegisterDockerServiceServer(grpc, s)
	}); err != nil {
		panic(err)
	}
}
