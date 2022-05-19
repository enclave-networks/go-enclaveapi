package enclave

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/enclave-networks/go-enclaveapi/data"
	"github.com/enclave-networks/go-enclaveapi/data/dns"
)

type DnsClient struct {
	base *ClientBase
}

func (client *DnsClient) GetPropertiesSummary() (dns.DnsSummary, error) {
	req, err := client.base.createRequest("/dns", http.MethodGet, nil)
	if err != nil {
		return dns.DnsSummary{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return dns.DnsSummary{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return dns.DnsSummary{}, err
	}

	dnsSummary := Decode[dns.DnsSummary](response)

	return *dnsSummary, nil
}

func (client *DnsClient) GetZones(pageNumber *int, perPage *int) ([]dns.DnsZoneSummary, error) {
	req, err := client.base.createRequest("/dns/zones", http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	buildDnsQuery(req, pageNumber, perPage)

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return nil, err
	}

	dnsZones := Decode[data.PaginatedResponse[dns.DnsZoneSummary]](response)

	return dnsZones.Items, nil
}

func (client *DnsClient) CreateZone(create dns.DnsZoneCreate) (dns.DnsZone, error) {
	body, err := Encode(create)
	if err != nil {
		return dns.DnsZone{}, err
	}

	req, err := client.base.createRequest("/dns/zones", http.MethodPost, body)
	if err != nil {
		return dns.DnsZone{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return dns.DnsZone{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return dns.DnsZone{}, err
	}

	dnsZone := Decode[dns.DnsZone](response)

	return *dnsZone, nil
}

func (client *DnsClient) GetZone(dnsZoneId dns.DnsZoneId) (dns.DnsZone, error) {
	route := fmt.Sprintf("/dns/zones/%v", dnsZoneId)
	req, err := client.base.createRequest(route, http.MethodGet, nil)
	if err != nil {
		return dns.DnsZone{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return dns.DnsZone{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return dns.DnsZone{}, err
	}

	dnsZone := Decode[dns.DnsZone](response)

	return *dnsZone, nil
}

func (client *DnsClient) UpdateZone(dnsZoneId dns.DnsZoneId, patch dns.DnsZonePatch) (dns.DnsZone, error) {
	body, err := Encode(patch)
	if err != nil {
		return dns.DnsZone{}, err
	}

	route := fmt.Sprintf("/dns/zones/%v", dnsZoneId)
	req, err := client.base.createRequest(route, http.MethodPatch, body)
	if err != nil {
		return dns.DnsZone{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return dns.DnsZone{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return dns.DnsZone{}, err
	}

	dnsZone := Decode[dns.DnsZone](response)

	return *dnsZone, nil
}

func (client *DnsClient) DeleteZone(dnsZoneId dns.DnsZoneId) (dns.DnsZone, error) {
	route := fmt.Sprintf("/dns/zones/%v", dnsZoneId)
	req, err := client.base.createRequest(route, http.MethodDelete, nil)
	if err != nil {
		return dns.DnsZone{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return dns.DnsZone{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return dns.DnsZone{}, err
	}

	dnsZone := Decode[dns.DnsZone](response)

	return *dnsZone, nil
}

//TODO: Need to add more values here
func (client *DnsClient) GetRecords(pageNumber *int, perPage *int) ([]dns.DnsRecordSummary, error) {
	req, err := client.base.createRequest("/dns/records", http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	buildDnsQuery(req, pageNumber, perPage)

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return nil, err
	}

	dnsRecords := Decode[data.PaginatedResponse[dns.DnsRecordSummary]](response)

	return dnsRecords.Items, nil
}

func (client *DnsClient) CreateRecord(create dns.DnsRecordCreate) (dns.DnsRecord, error) {
	// If we haven't set a type set it to ENCLAVE as this is the default
	if len(create.Type) == 0 {
		create.Type = "ENCLAVE"
	}

	body, err := Encode(create)
	if err != nil {
		return dns.DnsRecord{}, err
	}

	req, err := client.base.createRequest("/dns/records", http.MethodPost, body)
	if err != nil {
		return dns.DnsRecord{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return dns.DnsRecord{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return dns.DnsRecord{}, err
	}

	dnsRecord := Decode[dns.DnsRecord](response)

	return *dnsRecord, nil
}

func (client *DnsClient) GetRecord(dnsRecordId dns.DnsRecordId) (dns.DnsRecord, error) {
	route := fmt.Sprintf("/dns/records/%v", dnsRecordId)
	req, err := client.base.createRequest(route, http.MethodGet, nil)
	if err != nil {
		return dns.DnsRecord{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return dns.DnsRecord{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return dns.DnsRecord{}, err
	}

	dnsRecord := Decode[dns.DnsRecord](response)

	return *dnsRecord, nil
}

func (client *DnsClient) UpdateRecord(dnsRecordId dns.DnsRecordId, patch dns.DnsRecordPatch) (dns.DnsRecord, error) {
	body, err := Encode(patch)
	if err != nil {
		return dns.DnsRecord{}, err
	}

	route := fmt.Sprintf("/dns/records/%v", dnsRecordId)
	req, err := client.base.createRequest(route, http.MethodPatch, body)
	if err != nil {
		return dns.DnsRecord{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return dns.DnsRecord{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return dns.DnsRecord{}, err
	}

	dnsRecord := Decode[dns.DnsRecord](response)

	return *dnsRecord, nil
}

func (client *DnsClient) DeleteRecord(dnsRecordId dns.DnsRecordId) (dns.DnsRecord, error) {
	route := fmt.Sprintf("/dns/records/%v", dnsRecordId)
	req, err := client.base.createRequest(route, http.MethodDelete, nil)
	if err != nil {
		return dns.DnsRecord{}, err
	}

	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return dns.DnsRecord{}, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return dns.DnsRecord{}, err
	}

	dnsRecord := Decode[dns.DnsRecord](response)

	return *dnsRecord, nil
}

func (client *DnsClient) BulkDeleteRecord(recordIds ...dns.DnsRecordId) (int, error) {
	if recordIds == nil {
		err := fmt.Errorf("no record Ids")
		return 0, err
	}

	requestBody, err := Encode(dns.DnsBulkAction{RecordIds: recordIds})
	if err != nil {
		return -1, err
	}

	req, err := client.base.createRequest("/dns/records", http.MethodDelete, requestBody)
	if err != nil {
		return -1, err
	}
	response, err := client.base.httpClient.Do(req)
	if err != nil {
		return -1, err
	}
	defer response.Body.Close()

	err = isSuccessStatusCode(response.StatusCode)
	if err != nil {
		return -1, err
	}

	result := Decode[dns.DnsRecordBulkDeleteResult](response)

	return result.DnsRecordsDeleted, nil
}

func buildDnsQuery(req *http.Request, pageNumber *int, perPage *int) {
	query := req.URL.Query()

	if pageNumber != nil {
		query.Add("page", strconv.FormatInt(int64(*pageNumber), 10))
	}

	if perPage != nil {
		query.Add("per_page", strconv.FormatInt(int64(*perPage), 10))
	}
}
