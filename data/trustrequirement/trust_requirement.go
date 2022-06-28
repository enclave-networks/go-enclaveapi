package trustrequirement

import "time"

type TrustRequirement struct {
	Id             TrustRequirementId
	Description    string
	Created        time.Time
	Modified       time.Time
	Type           TrustRequirementType
	UsedInTags     int
	UsedInPolicies int
	Notes          string
	Settings       TrustRequirementSettings
}
