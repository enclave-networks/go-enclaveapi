package client

import "net/http"

func (client *EnclaveClient) GetSystems() error {
	url := client.getRequestUrl("/systems")
	request, err := http.NewRequest("GET", url.String(), nil)

	if err != nil {
		return err
	}

	client.setRequestHeader(request)

	return nil
}
