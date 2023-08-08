package account

type Status int

// we use iota here to assign incremental integer values to the constants
// status values will have 0, 1, 2 etc.
const (
	Pending Status = iota
	Active
	Suspended
)

func (s Status) String() string {
	return [...]string{"PENDING", "ACTIVE", "SUSPENDED"}[s]
}
