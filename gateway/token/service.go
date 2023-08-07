package authentication

import (
	"context"
	pb "gateway/generated"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceImpl struct {
	pb.UnimplementedTokenServiceServer
	Client pb.TokenServiceClient
}

func NewService(address string) pb.TokenServiceServer {
	connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceImpl{Client: pb.NewTokenServiceClient(connection)}
}

func (s *ServiceImpl) GetToken(ctx context.Context, request *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	return s.Client.GenerateToken(ctx, request)
}
