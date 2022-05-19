package enclaveclient_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/enclave-networks/go-enclaveapi/data"
	"github.com/enclave-networks/go-enclaveapi/data/organisation"
	"github.com/enclave-networks/go-enclaveapi/enclave"
)

var token string = "p9rcFksNsHALkfyqyfgRzYq4AXwcuxr22CN9Mc5PG42umHPUiPhnzX7kiRfdWM3"

// https://stackoverflow.com/questions/47436263/how-to-mock-http-client-that-returns-a-json-response
func Test_when_calling_organisation_get_returns_values(t *testing.T) {
	testServer, err := createTestServer(200, nil)
	if err != nil {
		t.Error(err)
	}

	defer testServer.Close()

	enclaveClient, err := enclave.CreateClientWithUrl(token, testServer.URL)
	if err != nil {
		t.Error(err)
	}

	orgs, err := enclaveClient.GetOrgs()
	if err != nil {
		t.Error(err)
	}

	organisationClient := enclaveClient.CreateOrganisationClient(orgs[0])
	org, err := organisationClient.Get()
	if err != nil {
		t.Error(err)
	}

	if org.Id == "" {
		t.Error("org is nil")
	}
}

func Test_when_calling_organisation_returns_values(t *testing.T) {
	enclaveClient := enclave.New(token)

	orgs, _ := enclaveClient.GetOrgs()

	organisationClient := enclaveClient.CreateOrganisationClient(orgs[0])

	email := "tom.soulard+1337@enclave.io"
	err := organisationClient.InviteUser(email)

	if err != nil {
		t.Error(err)
	}

	invites, err := organisationClient.GetPendingInvites()
	if err != nil {
		t.Error(err)
	}

	if len(invites) != 1 {
		t.Errorf("expected count of 1 got count of %v", len(invites))
	}

	err = organisationClient.CancelUser(email)
	if err != nil {
		t.Error(err)
	}
}

func createTestServer(expected int, httpMatches ...*HttpMatch) (*httptest.Server, error) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		accountOrganisationBody, _ := json.Marshal(&data.AccountOrganisation{
			OrgId: organisation.OrganisationId("thing"),
		})

		if req.RequestURI == "/account/orgs" {
			res.WriteHeader(expected)
			res.Write(accountOrganisationBody)
			return
		}

		if httpMatches == nil {
			return
		}

		for _, httpMatch := range httpMatches {
			if req.RequestURI == httpMatch.Uri {
				body, _ := json.Marshal(httpMatch.Body)

				res.WriteHeader(expected)
				res.Write(body)
				return
			}
		}
	}))

	return testServer, nil
}

type HttpMatch struct {
	Body any
	Uri  string
}
