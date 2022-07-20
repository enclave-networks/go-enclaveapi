package enrolmentkey

// The status of an enrolment key.
type EnrolmentKeyStatus string

const (
	Disabled        EnrolmentKeyStatus = "Disabled"
	Enabled         EnrolmentKeyStatus = "Enabled"
	NoUsesRemaining EnrolmentKeyStatus = "NoUsesRemaining"
)

// Defines the types of Enrolment Keys.
type EnrolmentKeyType string

const (
	GeneralPurpose EnrolmentKeyType = "GeneralPurpose"
	Ephemeral      EnrolmentKeyType = "Ephemeral"
)

// System Approval Mode.
type EnrolmentKeyApprovalMode string

const (
	Automatic EnrolmentKeyApprovalMode = "Automatic"
	Manual    EnrolmentKeyApprovalMode = "Manual"
)

// Enrolment Key Sort Order used when making an Enrolment Key request.
type EnrolmentKeySortOrder string

const (
	Description   EnrolmentKeySortOrder = "Description"
	LastUsed      EnrolmentKeySortOrder = "LastUsed"
	ApprovalMode  EnrolmentKeySortOrder = "ApprovalMode"
	UsesRemaining EnrolmentKeySortOrder = "UsesRemaining"
)
