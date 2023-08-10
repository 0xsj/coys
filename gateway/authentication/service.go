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

func (s *ServiceImpl) SignInByPhoneNumber(ctx context.Context, request *pb.SignInByPhoneNumberRequest) (*pb.SignInByPhoneNumberResponse, error) {
	return s.Client.SignInByPhoneNumber(ctx, request)
}

func (s *ServiceImpl) VerifyAccess(ctx context.Context, request *pb.VerifyAccessRequest) (*pb.VerifyAccessResponse, error) {
	return s.Client.VerifyAccess(ctx, request)
}
