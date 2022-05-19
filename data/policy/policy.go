package policy

import (
	"time"

	"github.com/enclave-networks/go-enclaveapi/data"
)

// An int backed Policy Id.
type PolicyId int

// Represents a single policy.
type Policy struct {
	Id           PolicyId
	Created      time.Time
	Description  string
	IsEnabled    bool
	State        PolicyState
	SenderTags   []data.TagReference
	RecieverTags []data.TagReference
	Acls         []PolicyAcl
	Notes        string
}
