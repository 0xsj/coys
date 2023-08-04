package main

import (
	"account/config"
	account "account/core"
	"account/db"
	pb "account/generated"
	"account/server"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var mongo_host = ""
var mongo_port = ""
var server_adddress = ""

func main() {

	cfg := config.Config{
		ServiceName:     "account",
		ServerAddress:   "0.0.0.0:8082",
		MongoHost:       "mongodb",
		MongoPort:       "27017",
		DatabaseName:    "account",
		CollectionItems: "items",
		ApiKey:          "your-api-key",
	}

	// cfg, err := config.LoadConfig()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	db := db.Connect(context.Background(), fmt.Sprintf("mongodb://%s:%s", cfg.MongoHost, cfg.MongoPort))
	defer db.Disconnect()

	accountRepository := account.NewRepository(db.Collection("HOST NAME", "HOST PORT"))
	accountUseCase := account.NewUseCase(accountRepository)
	accountService := account.NewService(accountUseCase, account.NewMapper())
	authInterceptor := server.NewInterceptor("Authorization", func(ctx context.Context, header string) error {
		if header != cfg.ApiKey {
			return status.Errorf(codes.Unauthenticated, "Invalid API key")
		}
		return nil
	})

	grpcServer := server.Server{Address: server_adddress}
	grpcServer.Launch(func(server *grpc.Server) {
		pb.RegisterAccountServiceServer(server, accountService)
	}, authInterceptor)
}
