package policy

// Defines the known protocols enforced in policy ACLs.
type PolicyAclProtocol int

const (
	Any  PolicyAclProtocol = 0
	Tcp  PolicyAclProtocol = 1
	Udp  PolicyAclProtocol = 2
	Icmp PolicyAclProtocol = 3
)
