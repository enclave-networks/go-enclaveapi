package enclaveclient_test

import (
	"testing"

	"github.com/enclave-networks/go-enclaveapi/client"
)

func TestThing(t *testing.T) {
	token := "p9rcFksNsHALkfyqyfgRzYq4AXwcuxr22CN9Mc5PG42umHPUiPhnzX7kiRfdWM3"
	enclaveClient := client.CreateClient(&token)

	_ = enclaveClient.GetOrgs()
}
