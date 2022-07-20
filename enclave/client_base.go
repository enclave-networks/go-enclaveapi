package enclave

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/enclave-networks/go-enclaveapi/data"
)

// Shared functionality that is needed by the other clients
type ClientBase struct {
	baseURL    *url.URL
	token      *string
	httpClient *http.Client
	currentOrg *data.AccountOrganisation
}

// Encode data of type any into a JSON byte buffer
func Encode(data any) (*bytes.Buffer, error) {
	postBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	requestBody := bytes.NewBuffer(postBody)

	return requestBody, nil
}

// Decode an HTTP Response to Type of T
func Decode[T any](response *http.Response) *T {
	var toDecode = new(T)
	err := json.NewDecoder(response.Body).Decode(toDecode)
	if err != nil {
		return nil
	}

	return toDecode
}

// Helper to create a HTTP request with the correct prefix using the client
func (client *ClientBase) createRequest(route string, method string, body io.Reader) (*http.Request, error) {
	orgRoute := fmt.Sprintf("org/%s%s", client.currentOrg.OrgId, route)

	reqUrl := getRequestUrl(*client.baseURL, orgRoute)
	req, err := http.NewRequest(method, reqUrl.String(), body)

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	if err != nil {
		return nil, err
	}

	setRequestHeader(client.token, req)

	return req, nil
}

// Get a complete URL from a relative string
func getRequestUrl(baseUrl url.URL, relative string) *url.URL {
	rel := &url.URL{Path: relative}
	reqUrl := baseUrl.ResolveReference(rel)

	return reqUrl
}

// Set request headers this
func setRequestHeader(token *string, request *http.Request) {
	request.Header.Set("Authorization", "Bearer"+*token)
	request.Header.Set("User-Agent", "go-enclaveapi")
}

// Check for a success status code (this logic is ported from the C# HttpClient)
func isSuccessStatusCode(response *http.Response) error {
	isSuccess := response.StatusCode >= 200 && response.StatusCode <= 299
	if !isSuccess {
		data := Decode[data.HttpErrorResponse](response)

		// if we have body data return error containing that information
		if data != nil {
			return fmt.Errorf("status code does not indicate a successful response %v \n error details \n title: %s \n detail: %s", response.StatusCode, data.Title, data.Detail)
		}
		return fmt.Errorf("status code does not indicate a successful response %v", response.StatusCode)
	}

	return nil
}
