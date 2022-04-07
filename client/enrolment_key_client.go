package client

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/enclave-networks/go-enclaveapi/data"
)

type EnrolmentKeyClient struct {
	base *ClientBase
}

func (client *EnrolmentKeyClient) GetEnrolmentKeys(searchTerm *string, includeDisabled *bool, sortOrder *int, pageNumber *int, perPage *int) ([]*data.EnrolmentKeySummary, error) {
	req, err := client.base.createRequest("/enrolment-keys", http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	buildEnrolmentKeyQuery(req, searchTerm, includeDisabled, sortOrder, pageNumber, perPage)

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return nil, err
	}

	enrolmentKeys := Decode[data.PaginatedResponse[data.EnrolmentKeySummary]](response)

	return enrolmentKeys.Items, nil
}

func (client *EnrolmentKeyClient) Create(create *data.EnrolmentKeyCreate) (*data.EnrolmentKey, error) {
	if create == nil {
		err := fmt.Errorf("create model is nil")
		return nil, err
	}

	body, err := Encode(create)
	if err != nil {
		return nil, err
	}

	req, err := client.base.createRequest("/enrolment-keys", http.MethodPost, body)
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

	enrolmentKey := Decode[data.EnrolmentKey](response)

	return enrolmentKey, nil
}

func (client *EnrolmentKeyClient) Get(enrolmentKeyId *string) (*data.EnrolmentKey, error) {
	route := fmt.Sprintf("/enrolment-keys/%s", *enrolmentKeyId)
	req, err := client.base.createRequest(route, http.MethodGet, nil)
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

	enrolmentKey := Decode[data.EnrolmentKey](response)

	return enrolmentKey, nil
}

func (client *EnrolmentKeyClient) Enable(enrolmentKeyId *string) (*data.EnrolmentKey, error) {
	route := fmt.Sprintf("/enrolment-keys/%s/enable", *enrolmentKeyId)
	req, err := client.base.createRequest(route, http.MethodPut, nil)
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

	enrolmentKey := Decode[data.EnrolmentKey](response)

	return enrolmentKey, nil
}

func (client *EnrolmentKeyClient) Disable(enrolmentKeyId *string) (*data.EnrolmentKey, error) {
	route := fmt.Sprintf("/enrolment-keys/%s/disable", *enrolmentKeyId)
	req, err := client.base.createRequest(route, http.MethodPut, nil)
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

	enrolmentKey := Decode[data.EnrolmentKey](response)

	return enrolmentKey, nil
}

func (client *EnrolmentKeyClient) BulkEnable(enrolmentKeyIds ...*string) (*data.EnrolmentKey, error) {
	if enrolmentKeyIds == nil {
		err := fmt.Errorf("no enrolmentKey Ids")
		return nil, err
	}

	body, err := Encode(enrolmentKeyIds)
	if err != nil {
		return nil, err
	}

	req, err := client.base.createRequest("/enrolment-keys/enable", http.MethodPut, body)
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

	enrolmentKey := Decode[data.EnrolmentKey](response)

	return enrolmentKey, nil
}

func (client *EnrolmentKeyClient) BulkDisable(enrolmentKeyIds ...*string) (*data.EnrolmentKey, error) {
	if enrolmentKeyIds == nil {
		err := fmt.Errorf("no enrolmentKey Ids")
		return nil, err
	}

	body, err := Encode(enrolmentKeyIds)
	if err != nil {
		return nil, err
	}

	req, err := client.base.createRequest("/enrolment-keys/disable", http.MethodPut, body)
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

	enrolmentKey := Decode[data.EnrolmentKey](response)

	return enrolmentKey, nil
}

func buildEnrolmentKeyQuery(req *http.Request, searchTerm *string, includeDisabled *bool, sortOrder *int, pageNumber *int, perPage *int) {
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
