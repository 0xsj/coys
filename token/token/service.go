package token

import (
	"context"
	pb "token/generated"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// implment the methods defined in the token service. we point to the generated pb code
type ServiceImpl struct {
	pb.UnimplementedTokenServiceServer
	useCase UseCase
}

// new constructor function, takes in a usecase as paramter and returns an instance of serviceImpl
// we create a memory address that points to the location of the specific intance of the service impl
// we can pass a refernce to the actual data this way, instead of copying the entire structure
func NewService(useCase UseCase) pb.TokenServiceServer {
	return &ServiceImpl{useCase: useCase}
}

// Generate Token service, s *ServiceImpl as the method receiver
// the receiver specifies that this method is associated with the instance of serviceimpl struct
// * indicates that method oeprates on a pointer, not the value itself.
func (s *ServiceImpl) GenerateToken(ctx context.Context, request *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	requestId := request.GetId()
	if requestId == "" {
		return nil, status.Error(codes.InvalidArgument, "Value cannot be empty")
	}

	tokenPair, err := s.useCase.GenerateTokenPair(ctx, requestId)
	if err != nil {
		return nil, err
	}
	return &pb.GenerateTokenResponse{AccessToken: tokenPair.AccessToken, RefreshToken: tokenPair.RefreshToken}, nil
}
