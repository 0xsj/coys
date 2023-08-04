package main

import (
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
	db := db.Connect(context.Background(), fmt.Sprintf("mongodb://%s:%s", mongo_host, mongo_port))
	defer db.Disconnect()

	accountRepository := account.NewRepository(db.Collection("HOST NAME", "HOST PORT"))
	accountUseCase := account.NewUseCase(accountRepository)
	accountInterceptor := server.NewInterceptor("")
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
