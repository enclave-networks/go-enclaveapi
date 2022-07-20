package policy

// Defines the known protocols enforced in policy ACLs.
type PolicyAclProtocol string

const (
	Any  PolicyAclProtocol = "Any"
	Tcp  PolicyAclProtocol = "Tcp"
	Udp  PolicyAclProtocol = "Udp"
	Icmp PolicyAclProtocol = "Icmp"
)
