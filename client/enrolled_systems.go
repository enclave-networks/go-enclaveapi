package client

import (
	"net/http"
	"strconv"

	"github.com/enclave-networks/go-enclaveapi/data"
)

type EnrolledSystemsClient struct {
	base *ClientBase
}

func (client *EnrolledSystemsClient) GetSystems(
	enrolmentKeyId *int,
	searchTerm *string,
	includeDisabled *bool,
	sortOrder *int,
	dnsName *string,
	pageNumber *int,
	perPage *int) (*data.PaginatedResponse[data.EnrolledSystemSummary], error) {
	req, err := client.base.createRequest("/systems", http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	buildSystemsQuery(req, enrolmentKeyId, searchTerm, includeDisabled, sortOrder, dnsName, pageNumber, perPage)

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return nil, err
	}

	systems := Decode[data.PaginatedResponse[data.EnrolledSystemSummary]](response)

	return systems, nil
}

func (client *EnrolledSystemsClient) RevokeSystems(systemIds ...*string) (*int, error) {
	requestBody, err := Encode(systemIds)
	if err != nil {
		return nil, err
	}

	req, err := client.base.createRequest("/systems", http.MethodDelete, requestBody)
	if err != nil {
		return nil, err
	}
	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

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
