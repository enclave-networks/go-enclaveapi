package enclave

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/enclave-networks/go-enclaveapi/data"
)

type Client struct {
	baseURL    *url.URL
	token      *string
	httpClient *http.Client
}

func CreateClient(token string) *Client {
	httpClient := &http.Client{Timeout: time.Minute}

	baseUrl := &url.URL{
		Scheme: data.Scheme,
		Host:   data.BaseUrl,
	}

	return &Client{
		httpClient: httpClient,
		token:      &token,
		baseURL:    baseUrl,
	}
}

func CreateClientWithUrl(token string, baseUrl string) (*Client, error) {
	httpClient := &http.Client{Timeout: time.Minute}

	parsedUrl, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	return &Client{
		httpClient: httpClient,
		token:      &token,
		baseURL:    parsedUrl,
	}, nil
}

func (client *Client) GetOrgs() (*[]data.AccountOrganisation, error) {
	req, err := client.createEnclaveRequest("/account/orgs", http.MethodGet, nil)
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

	return &orgTopLevel.Orgs, nil
}

func (client *Client) CreateOrganisationClient(org data.AccountOrganisation) *OrganisationClient {
	base := &ClientBase{
		baseURL:    client.baseURL,
		token:      client.token,
		httpClient: client.httpClient,
		currentOrg: &org,
	}
	return &OrganisationClient{
		base: base,
		Systems: &EnrolledSystemsClient{
			base: base,
		},
		EnrolmentKey: &EnrolmentKeyClient{
			base: base,
		},
	}
}

func (client *Client) createEnclaveRequest(route string, method string, body io.Reader) (*http.Request, error) {
	reqUrl := getRequestUrl(*client.baseURL, route)
	req, err := http.NewRequest(method, reqUrl.String(), body)

	fmt.Println(reqUrl.String())

	if err != nil {
		return nil, err
	}

	setRequestHeader(client.token, req)

	return req, nil
}
