package core

type Account struct {
	Id        string
	Email     string
	Role      Role
	Status    Status
	CreatedAt int64
}
