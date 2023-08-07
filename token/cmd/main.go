package main

import (
	"context"
	"fmt"
	"log"
	"token/config"
	"token/database"
	pb "token/generated"
	"token/server"
	"token/token"

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

	// init services
	accountRepository := token.NewRepository(cfg, client)
	accountUseCase := token.NewUseCase(accountRepository)
	accountService := token.NewService(accountUseCase)

	grpcServer := server.Server{Address: cfg.ServerAddress}
	grpcServer.Launch(func(server *grpc.Server) {
		pb.RegisterTokenServiceServer(server, accountService)
	})
}
