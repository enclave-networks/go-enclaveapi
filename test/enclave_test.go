package enclaveclient_test

import (
	"testing"

	"github.com/enclave-networks/go-enclaveapi/client"
)

func Test_when_making_a_call_to_GetSystems_should_return_items(t *testing.T) {
	token := "p9rcFksNsHALkfyqyfgRzYq4AXwcuxr22CN9Mc5PG42umHPUiPhnzX7kiRfdWM3"
	enclaveClient := client.CreateClient(&token)

	orgs, _ := enclaveClient.GetOrgs()

	enclaveClient.SetCurrentOrg(orgs[0])
	sys, err := enclaveClient.GetSystems(nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("error: %s", err)
	}

	if len(sys.Items) == 0 {
		t.Errorf("no items in return type")
	}
}
