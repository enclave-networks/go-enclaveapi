package data

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
