package policy

type PolicyAcl struct {
	Protocol    PolicyAclProtocol
	Ports       string
	Description string
}
