// https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/master/resources/groups.go
// https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/master/compute/container_cluster.go

package azureutil

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/go-autorest/autorest/azure/auth"

	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2018-03-31/containerservice"
	"github.com/Azure/go-autorest/autorest/to"

)

// ClusterParameters are used to create a new aks cluster
type ClusterParameters struct {
	Name              string
	Location          string
	KubernetesVersion string
	ClientID          string
	ClientSecret      string
	AgentPools        []Agent
	Tags              map[string]*string
}

// Agent is used to define properties for each instance group
type Agent struct {
	Name    *string
	Count   *int32
	MaxPods *int32
	Type    string
}

// GetClusterClient creates a new aks cluster client and authorizes access
func GetClusterClient(tenantID string, clientID string, clientSecret string, subscriptionID string) (containerservice.ManagedClustersClient, error) {
	// create new cluster client in subscription
	aksClient := containerservice.NewManagedClustersClient(subscriptionID)

	// authorize request
	clientConfig := auth.NewClientCredentialsConfig(clientID, clientSecret, tenantID)
	authorizer, err := clientConfig.Authorizer()
	if err != nil {
		return aksClient, fmt.Errorf("failed to initialize authorizer: %v", err)
	}

	aksClient.Authorizer = authorizer
	return aksClient, nil
}

// GetCluster retrieves a specific aks cluster
func GetCluster(ctx context.Context, clusterClient containerservice.ManagedClustersClient, resourceName string) (c containerservice.ManagedCluster, kubeConfig string, err error) {
	resourceGroupName := resourceName + "-group"

	// get kubeconfig for environment
	credentialResults, err := clusterClient.ListClusterAdminCredentials(ctx, resourceGroupName, resourceName)
	if err != nil {
		log.Printf("err getting cluster credentails: %v", err)
	}
	if credentialResults.Kubeconfigs != nil {
		for _, v := range *credentialResults.Kubeconfigs {
			if *v.Name == "clusterAdmin" {
				kubeConfig = string(*v.Value)
			}
		}
	}

	c, err = clusterClient.Get(ctx, resourceGroupName, resourceName)
	if err != nil {
		fmt.Printf("Error getting cluster %v: %v\n", resourceName, err)
	}

	return c, kubeConfig, err
}

// CreateCluster creates a new managed Kubernetes cluster
func CreateCluster(ctx context.Context, clusterClient containerservice.ManagedClustersClient, parameters ClusterParameters) (status string, err error) {
	resourceGroupName := parameters.Name + "-group"

	// create map of containerservice.ManagedClusterAgentPoolProfile from parameters.AgentPools
	agentPoolProfiles := make([]containerservice.ManagedClusterAgentPoolProfile, len(parameters.AgentPools))
	for i := range parameters.AgentPools {
		var vmSize containerservice.VMSizeTypes

		// get list of available VM size types
		vmSizeTypes := containerservice.PossibleVMSizeTypesValues()
		for j := range vmSizeTypes {
			// convert the vmSizeTypes to a string
			typeAsStr := string(vmSizeTypes[j])
			// compare input type against available vm size types
			if parameters.AgentPools[i].Type == typeAsStr {
				vmSize = vmSizeTypes[j]
			}
		}
		if vmSize == "" {
			return "", fmt.Errorf("invalid VM Size selected")
		}

		agentPoolProfiles[i] = containerservice.ManagedClusterAgentPoolProfile{
			Name:    parameters.AgentPools[i].Name,
			Count:   parameters.AgentPools[i].Count,
			VMSize:  vmSize,
		}
	}

	future, err := clusterClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		parameters.Name,
		containerservice.ManagedCluster{
			Name:     &parameters.Name,
			Location: &parameters.Location,
			ManagedClusterProperties: &containerservice.ManagedClusterProperties{
				DNSPrefix:         &parameters.Name,
				KubernetesVersion: &parameters.KubernetesVersion,
				AgentPoolProfiles: &agentPoolProfiles,
				ServicePrincipalProfile: &containerservice.ManagedClusterServicePrincipalProfile{
					ClientID: to.StringPtr(parameters.ClientID),
					Secret:   to.StringPtr(parameters.ClientSecret),
				},
			},
			Tags: parameters.Tags,
		},
	)
	if err != nil {
		return "", fmt.Errorf("cannot create aks cluster: %v", err)
	}

	status = future.Status()
	if status != "Creating" {
		return "", fmt.Errorf("cannot create cluster: %v", status)
	}

	return status, nil
}

