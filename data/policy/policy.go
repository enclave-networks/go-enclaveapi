package policy

import (
	"time"

	"github.com/enclave-networks/go-enclaveapi/data"
	"github.com/enclave-networks/go-enclaveapi/data/trustrequirement"
)

// An int backed Policy Id.
type PolicyId int

// Represents a single policy.
type Policy struct {
	Id                PolicyId
	Created           time.Time
	Description       string
	IsEnabled         bool
	State             PolicyState
	SenderTags        []data.TagReference
	ReceiverTags      []data.TagReference
	Acls              []PolicyAcl
	Notes             string
	TrustRequirements []trustrequirement.TrustRequirementId
}
