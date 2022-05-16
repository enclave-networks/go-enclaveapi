package enclave

import (
	"net/http"

	"github.com/enclave-networks/go-enclaveapi/data"
	"github.com/enclave-networks/go-enclaveapi/data/policy"
)

type PolicyClient struct {
	base *ClientBase
}

func (client *PolicyClient) GetPolicies(searchTerm *string, includeDisabled *bool, sortOrder *policy.PolicySortOrder, pageNumber *int, perPage *int) ([]policy.Policy, error) {
	req, err := client.base.createRequest("/policy", http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	buildPolicyQuery(req, searchTerm, includeDisabled, sortOrder, pageNumber, perPage)

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return nil, err
	}

	policies := Decode[data.PaginatedResponse[policy.Policy]](response)

	return policies.Items, nil
}

func buildPolicyQuery(req *http.Request, searchTerm *string, includeDisabled *bool, sortOrder *policy.PolicySortOrder, pageNumber *int, perPage *int) {
	panic("unimplemented")
}
