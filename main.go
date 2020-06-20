package main

import (
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/network/mgmt/network"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/writeameer/mesh-builder/azure"
)

var (
	azureClients = &azure.Clients{
		Credential: &azure.Credential{},
		Location:   to.StringPtr("EastAsia"),
	}
)

// Authenticate with the Azure services using file-based authentication
func main() {

	azureClients.Credential.AuthorizeFromFile()

	group, _ := createGroup("AmeerTest")
	log.Printf("Group %s created in %s", *group.Name, *group.Location)

	myvnet, _ := createVNET(group, "TestNET")
	log.Printf("VNET created %s", *myvnet.Name)

	createSubNet, _ := createSubnet(group, "TestNET", "coolsubnet")
	log.Printf("Subnet created: %s", *createSubNet.Name)

	myNSG, _ := createNSG(group, "TestNSG")
	log.Printf("Subnet created: %s", *myNSG.Name)

}

// Create a resource group for the deployment.
func createGroup(groupName string) (group resources.Group, err error) {

	client := azureClients.ResourcesGroupsClient()

	return client.CreateOrUpdate(
		azureClients.Credential.Ctx,
		groupName,
		resources.Group{
			Location: *&azureClients.Location,
		},
	)
}

func createVNET(group resources.Group, vnetName string) (vnet network.VirtualNetwork, err error) {

	client := azureClients.VirtualNetworksClient()

	future, err := client.CreateOrUpdate(
		azureClients.Credential.Ctx,
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

	err = future.WaitForCompletionRef(azureClients.Credential.Ctx, client.Client)
	if err != nil {
		return vnet, fmt.Errorf("cannot get the vnet create or update future response: %v", err)
	}

	return future.Result(client)
}

func createSubnet(group resources.Group, vnetName string, subnetName string) (subnet network.Subnet, err error) {

	client := azureClients.SubnetsClient()
	future, err := client.CreateOrUpdate(
		azureClients.Credential.Ctx,
		*group.Name,
		vnetName,
		subnetName,
		network.Subnet{
			SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
				AddressPrefix: to.StringPtr("10.0.0.0/16"),
			},
		})

	err = future.WaitForCompletionRef(azureClients.Credential.Ctx, client.Client)
	if err != nil {
		return subnet, fmt.Errorf("cannot get the subnet create or update future response: %v", err)
	}

	return future.Result(client)
}

func createNSG(group resources.Group, nsgName string) (nsg network.SecurityGroup, err error) {

	client := azureClients.NewSecurityGroupsClient()
	future, err := client.CreateOrUpdate(
		azureClients.Credential.Ctx,
		*group.Name,
		nsgName,
		network.SecurityGroup{
			Location: to.StringPtr(*group.Location),
		},
	)

	err = future.WaitForCompletionRef(azureClients.Credential.Ctx, client.Client)
	if err != nil {
		return nsg, fmt.Errorf("cannot get the nsg create or update future response: %v", err)
	}

	return future.Result(client)
}
