package main

import (
	"account/config"
	account "account/core"
	"account/db"
	pb "account/generated"
	"account/server"
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var mongo_host = ""
var mongo_port = ""
var server_adddress = ""

func main() {

	productionMode := flag.Bool("production", false, "enable production mode")
	flag.Parse()
	cfgName := func() string {
		if *productionMode {
			return "prod"
		}
		return "dev"
	}()
	cfg, err := config.LoadConfig(cfgName)
	if err != nil {
		log.Fatal(err)
	}

	db := db.Connect(context.Background(), fmt.Sprintf("mongodb://%s:%s", mongo_host, mongo_port))
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
