package dns

// Used to create a DNS record leave fields blank that you do not wish to use
type DnsRecordCreate struct {
	Name    string    `json:"Name,omitempty"`
	ZoneId  DnsZoneId `json:"ZoneId,omitempty"`
	Type    string    `json:"Type,omitempty"`
	Tags    []string  `json:"Tags,omitempty"`
	Systems []string  `json:"Systems,omitempty"`
	Notes   string    `json:"Notes,omitempty"`
}
