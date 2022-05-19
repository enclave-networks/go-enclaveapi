package enrolledSystem

// Defines a system reference model.
type SystemReference struct {
	ConnectedFrom string
	Id            SystemId
	MachineName   string
	Name          string
	PlatformType  string
	State         SystemState
}
