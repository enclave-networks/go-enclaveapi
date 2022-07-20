package dns

import (
	"github.com/enclave-networks/go-enclaveapi/data/enrolledsystem"
	"github.com/enclave-networks/go-enclaveapi/data/tag"
)

// Model representing a summary of a DNS record.
type DnsRecordSummary struct {
	Id       DnsRecordId
	Name     string
	Type     string
	ZoneId   DnsZoneId
	ZoneName string
	Fqdn     string
	Tags     []tag.TagReference
	Systems  []enrolledsystem.SystemReference
}
