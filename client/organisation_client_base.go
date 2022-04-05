package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Encode(data any) (*bytes.Buffer, error) {
	postBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	requestBody := bytes.NewBuffer(postBody)

	return requestBody, nil
}

func Decode[T any](response *http.Response) *T {
	var toDecode = new(T)
	json.NewDecoder(response.Body).Decode(toDecode)
	return toDecode
}

func (client *OrganisationClient) createRequest(route string, method string, body io.Reader) (*http.Request, error) {
	orgRoute := fmt.Sprintf("org/%s%s", *client.currentOrg.OrgId, route)
	reqUrl := getRequestUrl(client.baseURL, orgRoute)
	req, err := http.NewRequest(method, reqUrl.String(), body)
	fmt.Println(reqUrl.String())

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	if err != nil {
		return nil, err
	}

	setRequestHeader(client.token, req)

	return req, nil
}

func getRequestUrl(baseUrl *url.URL, relative string) *url.URL {
	rel := &url.URL{Path: relative}
	reqUrl := baseUrl.ResolveReference(rel)

	return reqUrl
}

func setRequestHeader(token *string, request *http.Request) {
	request.Header.Set("Authorization", "Bearer"+*token)
	request.Header.Set("User-Agent", "go-enclaveapi")
}

func isSuccessStatusCode(statusCode int) error {
	isSuccess := statusCode >= 200 && statusCode <= 299
	if !isSuccess {
		return fmt.Errorf("status code does not indicate a successful response %v", statusCode)
	}

	return nil
}
