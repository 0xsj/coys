package main

import (
	"context"
	"fmt"
	"log"
	"token/config"
	"token/database"
	"token/server"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Main running")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	client := database.NewClient(context.Background(), fmt.Sprintf("%s:%s", cfg.RedisHostname, cfg.RedisPort))
	defer client.Close()

	grpcServer := server.Server{Address: cfg.ServerAddress}
	grpcServer.Launch(func(server *grpc.Server) {
		// pb.RegisterTokenServiceServer(server)
	})

}
