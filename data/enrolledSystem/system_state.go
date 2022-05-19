package enrolledSystem

// The possible states of an enrolled system
type SystemState int

const (
	Disabled     SystemState = 0
	Disconnected SystemState = 1
	Connected    SystemState = 2
)
