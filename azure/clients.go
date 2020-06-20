package azure

import (
	"github.com/Azure/azure-sdk-for-go/profiles/latest/network/mgmt/network"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
)

// Clients is a convenience struct for generating clients for Azure APIs
type Clients struct {
	Credential *Credential
	Location   *string
}

// ResourcesGroupsClient Returns a resources.GroupsClient
func (clients *Clients) ResourcesGroupsClient() resources.GroupsClient {

	creds := clients.Credential

	client := resources.NewGroupsClient(*&creds.ServicePrincipal.SubscriptionID)
	client.Authorizer = creds.Authorizer
	return client
}

// VirtualNetworksClient Returns a network.VirtualNetworksClient
func (clients *Clients) VirtualNetworksClient() network.VirtualNetworksClient {

	creds := clients.Credential

	client := (network.NewVirtualNetworksClient(*&creds.ServicePrincipal.SubscriptionID))
	client.Authorizer = creds.Authorizer
	return client
}

// SubnetsClient Returns a network.SubnetsClient
func (clients *Clients) SubnetsClient() network.SubnetsClient {

	creds := clients.Credential

	client := (network.NewSubnetsClient(*&creds.ServicePrincipal.SubscriptionID))
	client.Authorizer = creds.Authorizer

	return client
}

// NewSecurityGroupsClient Returns a network.NewSecurityGroupsClient
func (clients *Clients) NewSecurityGroupsClient() network.SecurityGroupsClient {

	creds := clients.Credential

	client := network.NewSecurityGroupsClient(*&creds.ServicePrincipal.SubscriptionID)
	client.Authorizer = creds.Authorizer
	return client
}
