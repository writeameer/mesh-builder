package main

import (
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/network/mgmt/network"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/writeameer/mesh-builder/helpers"
)

var (
	azureCredential = &helpers.AzureCredential{}
)

// Authenticate with the Azure services using file-based authentication
func main() {

	azureCredential.AuthorizeFromFile("EastAsia")

	group, err := createGroup("AmeerTest")
	if err != nil {
		log.Println(err.Error())

	} else {
		log.Printf("Group %s created in %s", *group.Name, *group.Location)
	}

	myvnet, err := createVNET(group, "TestNET")

	if err != nil {
		log.Println(err.Error())
	} else {
		log.Printf("VNET created %s", *myvnet.Name)
	}

	createSubNet, err := createSubnet(azureCredential, group, "TestNET", "coolsubnet")

	log.Printf("Subnet created: %s", *createSubNet.Name)
}

// Create a resource group for the deployment.
func createGroup(groupName string) (group resources.Group, err error) {

	groupsClient := azureCredential.ResourcesGroupsClient()

	return groupsClient.CreateOrUpdate(
		azureCredential.Ctx,
		groupName,
		resources.Group{
			Location: *&azureCredential.Location,
		},
	)
}

func createVNET(group resources.Group, vnetName string) (vnet network.VirtualNetwork, err error) {

	vnetClient := azureCredential.VirtualNetworksClient()

	future, err := vnetClient.CreateOrUpdate(
		azureCredential.Ctx,
		*group.Name,
		vnetName,
		network.VirtualNetwork{
			Location: group.Location,
			VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
				AddressSpace: &network.AddressSpace{
					AddressPrefixes: &[]string{"10.0.0.0/8"},
				},
			},
		})

	err = future.WaitForCompletionRef(azureCredential.Ctx, vnetClient.Client)
	if err != nil {
		return vnet, fmt.Errorf("cannot get the vnet create or update future response: %v", err)
	}

	return future.Result(vnetClient)
}

func createSubnet(azureCredential *helpers.AzureCredential, group resources.Group, vnetName string, subnetName string) (subnet network.Subnet, err error) {

	subnetsClient := azureCredential.SubnetsClient()
	future, err := subnetsClient.CreateOrUpdate(
		azureCredential.Ctx,
		*group.Name,
		vnetName,
		subnetName,
		network.Subnet{
			SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
				AddressPrefix: to.StringPtr("10.0.0.0/16"),
			},
		})

	err = future.WaitForCompletionRef(azureCredential.Ctx, subnetsClient.Client)
	if err != nil {
		return subnet, fmt.Errorf("cannot get the vnet create or update future response: %v", err)
	}

	return future.Result(subnetsClient)
}
