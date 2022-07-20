package enclave

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/enclave-networks/go-enclaveapi/data"
	"github.com/enclave-networks/go-enclaveapi/data/tag"
)

// Provides operations to get, create, and manipulate Tags.
type TagsClient struct {
	base *ClientBase
}

// Gets a paginated list of Tags which can be searched and iterated upon.
func (client *TagsClient) GetTags(
	searchTerm *string,
	sortOrder *tag.TagQuerySortOrder,
	pageNumber *int,
	perPage *int) (*data.PaginatedResponse[tag.BasicTag], error) {
	req, err := client.base.createRequest("/tags", http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	buildTagQuery(req, searchTerm, sortOrder, pageNumber, perPage)

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response)
	if err != nil {
		return nil, err
	}

	tags := Decode[data.PaginatedResponse[tag.BasicTag]](response)

	return tags, nil
}

// Creates a Tag using a "TagCreate" Struct.
func (client *TagsClient) Create(create tag.TagCreate) (tag.DetailedTag, error) {
	body, err := Encode(create)
	if err != nil {
		return tag.DetailedTag{}, err
	}

	req, err := client.base.createRequest("/tags", http.MethodPost, body)
	if err != nil {
		return tag.DetailedTag{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return tag.DetailedTag{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response)
	if err != nil {
		return tag.DetailedTag{}, err
	}

	tag := Decode[tag.DetailedTag](response)

	return *tag, nil
}

// Delete multiple Tags.
func (client *TagsClient) DeleteTags(tagNameOrRef ...string) (int, error) {
	if tagNameOrRef == nil {
		err := fmt.Errorf("no Tag Ids")
		return 0, err
	}

	body, err := Encode(tag.BulkTagAction{Tags: tagNameOrRef})
	if err != nil {
		return 0, err
	}

	req, err := client.base.createRequest("/tags", http.MethodDelete, body)
	if err != nil {
		return 0, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response)
	if err != nil {
		return 0, err
	}

	bulkResult := Decode[tag.BulkTagDeleteResult](response)

	return bulkResult.TagsDeleted, nil
}

// Get a specific Tag.
func (client *TagsClient) Get(tagNameOrRef string) (tag.DetailedTag, error) {
	route := fmt.Sprintf("/tags/%v", tagNameOrRef)
	req, err := client.base.createRequest(route, http.MethodGet, nil)
	if err != nil {
		return tag.DetailedTag{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return tag.DetailedTag{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response)
	if err != nil {
		return tag.DetailedTag{}, err
	}

	tag := Decode[tag.DetailedTag](response)

	return *tag, nil
}

// Starts an update patch request.
func (client *TagsClient) Update(tagNameOrRef string, patch tag.TagPatch) (tag.DetailedTag, error) {
	body, err := Encode(patch)
	if err != nil {
		return tag.DetailedTag{}, err
	}

	route := fmt.Sprintf("/tags/%v", tagNameOrRef)
	req, err := client.base.createRequest(route, http.MethodPatch, body)
	if err != nil {
		return tag.DetailedTag{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return tag.DetailedTag{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response)
	if err != nil {
		return tag.DetailedTag{}, err
	}

	tag := Decode[tag.DetailedTag](response)

	return *tag, nil
}

// Delete a Tag.
func (client *TagsClient) Delete(tagNameOrRef string) (tag.DetailedTag, error) {
	route := fmt.Sprintf("/tags/%v", tagNameOrRef)
	req, err := client.base.createRequest(route, http.MethodDelete, nil)
	if err != nil {
		return tag.DetailedTag{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return tag.DetailedTag{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response)
	if err != nil {
		return tag.DetailedTag{}, err
	}

	tag := Decode[tag.DetailedTag](response)

	return *tag, nil
}

func buildTagQuery(req *http.Request, searchTerm *string, sortOrder *tag.TagQuerySortOrder, pageNumber *int, perPage *int) {
	query := req.URL.Query()

	if searchTerm != nil {
		query.Add("search", *searchTerm)
	}

	if sortOrder != nil {
		query.Add("sort", string(*sortOrder))
	}

	if pageNumber != nil {
		query.Add("page", strconv.FormatInt(int64(*pageNumber), 10))
	}

	if perPage != nil {
		query.Add("per_page", strconv.FormatInt(int64(*perPage), 10))
	}
}
