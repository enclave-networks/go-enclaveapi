package enclaveclient_test

import (
	"testing"

	"github.com/enclave-networks/go-enclaveapi/client"
)

var token string = "p9rcFksNsHALkfyqyfgRzYq4AXwcuxr22CN9Mc5PG42umHPUiPhnzX7kiRfdWM3"

// https://stackoverflow.com/questions/47436263/how-to-mock-http-client-that-returns-a-json-response
func Test_when_calling_organisation_get_returns_values(t *testing.T) {
	enclaveClient := client.CreateClient(&token)

	orgs, _ := enclaveClient.GetOrgs()

	organisationClient := enclaveClient.CreateOrganisationClient(orgs[0])
	org, err := organisationClient.Get()
	if err != nil {
		t.Error(err)
	}

	if org == nil {
		t.Error("org is nil")
	}
}

func Test_when_calling_organisation_returns_values(t *testing.T) {
	enclaveClient := client.CreateClient(&token)

	orgs, _ := enclaveClient.GetOrgs()

	organisationClient := enclaveClient.CreateOrganisationClient(orgs[0])

	email := "tom.soulard+1337@enclave.io"
	err := organisationClient.InviteUser(&email)

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

	err = organisationClient.CancelUser(&email)
	if err != nil {
		t.Error(err)
	}
}
