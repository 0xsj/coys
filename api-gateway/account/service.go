package account

import (
	pb "api-gateway/generated" // Import the generated protobuf package
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceImpl struct {
	pb.UnimplementedAuthenticationServiceServer
	Client pb.AuthenticationServiceClient
}

func NewService(address string) pb.AuthenticationServiceServer {
	connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceImpl{Client: pb.NewAuthenticationServiceClient(connection)}
}
