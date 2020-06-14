package main

// import (
// 	"context"
// 	"fmt"

// 	"github.com/Azure/azure-sdk-for-go/profiles/latest/network/mgmt/network"
// 	"github.com/Azure/azure-sdk-for-go/sdk/to"
// )

// func CreateNetworkSecurityGroup(ctx context.Context, nsgName string) (nsg network.SecurityGroup, err error) {
// 	nsgClient := getNsgClient()
// 	future, err := nsgClient.CreateOrUpdate(
// 		ctx,
// 		config.GroupName(),
// 		nsgName,
// 		network.SecurityGroup{
// 			Location: to.StringPtr(config.Location()),
// 			SecurityGroupPropertiesFormat: &network.SecurityGroupPropertiesFormat{
// 				SecurityRules: &[]network.SecurityRule{
// 					{
// 						Name: to.StringPtr("allow_ssh"),
// 						SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
// 							Protocol:                 network.SecurityRuleProtocolTCP,
// 							SourceAddressPrefix:      to.StringPtr("0.0.0.0/0"),
// 							SourcePortRange:          to.StringPtr("1-65535"),
// 							DestinationAddressPrefix: to.StringPtr("0.0.0.0/0"),
// 							DestinationPortRange:     to.StringPtr("22"),
// 							Access:                   network.SecurityRuleAccessAllow,
// 							Direction:                network.SecurityRuleDirectionInbound,
// 							Priority:                 to.Int32Ptr(100),
// 						},
// 					},
// 					{
// 						Name: to.StringPtr("allow_https"),
// 						SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
// 							Protocol:                 network.SecurityRuleProtocolTCP,
// 							SourceAddressPrefix:      to.StringPtr("0.0.0.0/0"),
// 							SourcePortRange:          to.StringPtr("1-65535"),
// 							DestinationAddressPrefix: to.StringPtr("0.0.0.0/0"),
// 							DestinationPortRange:     to.StringPtr("443"),
// 							Access:                   network.SecurityRuleAccessAllow,
// 							Direction:                network.SecurityRuleDirectionInbound,
// 							Priority:                 to.Int32Ptr(200),
// 						},
// 					},
// 				},
// 			},
// 		},
// 	)

// 	if err != nil {
// 		return nsg, fmt.Errorf("cannot create nsg: %v", err)
// 	}

// 	err = future.WaitForCompletionRef(ctx, nsgClient.Client)
// 	if err != nil {
// 		return nsg, fmt.Errorf("cannot get nsg create or update future response: %v", err)
// 	}

// 	return future.Result(nsgClient)
// }

// func getNsgClient() network.SecurityGroupsClient {
// 	nsgClient := network.NewSecurityGroupsClient(config.SubscriptionID())
// 	a, _ := iam.GetResourceManagementAuthorizer()
// 	nsgClient.Authorizer = a
// 	nsgClient.AddToUserAgent(config.UserAgent())
// 	return nsgClient
// }
