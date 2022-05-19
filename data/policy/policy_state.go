package policy

// Defines the possible states of policies.
type PolicyState int

const (
	Disabled       PolicyState = 0
	Active         PolicyState = 1
	InactiveNoAcls PolicyState = 2
)
