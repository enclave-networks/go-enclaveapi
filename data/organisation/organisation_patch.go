package organisation

// Model defining the modifiable properties for an organisation.
type OrganisationPatch struct {
	Name    string `json:"Name,omitempty"`
	Website string `json:"Website,omitempty"`
	Phone   string `json:"Phone,omitempty"`
}
