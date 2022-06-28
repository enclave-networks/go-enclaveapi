package enclave

import (
	"fmt"
	"net/http"

	"github.com/enclave-networks/go-enclaveapi/data/account"
	"github.com/enclave-networks/go-enclaveapi/data/organisation"
)

// Provides access to organisation level API calls and organisation related clients.
// For more information please refer to our API docs at https://api.enclave.io/.
type OrganisationClient struct {
	base                    *ClientBase
	Systems                 *EnrolledSystemsClient
	EnrolmentKey            *EnrolmentKeyClient
	Dns                     *DnsClient
	Policy                  *PolicyClient
	TrustRequirementsClient *TrustRequirementsClient
}

// Get more detail on your current organisaiton.
func (client *OrganisationClient) Get() (organisation.Organisation, error) {
	req, err := client.base.createRequest("", http.MethodGet, nil)
	if err != nil {
		return organisation.Organisation{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return organisation.Organisation{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return organisation.Organisation{}, err
	}

	org := Decode[organisation.Organisation](response)

	return *org, nil
}

// Performs an update patch request.
func (client *OrganisationClient) Update(patch organisation.OrganisationPatch) (organisation.Organisation, error) {
	requestBody, err := Encode(patch)
	if err != nil {
		return organisation.Organisation{}, err
	}

	req, err := client.base.createRequest("", http.MethodPost, requestBody)
	if err != nil {
		return organisation.Organisation{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return organisation.Organisation{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return organisation.Organisation{}, err
	}

	org := Decode[organisation.Organisation](response)

	return *org, nil
}

// Gets the users that have access to the current organisaiton.
func (client *OrganisationClient) GetOrganisationUsers() ([]organisation.OrganisationUser, error) {
	req, err := client.base.createRequest("/users", http.MethodGet, nil)
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

	orgUsers := Decode[organisation.OrganisationUsersTopLevel](response)

	return orgUsers.Users, nil
}

// Removes a user from the organisation.
func (client *OrganisationClient) RemoveUser(accountId account.AccountId) error {
	route := fmt.Sprintf("/users/%s", accountId)
	req, err := client.base.createRequest(route, http.MethodDelete, nil)
	if err != nil {
		return err
	}

	response, err := client.base.httpClient.Do(req)
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

// Get a list of invites that haven't been accepted.
func (client *OrganisationClient) GetPendingInvites() ([]organisation.OrganisationInvite, error) {
	req, err := client.base.createRequest("/invites", http.MethodGet, nil)
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

	orgInvites := Decode[organisation.OrganisationPendingInvites](response)

	return orgInvites.Invites, nil
}

// Invite a user provided they haven't already been invited.
func (client *OrganisationClient) InviteUser(emailAddress string) error {
	requestBody, err := Encode(organisation.OrganisationInvite{
		EmailAddress: emailAddress,
	})
	if err != nil {
		return err
	}

	req, err := client.base.createRequest("/invites", http.MethodPost, requestBody)
	if err != nil {
		return err
	}

	response, err := client.base.httpClient.Do(req)
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

// Cancel and invite before it's accepted.
func (client *OrganisationClient) CancelUser(emailAddress string) error {
	requestBody, err := Encode(organisation.OrganisationInvite{
		EmailAddress: emailAddress,
	})
	if err != nil {
		return err
	}

	req, err := client.base.createRequest("/invites", http.MethodDelete, requestBody)
	if err != nil {
		return err
	}

	response, err := client.base.httpClient.Do(req)
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
