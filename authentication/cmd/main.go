package main

import (
	"authentication/config"
	"authentication/server"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := server.Server{Address: cfg.ServerAddress}
	grpcServer.Launch(func(server *grpc.Server) {
		// pb.RegisterAuthenticationServiceServer(server, "")
	})
}
