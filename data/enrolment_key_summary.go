package data

import "time"

type EnrolmentKeySummary struct {
	Id                           string
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
	Tags                         []TagReference
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
