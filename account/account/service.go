package account

import (
	pb "account/generated"
)

type ServiceImpl struct {
	pb.UnimplementedAccountServiceServer
	UseCase UseCase
}

func NewService() {}

func CreateAccount() {}

func GetAccountById() {}
