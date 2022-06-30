package tag

import (
	"time"

	"github.com/enclave-networks/go-enclaveapi/data/trustrequirement"
)

type DetailedTag struct {
	Tag               string
	Ref               TagRefId
	Colour            string
	LastReferenced    time.Time
	Systems           int
	Keys              int
	Policies          int
	DnsRecords        int
	Notes             string
	TrustRequirements []trustrequirement.TrustRequirementId
}
