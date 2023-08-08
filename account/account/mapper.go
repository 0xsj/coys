/*
	Mappings between protobuffer messages and custom account type
	- Message to Entity - pointer to the custom Account type
	- Entity to Message - points to the Account type, and also points to a generated gorpc
*/

package account

import pb "account/generated"

type Mapper interface {
	MessageToEntity(message *pb.Account) *Account
	EntityToMessage(entity *Account) *pb.Account
}

// empty struct to encapsulate
type MapperImpl struct{}

func NewMapper() Mapper {
	return &MapperImpl{}
}

// mapping between Protocol Buffers messages and custom domain entities
func (m *MapperImpl) MessageToEntity(message *pb.Account) *Account {
	return &Account{
		Id:          message.GetId(),
		PhoneNumber: message.GetPhoneNumber(),
		Role:        Role(message.GetRole()),
		Status:      Status(message.GetStatus()),
		CreatedAt:   message.GetCreatedAt(),
	}
}

// mapping between the entity and the pb definition
func (m *MapperImpl) EntityToMessage(entity *Account) *pb.Account {
	return &pb.Account{
		Id:          entity.Id,
		PhoneNumber: entity.PhoneNumber,
		Role:        pb.Role(entity.Role),
		CreatedAt:   entity.CreatedAt,
	}
}
