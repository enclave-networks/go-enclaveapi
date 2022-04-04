package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/enclave-networks/go-enclaveapi/data"
)

type EnclaveClient struct {
	baseURL    *url.URL
	token      *string
	httpClient *http.Client
}

func CreateClient(token *string) *EnclaveClient {
	httpClient := &http.Client{Timeout: time.Minute}

	baseUrl := &url.URL{
		Scheme: data.Scheme,
		Host:   data.BaseUrl,
	}

	return &EnclaveClient{
		httpClient: httpClient,
		token:      token,
		baseURL:    baseUrl,
	}
}

func (client *EnclaveClient) GetOrgs() ([]*data.AccountOrganisation, error) {
	req, err := client.createEnclaveRequest("/account/orgs", "GET", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var orgTopLevel data.AccountOrganisationTopLevel
	json.NewDecoder(response.Body).Decode(&orgTopLevel)

	return orgTopLevel.Orgs, nil
}

func (client *EnclaveClient) CreateOrganisationClient(org *data.AccountOrganisation) *OrganisationClient {
	return &OrganisationClient{
		baseURL:    client.baseURL,
		token:      client.token,
		httpClient: client.httpClient,
		currentOrg: org,
	}
}

func (client *EnclaveClient) createEnclaveRequest(route string, method string, body io.Reader) (*http.Request, error) {
	reqUrl := getRequestUrl(client.baseURL, route)
	req, err := http.NewRequest(method, reqUrl.String(), body)

	fmt.Println(reqUrl.String())

	if err != nil {
		return nil, err
	}

	setRequestHeader(client.token, req)

	return req, nil
}
