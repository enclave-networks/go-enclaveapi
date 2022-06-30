package policy

import "github.com/enclave-networks/go-enclaveapi/data/trustrequirement"

// Data required to create a new policy.
type PolicyCreate struct {
	Description       string                                `json:"Description,omitempty"`
	IsEnabled         bool                                  `json:"IsEnabled,omitempty"`
	SenderTags        []string                              `json:"SenderTags,omitempty"`
	ReceiverTags      []string                              `json:"ReceiverTags,omitempty"`
	Acls              []PolicyAcl                           `json:"Acls,omitempty"`
	Notes             string                                `json:"Notes,omitempty"`
	TrustRequirements []trustrequirement.TrustRequirementId `json:"TrustRequirements,omitempty"`
}
