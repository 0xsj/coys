package account

import (
	pb "authentication/generated"
)

type Repository interface {
}

type RepositoryImpl struct {
	client pb.AccountServiceClient
	mapper Mapper
}
