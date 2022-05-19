package dns

import (
	"github.com/enclave-networks/go-enclaveapi/data"
	"github.com/enclave-networks/go-enclaveapi/data/enrolledSystem"
)

// Model representing a summary of a DNS record.
type DnsRecordSummary struct {
	Id       DnsRecordId
	Name     string
	Type     string
	ZoneId   DnsZoneId
	ZoneName string
	Fqdn     string
	Tags     []data.TagReference
	Systems  []enrolledSystem.SystemReference
}
