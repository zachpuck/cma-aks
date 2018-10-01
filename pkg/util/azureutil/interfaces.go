package azureutil

import (
	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2018-03-31/containerservice"
)

type ClientInterface interface {
	GetClusterClient(ClusterClientInput) (ClusterClientOutput, error)
	CreateCluster(CreateClusterInput) (CreateClusterOutput, error)
	GetCluster(GetClusterInput) (GetClusterOutput, error)
	SetClient(client containerservice.ManagedClustersClient)
	DeleteCluster(DeleteClusterInput) (DeleteClusterOutput, error)
	ListClusters(ListClusterInput) (ListClusterOutput, error)
	GetClusterUpgrades(GetClusterUpgradeInput) (GetClusterUpgradeOutput, error)
	UpgradeCluster(UpgradeClusterInput) (UpgradeClusterOutput, error)
	GetClusterNodeCount(ClusterNodeCountInput) (ClusterNodeCountOutput, error)
	ScaleClusterNodeCount(ScaleClusterInput) (ScaleClusterOutput, error)
}
