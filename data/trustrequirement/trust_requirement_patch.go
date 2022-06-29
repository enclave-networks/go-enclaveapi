package trustrequirement

type TrustRequirementPatch struct {
	Description string                   `json:"Description,omitempty"`
	Notes       string                   `json:"Notes,omitempty"`
	Settings    TrustRequirementSettings `json:"Settings,omitempty"`
}
