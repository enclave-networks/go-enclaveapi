package policy

import (
	"time"

	"github.com/enclave-networks/go-enclaveapi/data/tag"
	"github.com/enclave-networks/go-enclaveapi/data/trustrequirement"
)

// An int backed Policy Id.
type PolicyId int

// Represents a single policy.
type Policy struct {
	Id                      PolicyId
	Created                 time.Time
	Description             string
	IsEnabled               bool
	State                   PolicyState
	SenderTags              []tag.TagReference
	ReceiverTags            []tag.TagReference
	Acls                    []PolicyAcl
	Notes                   string
	SenderTrustRequirements []trustrequirement.UsedTrustRequirement
}
