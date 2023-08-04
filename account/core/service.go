package core

import (
	pb "account/generated"
)

type ServiceImplm struct {
	pb.UnimplementedAccountServiceServer
	useCase UseCase
	mapper  Mapper
}
