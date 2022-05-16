package policy

type PolicyState int

const (
	Disabled       PolicyState = 0
	Active         PolicyState = 1
	InactiveNoAcls PolicyState = 2
)
