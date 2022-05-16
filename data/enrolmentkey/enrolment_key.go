package enrolmentkey

import (
	"time"

	"github.com/enclave-networks/go-enclaveapi/data"
)

type EnrolmentKey struct {
	Id                           int
	Created                      time.Time
	LastUsed                     time.Time
	Type                         EnrolmentKeyType
	ApprovalMode                 EnrolmentKeyApprovalMode
	Key                          string
	Description                  string
	IsEnabled                    bool
	UsesRemaining                int64
	EnrolledCount                int64
	UnapprovedCount              int64
	Tags                         []data.TagReference
	DisconnectedRetentionMinutes int
	IpConstraints                []EnrolmentKeyIpConstraint
	Notes                        string
}

func (enrolmentKey *EnrolmentKey) Status() EnrolmentKeyStatus {
	if enrolmentKey.UsesRemaining == 0 {
		return NoUsesRemaining
	} else if enrolmentKey.IsEnabled {
		return Enabled
	} else {
		return Disabled
	}
}