package authentication

import (
	pb "authentication/generated"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ServiceImpl struct {
	pb.UnimplementedAuthenticationServiceServer
	useCase UseCase
	md      metadata.MD
}

func NewService(useCase UseCase, md metadata.MD) pb.AuthenticationServiceServer {
	return &ServiceImpl{useCase: useCase, md: md}
}

func (s *ServiceImpl) SignInByPhoneNumber(ctx context.Context, request *pb.SignInByPhoneNumberRequest) (*pb.SignInByPhoneNumberResponse, error) {
	reqPhoneNumber := request.GetPhoneNumber()
	if reqPhoneNumber == "" {
		return nil, status.Error(codes.InvalidArgument, "Value cannot be empty")
	}
	retryAt, err := s.useCase.SignInByPhoneNumber(metadata.NewOutgoingContext(ctx, s.md), reqPhoneNumber)
	if err != nil {
		return nil, err
	}
	return &pb.SignInByPhoneNumberResponse{RetryAt: *retryAt}, nil
}
