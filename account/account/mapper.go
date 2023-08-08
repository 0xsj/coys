/*
	Mappings between protobuffer messages and custom account type
	- Message to Entity - pointer to the custom Account type
	- Entity to Message - points to the Account type, and also points to a generated gorpc
*/

package account

import pb "account/generated"

type Mapper interface {
	MessageToEntity(message *pb.Account) *Account
	// EntityToMessage(entity *Account) *pb.Account
}

type MapperImpl struct{}

func NewMapper() Mapper {
	return &MapperImpl{}
}

func (m *MapperImpl) MessageToEntity(message *pb.Account) *Account {
	return &Account{
		Id:          message.GetId(),
		PhoneNumber: message.GetPhoneNumber(),
		Role:        Role(message.GetRole()),
		Status:      Status(message.GetStatus()),
		CreatedAt:   message.GetCreatedAt(),
	}
}
