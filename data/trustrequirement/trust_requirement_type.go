package trustrequirement

//Defines a type of trust requirement, that generally indicates how that requirement is evaluated.
type TrustRequirementType int

const (
	UserAuthentication TrustRequirementType = 0
	PublicIp           TrustRequirementType = 1
)
