package data

type OrganisationPatch struct {
	Name    *string `json:"Name,omitempty"`
	Website *string `json:"Website,omitempty"`
	Phone   *string `json:"Phone,omitempty"`
}
