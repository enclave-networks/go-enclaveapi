package data

type EnrolmentKeyCreate struct {
	Type                         *EnrolmentKeyType
	ApprovalMode                 *EnrolmentKeyApprovalMode
	Description                  *string
	UsesRemaining                *int
	IpConstraints                []*EnrolmentKeyIpConstraint
	Tags                         []*string
	DisconnectedRetentionMinutes *int
	Notes                        *string
}
