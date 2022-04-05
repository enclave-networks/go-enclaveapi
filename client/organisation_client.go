package client

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/enclave-networks/go-enclaveapi/data"
)

type OrganisationClient struct {
	baseURL    *url.URL
	token      *string
	httpClient *http.Client
	currentOrg *data.AccountOrganisation
}

func (client *OrganisationClient) Get() (*data.Organisation, error) {
	req, err := client.createRequest("", "GET", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return nil, err
	}

	org := Decode[data.Organisation](response)

	return org, nil
}

func (client *OrganisationClient) Update(patch *data.OrganisationPatch) (*data.Organisation, error) {
	requestBody, err := Encode(patch)
	if err != nil {
		return nil, err
	}

	req, err := client.createRequest("", "POST", requestBody)
	if err != nil {
		return nil, err
	}

	response, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return nil, err
	}

	org := Decode[data.Organisation](response)

	return org, nil
}

func (client *OrganisationClient) GetOrganisationUsers() ([]*data.OrganisationUser, error) {
	req, err := client.createRequest("/users", "GET", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return nil, err
	}

	orgUsers := Decode[data.OrganisationUsersTopLevel](response)

	return orgUsers.Users, nil
}

func (client *OrganisationClient) RemoveUser(accountId *string) error {
	route := fmt.Sprintf("/users/%s", *accountId)
	req, err := client.createRequest(route, "DELETE", nil)
	if err != nil {
		return err
	}

	response, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return err
	}

	return nil
}

func (client *OrganisationClient) GetPendingInvites() ([]*data.OrganisationInvite, error) {
	req, err := client.createRequest("/invites", "GET", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return nil, err
	}

	orgInvites := Decode[data.OrganisationPendingInvites](response)

	return orgInvites.Invites, nil
}

func (client *OrganisationClient) InviteUser(emailAddress *string) error {
	requestBody, err := Encode(data.OrganisationInvite{
		EmailAddress: emailAddress,
	})
	if err != nil {
		return err
	}

	req, err := client.createRequest("/invites", "POST", requestBody)
	if err != nil {
		return err
	}

	response, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return err
	}

	return nil
}

func (client *OrganisationClient) CancelUser(emailAddress *string) error {
	requestBody, err := Encode(data.OrganisationInvite{
		EmailAddress: emailAddress,
	})
	if err != nil {
		return err
	}

	req, err := client.createRequest("/invites", "DELETE", requestBody)
	if err != nil {
		return err
	}

	response, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return err
	}

	return nil
}
