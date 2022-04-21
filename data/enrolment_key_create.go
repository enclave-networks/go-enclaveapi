package data

type EnrolmentKeyCreate struct {
	Type                         EnrolmentKeyType           `json:"Type,omitempty"`
	ApprovalMode                 EnrolmentKeyApprovalMode   `json:"ApprovalMode,omitempty"`
	Description                  string                     `json:"Description,omitempty"`
	UsesRemaining                int                        `json:"UsesRemaining,omitempty"`
	IpConstraints                []EnrolmentKeyIpConstraint `json:"IpConstraints,omitempty"`
	Tags                         []string                   `json:"Tags,omitempty"`
	DisconnectedRetentionMinutes int                        `json:"DisconnectedRetentionMinutes,omitempty"`
	Notes                        string                     `json:"Notes,omitempty"`
}
