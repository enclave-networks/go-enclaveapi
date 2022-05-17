package organisation

type OrganisationInvite struct {
	EmailAddress string
}

type OrganisationPendingInvites struct {
	Invites []OrganisationInvite
}
