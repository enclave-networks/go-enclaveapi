package trustrequirement

//Defines a type of trust requirement, that generally indicates how that requirement is evaluated.
type TrustRequirementType string

const (
	UserAuthentication TrustRequirementType = "UserAuthentication"
	PublicIp           TrustRequirementType = "PublicIp"
)
