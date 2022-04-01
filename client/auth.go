package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/enclave-networks/go-enclaveapi/data"
)

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

func (client *EnclaveClient) GetOrgs() (*[]data.AccountOrganisation, error) {
	reqUrl := client.getRequestUrl("/account/orgs")
	req, err := http.NewRequest("GET", reqUrl.String(), nil)

	fmt.Println(reqUrl.String())

	if err != nil {
		return nil, err
	}

	client.setRequestHeader(req)

	response, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var orgTopLevel data.AccountOrganisationTopLevel
	json.NewDecoder(response.Body).Decode(&orgTopLevel)

	return orgTopLevel.Orgs, nil
}

func (client *EnclaveClient) getRequestUrl(relative string) *url.URL {
	rel := &url.URL{Path: relative}
	reqUrl := client.baseURL.ResolveReference(rel)

	return reqUrl
}

func (client *EnclaveClient) setRequestHeader(request *http.Request) {
	request.Header.Set("Authorization", "Bearer"+*client.token)
	request.Header.Set("User-Agent", "go-enclaveapi")
}

type EnclaveClient struct {
	baseURL    *url.URL
	token      *string
	httpClient *http.Client
}
