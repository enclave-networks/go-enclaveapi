package policy

type PolicyAclProtocol int

const (
	Any  PolicyAclProtocol = 0
	Tcp  PolicyAclProtocol = 1
	Udp  PolicyAclProtocol = 2
	Icmp PolicyAclProtocol = 3
)
