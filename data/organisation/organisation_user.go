package organisation

import "time"

type OrganisationUser struct {
	Id           string
	EmailAddress string
	FirstName    string
	LastName     string
	JoinDate     time.Time
	Role         int
}

type OrganisationUsersTopLevel struct {
	Users []OrganisationUser
}
