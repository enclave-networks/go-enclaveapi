package organisation

import (
	"time"

	"github.com/enclave-networks/go-enclaveapi/data/account"
)

// Defines the properties of a user's membership of an organisation.
type OrganisationUser struct {
	Id           account.AccountId
	EmailAddress string
	FirstName    string
	LastName     string
	JoinDate     time.Time
	Role         int
}

// Top Level model for organisation user requests.
type OrganisationUsersTopLevel struct {
	Users []OrganisationUser
}
