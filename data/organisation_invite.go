package data

type OrganisationInvite struct {
	EmailAddress string
}

type OrganisationPendingInvites struct {
	Invites []OrganisationInvite
}
