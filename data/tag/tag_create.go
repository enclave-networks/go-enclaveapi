package tag

import "github.com/enclave-networks/go-enclaveapi/data/trustrequirement"

type TagCreate struct {
	Tag               string                                `json:"Tag,omitempty"`
	Colour            string                                `json:"Colour,omitempty"`
	Notes             string                                `json:"Notes,omitempty"`
	TrustRequirements []trustrequirement.TrustRequirementId `json:"TrustRequirements,omitempty"`
}
