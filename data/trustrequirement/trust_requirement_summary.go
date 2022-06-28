package trustrequirement

import "time"

type TrustRequirementSummary struct {
	Id             TrustRequirementId
	Description    string
	Created        time.Time
	Modified       time.Time
	Type           TrustRequirementType
	UsedInTags     int
	UsedInPolicies int
	Summary        string
}
