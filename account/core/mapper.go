/*
*	mapping between data and proto buffer
 */
package core

import pb "account/generated"

type Mapper interface {
	MessageToEntity(message *pb.Account) *Account
	EntityToMessage(entity *Account) *pb.Account
}

type MapperImpl struct{}

func NewMapper() Mapper {
	return &MapperImpl{}
}

func (m *MapperImpl) MessageToEntity(message *pb.Account) *Account {
	return &Account{
		Id:    message.GetId(),
		Email: message.GetEmail(),
		Role:  Role(message.GetRole()),
		// Status:    Status(message.GetStatus()),
		// CreatedAt: message.GetCreatedAt(),
	}
}

func (m *MapperImpl) EntityToMessage(entity *Account) *pb.Account {
	return &pb.Account{
		Id:    entity.Id,
		Email: entity.Email,
		Role:  pb.Role(entity.Role),
		// Status:    pb.Status(entity.Status),
		// CreatedAt: entity.CreatedAt,
	}
}
