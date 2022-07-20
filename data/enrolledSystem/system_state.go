package enrolledsystem

// The possible states of an enrolled system
type SystemState string

const (
	Disabled     SystemState = "Disabled"
	Disconnected SystemState = "Disconnected"
	Connected    SystemState = "Connected"
)
