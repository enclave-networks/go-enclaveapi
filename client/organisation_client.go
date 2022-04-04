package client

import (
	"net/http"
	"net/url"

	"github.com/enclave-networks/go-enclaveapi/data"
)

type IOrganisationClient interface {
	Get() (*Organisation, error)
	Update() error
	GetOrganisationUsers() ([]*OrganisationUser, error)
}

type OrganisationClient struct {
	baseURL    *url.URL
	token      *string
	httpClient *http.Client
	currentOrg *data.AccountOrganisation
}
