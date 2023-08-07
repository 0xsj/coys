package authentication

import (
	"context"
	pb "gateway/generated"
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

func (s *ServiceImpl) GetToken(ctx context.Context, request *pb.GenerateTokenRequest) {
	// return s.Client.GetToken(ctx, request)
}
