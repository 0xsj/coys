package account

import (
	pb "account/generated"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServiceImpl struct {
	pb.UnimplementedAccountServiceServer
	useCase UseCase
	mapper  Mapper
}

func NewService(useCase UseCase, mapper Mapper) pb.AccountServiceServer {
	return &ServiceImpl{useCase: useCase, mapper: mapper}
}

// method implementation in Go that handles the creation of an account using some protobuf-defined request and response messages
func (service *ServiceImpl) CreateAccount(ctx context.Context, request *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	req_phoneNumber := request.GetPhoneNumber()
	if req_phoneNumber == "" {
		return nil, status.Error(codes.InvalidArgument, "value cannot be empty")
	}

	req_role := request.GetRole()
	if req_role.String() == "" {
		return nil, status.Error(codes.InvalidArgument, "Value cannot be empty")
	}

	id, err := service.useCase.CreateAccount(ctx, req_phoneNumber, Role(req_role))
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccountResponse{Id: *id}, nil
}

func GetAccountById() {}
