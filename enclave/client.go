package enclave

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/enclave-networks/go-enclaveapi/data"
)

// The standard client for using the enclave API from here you can create associated clients
type Client struct {
	baseURL    *url.URL
	token      *string
	httpClient *http.Client
}

// Create a new client with a provided token
func New(token string) *Client {
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

// Create the client with the specified base url (this is primarly for testing)
func NewWithUrl(token string, baseUrl string) (*Client, error) {
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

// Get all orgs associated to the current token
func (client *Client) GetOrgs() ([]data.AccountOrganisation, error) {
	req, err := client.createEnclaveRequest("/account/orgs", http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	orgTopLevel := Decode[data.AccountOrganisationTopLevel](response)

	return orgTopLevel.Orgs, nil
}

// Create the base organisation client from here you can do all the expected requests as described in our docs https://api.enclave.io
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
		EnrolmentKeys: &EnrolmentKeyClient{
			base: base,
		},
		Dns: &DnsClient{
			base: base,
		},
		Policies: &PolicyClient{
			base: base,
		},
		TrustRequirements: &TrustRequirementsClient{
			base: base,
		},
	}
}

// create an enclave request that doesn't include the org prefix
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
