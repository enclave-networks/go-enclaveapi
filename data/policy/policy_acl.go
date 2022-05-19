package policy

// Struct representing a single ACL entry for a policy.
type PolicyAcl struct {
	Protocol    PolicyAclProtocol `json:"Protocol,omitempty"`
	Ports       string            `json:"Ports,omitempty"`
	Description string            `json:"Description,omitempty"`
}
