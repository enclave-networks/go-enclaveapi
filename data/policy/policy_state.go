package policy

// Defines the possible states of policies.
type PolicyState string

const (
	Disabled       PolicyState = "Disabled"
	Active         PolicyState = "Active"
	InactiveNoAcls PolicyState = "InactiveNoAcls"
)
