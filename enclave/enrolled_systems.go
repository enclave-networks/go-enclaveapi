package enclave

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/enclave-networks/go-enclaveapi/data"
	"github.com/enclave-networks/go-enclaveapi/data/enrolledsystem"
	"github.com/enclave-networks/go-enclaveapi/data/enrolmentkey"
)

// Provides operations to get, create, and manipulate Enrolled Systems.
type EnrolledSystemsClient struct {
	base *ClientBase
}

// Gets a paginated list of Systems which can be searched and iterated upon.
func (client *EnrolledSystemsClient) GetSystems(
	enrolmentKeyId *enrolmentkey.EnrolmentKeyId,
	searchTerm *string,
	includeDisabled *bool,
	sortOrder *int,
	dnsName *string,
	pageNumber *int,
	perPage *int) (*data.PaginatedResponse[enrolledsystem.EnrolledSystemSummary], error) {
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

	err = isSuccessStatusCode(response)
	if err != nil {
		return nil, err
	}

	systems := Decode[data.PaginatedResponse[enrolledsystem.EnrolledSystemSummary]](response)

	return systems, nil
}

// Permanetly revoke multiple systems.
func (client *EnrolledSystemsClient) RevokeSystems(systemIds ...enrolledsystem.SystemId) (int, error) {
	if systemIds == nil {
		err := fmt.Errorf("no system Ids")
		return 0, err
	}

	requestBody, err := Encode(enrolledsystem.EnrolledSystemBulkAction{SystemIds: systemIds})
	if err != nil {
		return -1, err
	}

	req, err := client.base.createRequest("/systems", http.MethodDelete, requestBody)
	if err != nil {
		return -1, err
	}
	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return -1, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response)
	if err != nil {
		return -1, err
	}

	result := Decode[enrolledsystem.EnrolledSystemBulkRevokedResult](response)

	return result.SystemsRevoked, nil
}

func buildSystemsQuery(
	req *http.Request,
	enrolmentKeyId *enrolmentkey.EnrolmentKeyId,
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
