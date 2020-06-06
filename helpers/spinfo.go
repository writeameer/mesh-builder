package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// readJSON Reads json and returns a map
func readJSON(path string) (*SPInfo, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	contents := make(map[string]string)
	json.Unmarshal(data, &contents)

	spInfo := &SPInfo{
		ClientID:                   contents["clientId"],
		ClientSecret:               contents["clientSecret"],
		SubscriptionID:             contents["subscriptionId"],
		TenantID:                   contents["tenantId"],
		ActiveDirectoryEndPointURL: contents["activeDirectoryEndpointUrl"],
		ManagementEndpointURL:      contents["managementEndpointUrl"],
	}

	return spInfo, nil
}

// SPInfo Contatins details about the service principal used to authenticate with Azure
type SPInfo struct {
	ClientID                   string
	ClientSecret               string
	SubscriptionID             string
	TenantID                   string
	ActiveDirectoryEndPointURL string
	ManagementEndpointURL      string
}
