package dns

// Patch model for updating a zone.
type DnsZonePatch struct {
	Name  string `json:"Name,omitempty"`
	Notes string `json:"Notes,omitempty"`
}
