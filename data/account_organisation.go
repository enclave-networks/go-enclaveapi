package data

type AccountOrganisation struct {
	OrgId   string
	OrgName string
	Role    UserOrganisationRole
}

type AccountOrganisationTopLevel struct {
	Orgs []AccountOrganisation
}
