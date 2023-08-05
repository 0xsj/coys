package main

import (
	"authentication/account"
	"authentication/config"
	"authentication/confirmation"
	authentication "authentication/core"
	pb "authentication/generated"
	"authentication/server"
	"authentication/token"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	cfg := config.Config{
		ServiceName:         "authentication",
		ServerAddress:       "0.0.0.0:8080",
		RedisHostname:       "0.0.0.0",
		RedisPort:           "6379",
		MongoHostname:       "0.0.0.0",
		MongoPort:           "27017",
		DatabaseName:        "authentication",
		CollectionItems:     "items",
		TokenAddress:        "0.0.0.0:8081",
		AccountAddress:      "0.0.0.0:8082",
		ConfirmationAddress: "0.0.0.0:8083",
		ApiKey:              "aW52b2x2ZWRzcHJpbmdjb25zb25hbnRzaGFsbG1lbWJlcmJ5bW9ua2V5ZXhlcmNpc2U=",
	}

	tokenClient := token.NewClient(cfg.TokenAddress)
	tokenRepository := token.NewRepository(tokenClient)

	accountClient := account.NewClient(cfg.AccountAddress)
	accountRepository := account.NewRepository(accountClient, account.NewMapper())

	confirmationClient := confirmation.NewClient(cfg.ConfirmationAddress)
	confirmationRepository := confirmation.NewRepository(confirmationClient)

	authenticationUseCase := authentication.NewUseCase(accountRepository, confirmationRepository, tokenRepository)

	md := metadata.New(map[string]string{"Authorization": cfg.ApiKey})
	authenticationService := authentication.NewService(authenticationUseCase, md)
	grpcServer := server.Server{Address: cfg.ServerAddress}
	grpcServer.Launch(func(server *grpc.Server) {
		pb.RegisterAuthenticationServiceServer(server, authenticationService)
	})
}
