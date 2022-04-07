package data

type EnrolmentKeyStatus int

const (
	Disabled        EnrolmentKeyStatus = 0
	Enabled         EnrolmentKeyStatus = 1
	NoUsesRemaining EnrolmentKeyStatus = 2
)

type EnrolmentKeyType int

const (
	GeneralPurpose EnrolmentKeyType = 0
	Ephemeral      EnrolmentKeyType = 1
)

type EnrolmentKeyApprovalMode int

const (
	Automatic EnrolmentKeyApprovalMode = 0
	Manual    EnrolmentKeyApprovalMode = 1
)
