package tag

import "time"

type BasicTag struct {
	Tag            string
	Ref            TagRefId
	Colour         string
	LastReferenced time.Time
	Systems        int
	Keys           int
	Policies       int
	DnsRecords     int
}
