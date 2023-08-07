package main

import (
	"gateway/config"
	pb "gateway/generated"
	"gateway/server"
	token "gateway/token"
	"log"

	"google.golang.org/grpc"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := server.Server{Address: cfg.ServerAddress}
	tokenService := token.NewService(cfg.TokenAddress)
	grpcServer.Launch(func(server *grpc.Server) {
		pb.RegisterTokenServiceServer(server, tokenService)
	})
}
