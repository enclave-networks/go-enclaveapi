package enclaveclient_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/enclave-networks/go-enclaveapi/data"
	"github.com/enclave-networks/go-enclaveapi/data/organisation"
	"github.com/enclave-networks/go-enclaveapi/data/tag"
	"github.com/enclave-networks/go-enclaveapi/enclave"
)

var token string = "mytoken"

// https://stackoverflow.com/questions/47436263/how-to-mock-http-client-that-returns-a-json-response
func Test_when_calling_organisation_get_returns_values(t *testing.T) {
	testServer, err := createTestServer(nil)
	if err != nil {
		t.Error(err)
	}

	defer testServer.Close()

	enclaveClient, err := enclave.NewWithUrl(token, testServer.URL)
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
	email := "test@test.com"

	testServer, err := createTestServer(
		&HttpMatch{
			Uri:      "/org/orgId/invites",
			Method:   http.MethodPost,
			Response: 200,
		},
		&HttpMatch{
			Uri:      "/org/orgId/invites",
			Method:   http.MethodDelete,
			Response: 200,
		},
		&HttpMatch{
			Uri:      "/org/orgId/invites",
			Method:   http.MethodGet,
			Response: 200,
			Body: &organisation.OrganisationPendingInvites{
				Invites: []organisation.OrganisationInvite{
					{
						EmailAddress: email,
					},
				},
			},
		},
	)

	if err != nil {
		t.Error(err)
	}

	defer testServer.Close()

	enclaveClient, err := enclave.NewWithUrl(token, testServer.URL)
	if err != nil {
		t.Error(err)
	}

	orgs, _ := enclaveClient.GetOrgs()

	organisationClient := enclaveClient.CreateOrganisationClient(orgs[0])

	err = organisationClient.InviteUser(email)

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

func Test_thing(t *testing.T) {
	token := "mKz68Bt2Pa43bs2qhagNEBs3k5mqZ2Kxpdj1nPkbf9DG3YzhSS8EbGatFMiEu36"
	enclaveClient, err := enclave.NewWithUrl(token, "http://localhost:8081")
	if err != nil {
		t.Error(err)
	}

	orgs, err := enclaveClient.GetOrgs()
	if err != nil {
		t.Error(err)
	}

	organisationClient := enclaveClient.CreateOrganisationClient(orgs[0])

	tagModel, err := organisationClient.Tags.Create(tag.TagCreate{
		Tag: "aws",
	})
	if err != nil {
		t.Error(err)
	}

	fmt.Print(tagModel.Tag)
}

func createTestServer(httpMatches ...*HttpMatch) (*httptest.Server, error) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		accountOrganisationBody, _ := json.Marshal(&data.AccountOrganisationTopLevel{
			Orgs: []data.AccountOrganisation{
				{
					OrgId: organisation.OrganisationId("orgId"),
				},
			},
		})

		if req.RequestURI == "/account/orgs" {
			res.WriteHeader(200)
			res.Write(accountOrganisationBody)
			return
		}

		if httpMatches == nil {
			return
		}

		for _, httpMatch := range httpMatches {
			if req.RequestURI == httpMatch.Uri && req.Method == httpMatch.Method {
				body, _ := json.Marshal(httpMatch.Body)

				res.WriteHeader(httpMatch.Response)
				res.Write(body)
				return
			}
		}
	}))

	return testServer, nil
}

type HttpMatch struct {
	Body     any
	Uri      string
	Response int
	Method   string
}
