package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/enclave-networks/go-enclaveapi/data"
)

func (client *OrganisationClient) GetSystems(
	enrolmentKeyId *int,
	searchTerm *string,
	includeDisabled *bool,
	sortOrder *int,
	dnsName *string,
	pageNumber *int,
	perPage *int) (*data.PaginatedResponse[data.EnrolledSystemSummary], error) {
	req, err := client.createOrgRequest("/systems", "GET", nil)
	if err != nil {
		return nil, err
	}

	buildSystemsQuery(req, enrolmentKeyId, searchTerm, includeDisabled, sortOrder, dnsName, pageNumber, perPage)

	response, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	systems := Decode[data.PaginatedResponse[data.EnrolledSystemSummary]](response)

	return systems, nil
}

func (client *OrganisationClient) RevokeSystems(systemIds ...*string) (*int, error) {
	postBody, err := json.Marshal(systemIds)
	if err != nil {
		return nil, err
	}

	requestBody := bytes.NewBuffer(postBody)

	req, err := client.createOrgRequest("/systems", "DELETE", requestBody)
	if err != nil {
		return nil, err
	}
	response, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return nil, err
	}

	count := Decode[int](response)

	return count, nil
}

func buildSystemsQuery(
	req *http.Request,
	enrolmentKeyId *int,
	searchTerm *string,
	includeDisabled *bool,
	sortOrder *int,
	dnsName *string,
	pageNumber *int,
	perPage *int) {
	query := req.URL.Query()

	if enrolmentKeyId != nil {
		query.Add("enrolment_key", strconv.FormatInt(int64(*enrolmentKeyId), 10))
	}

	if searchTerm != nil {
		query.Add("search", *searchTerm)
	}

	if includeDisabled != nil {
		query.Add("include_disabled", strconv.FormatBool(*includeDisabled))
	}

	if sortOrder != nil {
		query.Add("sort", strconv.FormatInt(int64(*sortOrder), 10))
	}

	if dnsName != nil {
		query.Add("dns", *dnsName)
	}

	if pageNumber != nil {
		query.Add("page", strconv.FormatInt(int64(*pageNumber), 10))
	}

	if perPage != nil {
		query.Add("per_page", strconv.FormatInt(int64(*perPage), 10))
	}
}
