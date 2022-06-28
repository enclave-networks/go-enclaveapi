package trustrequirement

type BulkTrustRequirementDeleteAction struct {
	RequirementIds []TrustRequirementId
}

type BulkTrustRequirementDeleteResult struct {
	RequirementsDeleted int
}