// DeleteCluster deletes an existing aks cluster
func DeleteCluster(ctx context.Context, clusterClient containerservice.ManagedClustersClient, resourceName string) (result string, err error) {
	resourceGroupName := resourceName + "-group"

	future, err := clusterClient.Delete(ctx, resourceGroupName, resourceName)
	if err != nil {
		return result, fmt.Errorf("error deleting cluster: %v", err)
	}

	result = future.Status()
	if result != "InProgress" {
		return "", fmt.Errorf("current status of delete: %v", result)
	}

	msg := "Deleting " + resourceName + " cluster"

	return msg, err

	// TODO: delete resource group also, if nothing else is in it
}

// ListClusters will list all clusters in the subscription
func ListClusters(ctx context.Context, clusterClient containerservice.ManagedClustersClient) ([]containerservice.ManagedCluster, error) {
	results, err := clusterClient.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing clusters: %v", err)
	}

	return results.Values(), nil
}

// GetClusterUpgrades lists the kubernetes upgrades available on the cluster
func GetClusterUpgrades(ctx context.Context, clusterClient containerservice.ManagedClustersClient, resourceName string) ([]string, error) {
	resourceGroupName := resourceName + "-group"

	result, err := clusterClient.GetUpgradeProfile(ctx, resourceGroupName, resourceName)
	if err != nil {
		return nil, fmt.Errorf("error getting available upgrades: %v", err)
	}

	for _, v := range *result.AgentPoolProfiles {
		if v.Upgrades != nil {
			return *v.Upgrades, nil
		}
	}

	return nil, nil
}

// UpgradeCluster upgrades the cluster to the provided kubernetes version
func UpgradeCluster(ctx context.Context, clusterClient containerservice.ManagedClustersClient, parameters ClusterParameters) (status string, err error) {
	resourceGroupName := parameters.Name + "-group"

	// Get the location from cluster properties
	c, err := clusterClient.Get(ctx, resourceGroupName, parameters.Name)
	if err != nil {
		fmt.Printf("Error getting location for cluster %v: %v\n", parameters.Name, err)
	}

	// check cluster status before upgrading
	if *c.ProvisioningState != "Succeeded" {
		return "", fmt.Errorf("Unable to upgrade cluster while it is currently %v", *c.ProvisioningState)
	}

	future, err := clusterClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		parameters.Name,
		containerservice.ManagedCluster{
			Location: c.Location,
			ManagedClusterProperties: &containerservice.ManagedClusterProperties{
				KubernetesVersion: &parameters.KubernetesVersion,
			},
		},
	)
	if err != nil {
		return "", fmt.Errorf("cannot upgrade cluster: %v", err)
	}

	status = future.Status()
	if status != "Upgrading" {
		return "", fmt.Errorf("cannot upgrade cluster: %v", status)
	}

	return status, nil
}

// GetClusterNodeCount returns the current number of nodes in the agent pool
func GetClusterNodeCount(ctx context.Context, clusterClient containerservice.ManagedClustersClient, resourceName string) (agent Agent, err error) {
	resourceGroupName := resourceName + "-group"

	c, err := clusterClient.Get(ctx, resourceGroupName, resourceName)
	if err != nil {
		fmt.Printf("Error getting cluster %v: %v\n", resourceName, err)
	}

	temp := c.ManagedClusterProperties.AgentPoolProfiles
	for _, v := range *temp {
		agent.Name = v.Name
		agent.Count = v.Count
	}
	return agent, nil
}

// ScaleClusterNodeCount sets the total number of nodes based on the count input
func ScaleClusterNodeCount(ctx context.Context, clusterClient containerservice.ManagedClustersClient, resourceName string, nodePool string, count int32) (status string, err error) {
	resourceGroupName := resourceName + "-group"

	// get current cluster
	c, err := clusterClient.Get(ctx, resourceGroupName, resourceName)
	if err != nil {
		fmt.Printf("Error getting cluster %v: %v\n", resourceName, err)
	}

	// check cluster status before scaling
	if *c.ProvisioningState != "Succeeded" {
		return "", fmt.Errorf("Unable to update cluster while it is currently %v", *c.ProvisioningState)
	}

	// get the current VMSize from the cluster
	var vmSize containerservice.VMSizeTypes
	for _, v := range *c.ManagedClusterProperties.AgentPoolProfiles {
		if *v.Name == nodePool {
			vmSize = v.VMSize
		}
	}

	// scale cluster
	future, err := clusterClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		resourceName,
		containerservice.ManagedCluster{
			Location: c.Location,
			ManagedClusterProperties: &containerservice.ManagedClusterProperties{
				AgentPoolProfiles: &[]containerservice.ManagedClusterAgentPoolProfile{
					{
						Name:  to.StringPtr(nodePool),
						Count: to.Int32Ptr(count),
						VMSize: vmSize,
					},
				},
			},
		},
	)
	if err != nil {
		return "", fmt.Errorf("cannot scale cluster: %v", err)
	}

	status = future.Status()
	if status != "Updating" {
		return "", fmt.Errorf("unable to scale: %v", err)
	}

	return status, nil
}
