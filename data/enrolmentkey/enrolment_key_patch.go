package enrolmentkey

type EnrolmentKeyPatch struct {
	Description                  string                     `json:"Description,omitempty"`
	IsEnabled                    bool                       `json:"IsEnabled,omitempty"`
	ApprovalMode                 EnrolmentKeyApprovalMode   `json:"ApprovalMode,omitempty"`
	UsesRemaining                int                        `json:"UsesRemaining,omitempty"`
	IpConstraints                []EnrolmentKeyIpConstraint `json:"IpConstraints,omitempty"`
	Tags                         []string                   `json:"Tags,omitempty"`
	DisconnectedRetentionMinutes int                        `json:"DisconnectedRetentionMinutes,omitempty"`
	Notes                        string                     `json:"Notes,omitempty"`
}
