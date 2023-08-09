package main

import (
	"account/account"
	"account/config"
	"account/database"
	pb "account/generated"
	"account/server"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db := database.Connect(context.Background(), fmt.Sprintf("mongodb://%s:%s", cfg.MongoHostname, cfg.MongoPort))
	defer db.Disconenct()

	accountRepository := account.NewRepository(db.Collection(cfg.DatabaseName, cfg.CollectionItems))
	accountUseCase := account.NewUseCase(accountRepository)
	accountService := account.NewService(accountUseCase, account.NewMapper())

	grpcServer := server.Server{Address: cfg.ServerAddress}
	grpcServer.Launch(func(server *grpc.Server) {
		pb.RegisterAccountServiceServer(server, accountService)
	})
}
