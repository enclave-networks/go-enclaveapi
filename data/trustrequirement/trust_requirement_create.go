package trustrequirement

type TrustRequirementCreate struct {
	Description string                   `json:"Description,omitempty"`
	Type        TrustRequirementType     `json:"Type,omitempty"`
	Notes       string                   `json:"Notes,omitempty"`
	Settings    TrustRequirementSettings `json:"Settings,omitempty"`
}
