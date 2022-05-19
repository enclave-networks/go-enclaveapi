package dns

// The patch model for a DNS record. Any values not provided will not be updated.
type DnsRecordPatch struct {
	Name    string   `json:"Name,omitempty"`
	Tags    []string `json:"Tags,omitempty"`
	Systems []string `json:"Systems,omitempty"`
	Notes   string   `json:"Notes,omitempty"`
}
