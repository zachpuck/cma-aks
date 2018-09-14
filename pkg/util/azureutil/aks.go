// https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/master/resources/groups.go
// https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/master/compute/container_cluster.go

package azureutil

import (
	"context"
	"fmt"
	"log"
	"os"

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

	// used to authenticate with azure
	// sets the environment variables expected by NewAuthorizerFromEnvironment()
	os.Setenv("AZURE_TENANT_ID", tenantID)
	os.Setenv("AZURE_CLIENT_ID", clientID)
	os.Setenv("AZURE_CLIENT_SECRET", clientSecret)

	// create new cluster client in subscription
	aksClient := containerservice.NewManagedClustersClient(subscriptionID)

	// create an authorizer from env vars
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		log.Fatalf("failed to initialize authorizer: %v\n", err)
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
			MaxPods: parameters.AgentPools[i].MaxPods,
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
