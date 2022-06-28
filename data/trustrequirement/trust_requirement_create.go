package trustrequirement

type TrustRequirementCreate struct {
	Description string
	Type        TrustRequirementType
	Notes       string
	Settings    TrustRequirementSettings
}
