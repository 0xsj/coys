package core

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

func (s *ServiceImpl) CreateAccount(ctx context.Context, request *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	reqEmail := request.GetEmail()
	if reqEmail == "" {
		return nil, status.Error(codes.InvalidArgument, "Value cannot be empty")
	}
	reqRole := request.GetRole()
	if reqRole.String() == "" {
		return nil, status.Error(codes.InvalidArgument, "Value cannot be empty")
	}
	id, err := s.useCase.CreateAccount(ctx, reqEmail, Role(reqRole))
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccountResponse{Id: *id}, nil
}

func (s *ServiceImpl) GetAccountById(ctx context.Context, request *pb.GetAccountByIdRequest) (*pb.GetAccountByIdResponse, error) {
	reqId := request.GetId()
	if reqId == "" {
		return nil, status.Error(codes.InvalidArgument, "Value cannot be empty")
	}
	account, err := s.useCase.GetAccountById(ctx, reqId)
	if err != nil {
		return nil, err
	}
	if account != nil {
		return &pb.GetAccountByIdResponse{Account: s.mapper.EntityToMessage(account)}, nil
	}
	return nil, status.Error(codes.NotFound, codes.NotFound.String())
}
