package enclave

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/enclave-networks/go-enclaveapi/data"
	"github.com/enclave-networks/go-enclaveapi/data/trustrequirement"
)

// Provides operations to get, create, and manipulate Trust Requirements.
type TrustRequirementsClient struct {
	base *ClientBase
}

// Gets a paginated list of Trust Requirements which can be searched and iterated upon.
func (client *TrustRequirementsClient) GetTrustRequirements(
	searchTerm *string,
	sortOrder *trustrequirement.TrustRequirementSortOrder,
	pageNumber *int,
	perPage *int) (*data.PaginatedResponse[trustrequirement.TrustRequirementSummary], error) {
	req, err := client.base.createRequest("/trust-requirements", http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	buildTrustRequirementQuery(req, searchTerm, sortOrder, pageNumber, perPage)

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return nil, err
	}

	trustRequirements := Decode[data.PaginatedResponse[trustrequirement.TrustRequirementSummary]](response)

	return trustRequirements, nil
}

// Creates a Trust Requirement using a "TrustRequirementCreate" Struct.
func (client *TrustRequirementsClient) Create(create trustrequirement.TrustRequirementCreate) (trustrequirement.TrustRequirement, error) {
	body, err := Encode(create)
	if err != nil {
		return trustrequirement.TrustRequirement{}, err
	}

	req, err := client.base.createRequest("/trust-requirements", http.MethodPost, body)
	if err != nil {
		return trustrequirement.TrustRequirement{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return trustrequirement.TrustRequirement{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return trustrequirement.TrustRequirement{}, err
	}

	trustRequirement := Decode[trustrequirement.TrustRequirement](response)

	return *trustRequirement, nil
}

// Delete multiple Trust Requirements.
func (client *TrustRequirementsClient) DeleteTrustRequirements(trustRequirementIds ...trustrequirement.TrustRequirementId) (int, error) {
	if trustRequirementIds == nil {
		err := fmt.Errorf("no trust requirement Ids")
		return 0, err
	}

	body, err := Encode(trustrequirement.BulkTrustRequirementDeleteAction{RequirementIds: trustRequirementIds})
	if err != nil {
		return 0, err
	}

	req, err := client.base.createRequest("/trust-requirements", http.MethodDelete, body)
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

	bulkResult := Decode[trustrequirement.BulkTrustRequirementDeleteResult](response)

	return bulkResult.RequirementsDeleted, nil
}

// Get a specific Trust Requirement.
func (client *TrustRequirementsClient) Get(trustRequirementId trustrequirement.TrustRequirementId) (trustrequirement.TrustRequirement, error) {
	route := fmt.Sprintf("/trust-requirements/%v", trustRequirementId)
	req, err := client.base.createRequest(route, http.MethodGet, nil)
	if err != nil {
		return trustrequirement.TrustRequirement{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return trustrequirement.TrustRequirement{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return trustrequirement.TrustRequirement{}, err
	}

	trustRequirement := Decode[trustrequirement.TrustRequirement](response)

	return *trustRequirement, nil
}

// Starts an update patch request.
func (client *TrustRequirementsClient) Update(requirementId trustrequirement.TrustRequirementId, patch trustrequirement.TrustRequirementPatch) (trustrequirement.TrustRequirement, error) {
	body, err := Encode(patch)
	if err != nil {
		return trustrequirement.TrustRequirement{}, err
	}

	route := fmt.Sprintf("/trust-requirements/%v", requirementId)
	req, err := client.base.createRequest(route, http.MethodPatch, body)
	if err != nil {
		return trustrequirement.TrustRequirement{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return trustrequirement.TrustRequirement{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return trustrequirement.TrustRequirement{}, err
	}

	trustRequirement := Decode[trustrequirement.TrustRequirement](response)

	return *trustRequirement, nil
}

// Delete a Trust Requirement.
func (client *TrustRequirementsClient) Delete(requirementId trustrequirement.TrustRequirementId) (trustrequirement.TrustRequirement, error) {
	route := fmt.Sprintf("/trust-requirements/%v", requirementId)
	req, err := client.base.createRequest(route, http.MethodDelete, nil)
	if err != nil {
		return trustrequirement.TrustRequirement{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return trustrequirement.TrustRequirement{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return trustrequirement.TrustRequirement{}, err
	}

	trustRequirement := Decode[trustrequirement.TrustRequirement](response)

	return *trustRequirement, nil
}

func buildTrustRequirementQuery(req *http.Request, searchTerm *string, sortOrder *trustrequirement.TrustRequirementSortOrder, pageNumber *int, perPage *int) {
	query := req.URL.Query()

	if searchTerm != nil {
		query.Add("search", *searchTerm)
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
