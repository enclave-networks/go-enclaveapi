package organisation

// Organisation invite struct
type OrganisationInvite struct {
	EmailAddress string
}

// A struct reprisenting pending organisation invites
type OrganisationPendingInvites struct {
	Invites []OrganisationInvite
}
