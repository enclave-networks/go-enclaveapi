package policy

type PolicyAcl struct {
	Protocol    PolicyAclProtocol `json:"Protocol,omitempty"`
	Ports       string            `json:"Ports,omitempty"`
	Description string            `json:"Description,omitempty"`
}
