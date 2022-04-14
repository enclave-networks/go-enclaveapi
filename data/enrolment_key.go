package data

import "time"

type EnrolmentKey struct {
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
