/*
*	custom type role
 */
package core

type Role int

const (
	Root Role = iota
	Staff
	Courier
	Customer
)

func (r Role) String() string {
	return [...]string{"ROOT", "CREATOR", "CUSTOMER"}[r]
}
