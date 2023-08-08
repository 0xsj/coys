package account

type Account struct {
	Id          string `bson:"_id"`
	PhoneNumber string `bson:"phonenumber"`
	Role        Role   `bson:"role,omitempty"`
	Status      Status `bson:"status, omitempty"`
	CreatedAt   int64  `bson:"createdAt, omitempty"`
}
