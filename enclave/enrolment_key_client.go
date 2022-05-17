package enclave

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/enclave-networks/go-enclaveapi/data"
	"github.com/enclave-networks/go-enclaveapi/data/enrolmentkey"
)

type EnrolmentKeyClient struct {
	base *ClientBase
}

func (client *EnrolmentKeyClient) GetEnrolmentKeys(searchTerm *string, includeDisabled *bool, sortOrder *enrolmentkey.EnrolmentKeySortOrder, pageNumber *int, perPage *int) ([]enrolmentkey.EnrolmentKeySummary, error) {
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

	enrolmentKeys := Decode[data.PaginatedResponse[enrolmentkey.EnrolmentKeySummary]](response)

	return enrolmentKeys.Items, nil
}

func (client *EnrolmentKeyClient) Create(create enrolmentkey.EnrolmentKeyCreate) (enrolmentkey.EnrolmentKey, error) {
	body, err := Encode(create)
	if err != nil {
		return enrolmentkey.EnrolmentKey{}, err
	}

	req, err := client.base.createRequest("/enrolment-keys", http.MethodPost, body)
	if err != nil {
		return enrolmentkey.EnrolmentKey{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return enrolmentkey.EnrolmentKey{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return enrolmentkey.EnrolmentKey{}, err
	}

	enrolmentKey := Decode[enrolmentkey.EnrolmentKey](response)

	return *enrolmentKey, nil
}

func (client *EnrolmentKeyClient) Get(enrolmentKeyId int) (enrolmentkey.EnrolmentKey, error) {
	route := fmt.Sprintf("/enrolment-keys/%v", enrolmentKeyId)
	req, err := client.base.createRequest(route, http.MethodGet, nil)
	if err != nil {
		return enrolmentkey.EnrolmentKey{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return enrolmentkey.EnrolmentKey{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return enrolmentkey.EnrolmentKey{}, err
	}

	enrolmentKey := Decode[enrolmentkey.EnrolmentKey](response)

	return *enrolmentKey, nil
}

func (client *EnrolmentKeyClient) Update(enrolmentKeyId int, patch enrolmentkey.EnrolmentKeyPatch) (enrolmentkey.EnrolmentKey, error) {
	body, err := Encode(patch)
	if err != nil {
		return enrolmentkey.EnrolmentKey{}, err
	}

	route := fmt.Sprintf("/enrolment-keys/%v", enrolmentKeyId)
	req, err := client.base.createRequest(route, http.MethodPatch, body)
	if err != nil {
		return enrolmentkey.EnrolmentKey{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return enrolmentkey.EnrolmentKey{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return enrolmentkey.EnrolmentKey{}, err
	}

	enrolmentKey := Decode[enrolmentkey.EnrolmentKey](response)

	return *enrolmentKey, nil
}

func (client *EnrolmentKeyClient) Enable(enrolmentKeyId int) (enrolmentkey.EnrolmentKey, error) {
	route := fmt.Sprintf("/enrolment-keys/%v/enable", enrolmentKeyId)
	req, err := client.base.createRequest(route, http.MethodPut, nil)
	if err != nil {
		return enrolmentkey.EnrolmentKey{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return enrolmentkey.EnrolmentKey{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return enrolmentkey.EnrolmentKey{}, err
	}

	enrolmentKey := Decode[enrolmentkey.EnrolmentKey](response)

	return *enrolmentKey, nil
}

func (client *EnrolmentKeyClient) Disable(enrolmentKeyId int) (enrolmentkey.EnrolmentKey, error) {
	route := fmt.Sprintf("/enrolment-keys/%v/disable", enrolmentKeyId)
	req, err := client.base.createRequest(route, http.MethodPut, nil)
	if err != nil {
		return enrolmentkey.EnrolmentKey{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return enrolmentkey.EnrolmentKey{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return enrolmentkey.EnrolmentKey{}, err
	}

	enrolmentKey := Decode[enrolmentkey.EnrolmentKey](response)

	return *enrolmentKey, nil
}

func (client *EnrolmentKeyClient) BulkEnable(enrolmentKeyIds ...int) (int, error) {
	if enrolmentKeyIds == nil {
		err := fmt.Errorf("no enrolmentKey Ids")
		return 0, err
	}

	body, err := Encode(enrolmentKeyIds)
	if err != nil {
		return 0, err
	}

	req, err := client.base.createRequest("/enrolment-keys/enable", http.MethodPut, body)
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

	bulkResult := Decode[enrolmentkey.EnrolmentKeyBulkActionResult](response)

	return bulkResult.KeysModified, nil
}

func (client *EnrolmentKeyClient) BulkDisable(enrolmentKeyIds ...int) (int, error) {
	if enrolmentKeyIds == nil {
		err := fmt.Errorf("no enrolmentKey Ids")
		return 0, err
	}

	body, err := Encode(enrolmentKeyIds)
	if err != nil {
		return 0, err
	}

	req, err := client.base.createRequest("/enrolment-keys/disable", http.MethodPut, body)
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

	bulkResult := Decode[enrolmentkey.EnrolmentKeyBulkActionResult](response)

	return bulkResult.KeysModified, nil
}

func buildEnrolmentKeyQuery(req *http.Request, searchTerm *string, includeDisabled *bool, sortOrder *enrolmentkey.EnrolmentKeySortOrder, pageNumber *int, perPage *int) {
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
