package data

import "github.com/enclave-networks/go-enclaveapi/data/organisation"

type AccountOrganisation struct {
	OrgId   organisation.OrganisationId
	OrgName string
	Role    UserOrganisationRole
}

type AccountOrganisationTopLevel struct {
	Orgs []AccountOrganisation
}
