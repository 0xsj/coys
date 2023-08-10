package authentication

import (
	pb "authentication/generated"

	"google.golang.org/grpc/metadata"
)

type ServiceImpl struct {
	pb.UnimplementedAuthenticationServiceServer
	useCase UseCase
	md      metadata.MD
}

func NewService(useCase UseCase, md metadata.MD) pb.AuthenticationServiceServer {
	return &ServiceImpl{useCase: useCase, md: md}
}
