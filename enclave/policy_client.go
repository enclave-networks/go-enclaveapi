package enclave

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/enclave-networks/go-enclaveapi/data"
	"github.com/enclave-networks/go-enclaveapi/data/policy"
)

type PolicyClient struct {
	base *ClientBase
}

func (client *PolicyClient) GetPolicies(searchTerm *string, includeDisabled *bool, sortOrder *policy.PolicySortOrder, pageNumber *int, perPage *int) ([]policy.Policy, error) {
	req, err := client.base.createRequest("/policies", http.MethodGet, nil)
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

func (client *PolicyClient) Create(create policy.PolicyCreate) (policy.Policy, error) {
	body, err := Encode(create)
	if err != nil {
		return policy.Policy{}, err
	}

	req, err := client.base.createRequest("/policies", http.MethodPost, body)
	if err != nil {
		return policy.Policy{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return policy.Policy{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return policy.Policy{}, err
	}

	policy := Decode[policy.Policy](response)

	return *policy, nil
}

func (client *PolicyClient) Get(policyId int) (policy.Policy, error) {
	route := fmt.Sprintf("/policies/%v", policyId)
	req, err := client.base.createRequest(route, http.MethodGet, nil)
	if err != nil {
		return policy.Policy{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return policy.Policy{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return policy.Policy{}, err
	}

	policy := Decode[policy.Policy](response)

	return *policy, nil
}

func (client *PolicyClient) Update(policyId int, patch policy.PolicyPatch) (policy.Policy, error) {
	body, err := Encode(patch)
	if err != nil {
		return policy.Policy{}, err
	}

	route := fmt.Sprintf("/policies/%v", policyId)
	req, err := client.base.createRequest(route, http.MethodPatch, body)
	if err != nil {
		return policy.Policy{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return policy.Policy{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return policy.Policy{}, err
	}

	policy := Decode[policy.Policy](response)

	return *policy, nil
}

func (client *PolicyClient) Delete(policyId int) (policy.Policy, error) {
	route := fmt.Sprintf("/policies/%v/delete", policyId)
	req, err := client.base.createRequest(route, http.MethodDelete, nil)
	if err != nil {
		return policy.Policy{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return policy.Policy{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return policy.Policy{}, err
	}

	policy := Decode[policy.Policy](response)

	return *policy, nil
}

func (client *PolicyClient) Enable(policyId int) (policy.Policy, error) {
	route := fmt.Sprintf("/policies/%v/enable", policyId)
	req, err := client.base.createRequest(route, http.MethodPut, nil)
	if err != nil {
		return policy.Policy{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return policy.Policy{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return policy.Policy{}, err
	}

	policy := Decode[policy.Policy](response)

	return *policy, nil
}

func (client *PolicyClient) Disable(policyId int) (policy.Policy, error) {
	route := fmt.Sprintf("/policies/%v/disable", policyId)
	req, err := client.base.createRequest(route, http.MethodPut, nil)
	if err != nil {
		return policy.Policy{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return policy.Policy{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return policy.Policy{}, err
	}

	policy := Decode[policy.Policy](response)

	return *policy, nil
}

func (client *PolicyClient) BulkDelete(policyIds ...int) (int, error) {
	if policyIds == nil {
		err := fmt.Errorf("no policy Ids")
		return 0, err
	}

	body, err := Encode(policyIds)
	if err != nil {
		return 0, err
	}

	req, err := client.base.createRequest("/policies/delete", http.MethodDelete, body)
	if err != nil {
		return 0, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return 0, err
	}

	bulkResult := Decode[policy.PolicyBulkDeleteResult](response)

	return bulkResult.PoliciesDeleted, nil
}

func (client *PolicyClient) BulkEnable(policyIds ...int) (int, error) {
	if policyIds == nil {
		err := fmt.Errorf("no policy Ids")
		return 0, err
	}

	body, err := Encode(policyIds)
	if err != nil {
		return 0, err
	}

	req, err := client.base.createRequest("/policies/enable", http.MethodPut, body)
	if err != nil {
		return 0, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return 0, err
	}

	bulkResult := Decode[policy.PolicyBulkUpdateResult](response)

	return bulkResult.PoliciesUpdated, nil
}

func (client *PolicyClient) BulkDisable(policyIds ...int) (int, error) {
	if policyIds == nil {
		err := fmt.Errorf("no policy Ids")
		return 0, err
	}

	body, err := Encode(policyIds)
	if err != nil {
		return 0, err
	}

	req, err := client.base.createRequest("/policies/disable", http.MethodPut, body)
	if err != nil {
		return 0, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return 0, err
	}

	bulkResult := Decode[policy.PolicyBulkUpdateResult](response)

	return bulkResult.PoliciesUpdated, nil
}

func buildPolicyQuery(req *http.Request, searchTerm *string, includeDisabled *bool, sortOrder *policy.PolicySortOrder, pageNumber *int, perPage *int) {
	query := req.URL.Query()

	if searchTerm != nil {
		query.Add("search", *searchTerm)
	}

	if includeDisabled != nil {
		query.Add("include_disabled", strconv.FormatBool(*includeDisabled))
	}

	if sortOrder != nil {
		query.Add("sort", strconv.FormatInt(int64(*sortOrder), 10))
	}

	if pageNumber != nil {
		query.Add("page", strconv.FormatInt(int64(*pageNumber), 10))
	}

	if perPage != nil {
		query.Add("per_page", strconv.FormatInt(int64(*perPage), 10))
	}
}
