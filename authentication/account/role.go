package account

type Role int

const (
	Admin Role = iota
	Creator
	User
)

// @[r] - we can access an element from the sliced string
func (r Role) String() string {
	return [...]string{"ADMIN", "CREATOR", "USER"}[r]
}
