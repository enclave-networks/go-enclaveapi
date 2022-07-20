package dns

import (
	"github.com/enclave-networks/go-enclaveapi/data/enrolledsystem"
	"github.com/enclave-networks/go-enclaveapi/data/tag"
)

// An int backed Dns Record Id.
type DnsRecordId int

// Detailed model representing a DNS record.
type DnsRecord struct {
	Id       DnsRecordId
	Name     string
	Type     string
	ZoneId   DnsZoneId
	ZoneName string
	Fqdn     string
	Tags     []tag.TagReference
	Systems  []enrolledsystem.SystemReference
}
