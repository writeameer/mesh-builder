package azure

import (
	"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

// Credential Contains Azure client details, an authorizer token and context
type Credential struct {
	ServicePrincipal *ServicePrincipal
	Authorizer       autorest.Authorizer
	Ctx              context.Context
	Location         *string
}

// AuthorizeFromFile Authorizes the Azure API client from file and returns an AzureCredential struct
func (creds *Credential) AuthorizeFromFile(location ...string) {

	defaultLocation := "EastAsia"

	if location != nil {
		defaultLocation = location[0]
	}
	authorizer, err := auth.NewAuthorizerFromFile(azure.PublicCloud.ResourceManagerEndpoint)
	if err != nil {
		log.Fatalf("Failed to get OAuth config: %v", err)
	}

	servicePrincipal, err := readJSON(os.Getenv("AZURE_AUTH_LOCATION"))

	if err != nil {
		log.Fatalf("Failed to read JSON: %+v", err)
	}

	creds.ServicePrincipal = servicePrincipal
	creds.Authorizer = authorizer
	creds.Ctx = context.Background()
	creds.Location = &defaultLocation
}

// ResourcesGroupsClient Returns a resources.GroupsClient
func (creds *Credential) ResourcesGroupsClient() resources.GroupsClient {
	client := (resources.NewGroupsClient(*&creds.ServicePrincipal.SubscriptionID))
	client.Authorizer = creds.Authorizer
	return client
}
