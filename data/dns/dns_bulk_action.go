package dns

// Used to send a set of RecordIds to the API
type DnsBulkAction struct {
	RecordIds []DnsRecordId
}

// Response from the API when doing a bulk delete
type DnsRecordBulkDeleteResult struct {
	DnsRecordsDeleted int
}
