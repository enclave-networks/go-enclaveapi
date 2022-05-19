package dns

import "time"

// Model representing a summary of a DNS record.
type DnsZoneSummary struct {
	Id               DnsZoneId
	Name             string
	Created          time.Time
	RecordCount      int
	RecordTypeCounts map[int]string
}
