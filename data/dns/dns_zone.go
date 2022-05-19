package dns

import "time"

// Int backed Dns Zone Id.
type DnsZoneId int

// Detailed model representing a DNS zone.
type DnsZone struct {
	Id               DnsZoneId
	Name             string
	Created          time.Time
	RecordCount      int
	RecordTypeCounts map[string]int
	Notes            string
}
