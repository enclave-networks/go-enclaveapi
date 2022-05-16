package policy

import (
	"time"

	"github.com/enclave-networks/go-enclaveapi/data"
)

type Policy struct {
	Id           int
	Created      time.Time
	Description  string
	IsEnabled    bool
	State        PolicyState
	SenderTags   []data.TagReference
	RecieverTags []data.TagReference
	Acls         []PolicyAcl
	Notes        string
}
