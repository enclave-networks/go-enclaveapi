package enrolmentkey

import (
	"time"

	"github.com/enclave-networks/go-enclaveapi/data/tag"
)

// Represents a single Enclave Enrolment Key.
type EnrolmentKeySummary struct {
	Id                           EnrolmentKeyId
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
	Tags                         []tag.TagReference
	DisconnectedRetentionMinutes int
}

func (enrolmentKeySummary *EnrolmentKeySummary) Status() EnrolmentKeyStatus {
	if enrolmentKeySummary.UsesRemaining == 0 {
		return NoUsesRemaining
	} else if enrolmentKeySummary.IsEnabled {
		return Enabled
	} else {
		return Disabled
	}
}
