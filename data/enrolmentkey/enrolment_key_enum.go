package enrolmentkey

type EnrolmentKeyStatus string

const (
	Disabled        EnrolmentKeyStatus = "Disabled"
	Enabled         EnrolmentKeyStatus = "Enabled"
	NoUsesRemaining EnrolmentKeyStatus = "NoUsesRemaining"
)

type EnrolmentKeyType string

const (
	GeneralPurpose EnrolmentKeyType = "GeneralPurpose"
	Ephemeral      EnrolmentKeyType = "Ephemeral"
)

type EnrolmentKeyApprovalMode string

const (
	Automatic EnrolmentKeyApprovalMode = "Automatic"
	Manual    EnrolmentKeyApprovalMode = "Manual"
)

type EnrolmentKeySortOrder int

const (
	Description   EnrolmentKeySortOrder = 0
	LastUsed      EnrolmentKeySortOrder = 1
	ApprovalMode  EnrolmentKeySortOrder = 2
	UsesRemaining EnrolmentKeySortOrder = 3
)
