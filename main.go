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

	createSubNet, err := createSubnet(group, "TestNET", "coolsubnet")
	log.Printf("Subnet created: %s", *createSubNet.Name)

	myNSG, err := createNSG(group, "TestNSG")
	log.Printf("Subnet created: %s", *myNSG.Name)

}

// Create a resource group for the deployment.
func createGroup(groupName string) (group resources.Group, err error) {

	client := azureCredential.ResourcesGroupsClient()

	return client.CreateOrUpdate(
		azureCredential.Ctx,
		groupName,
		resources.Group{
			Location: *&azureCredential.Location,
		},
	)
}

func createVNET(group resources.Group, vnetName string) (vnet network.VirtualNetwork, err error) {

	client := azureCredential.VirtualNetworksClient()

	future, err := client.CreateOrUpdate(
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

	err = future.WaitForCompletionRef(azureCredential.Ctx, client.Client)
	if err != nil {
		return vnet, fmt.Errorf("cannot get the vnet create or update future response: %v", err)
	}

	return future.Result(client)
}

func createSubnet(group resources.Group, vnetName string, subnetName string) (subnet network.Subnet, err error) {

	client := azureCredential.SubnetsClient()
	future, err := client.CreateOrUpdate(
		azureCredential.Ctx,
		*group.Name,
		vnetName,
		subnetName,
		network.Subnet{
			SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
				AddressPrefix: to.StringPtr("10.0.0.0/16"),
			},
		})

	err = future.WaitForCompletionRef(azureCredential.Ctx, client.Client)
	if err != nil {
		return subnet, fmt.Errorf("cannot get the subnet create or update future response: %v", err)
	}

	return future.Result(client)
}

func createNSG(group resources.Group, nsgName string) (nsg network.SecurityGroup, err error) {

	client := azureCredential.NewSecurityGroupsClient()
	future, err := client.CreateOrUpdate(
		azureCredential.Ctx,
		*group.Name,
		nsgName,
		network.SecurityGroup{
			Location: to.StringPtr(*group.Location),
		},
	)

	err = future.WaitForCompletionRef(azureCredential.Ctx, client.Client)
	if err != nil {
		return nsg, fmt.Errorf("cannot get the nsg create or update future response: %v", err)
	}

	return future.Result(client)
}
