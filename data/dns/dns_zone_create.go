package dns

// The model for creating a new zone.
type DnsZoneCreate struct {
	Name  string `json:"Name,omitempty"`
	Notes string `json:"Notes,omitempty"`
}
