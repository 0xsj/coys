package core

type Role int

const (
	Root Role = iota
	Creator
	Customer
)

func (r Role) String() string {
	return [...]string{"ROOT", "CREATOR", "CUSTOMER"}[r]
}
