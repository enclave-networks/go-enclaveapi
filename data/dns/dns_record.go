package dns

import (
	"github.com/enclave-networks/go-enclaveapi/data"
	"github.com/enclave-networks/go-enclaveapi/data/enrolledSystem"
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
	Tags     []data.TagReference
	Systems  []enrolledSystem.SystemReference
}
