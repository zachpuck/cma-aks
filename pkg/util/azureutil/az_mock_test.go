package azureutil_test

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2018-03-31/containerservice"
	"github.com/Azure/go-autorest/autorest/to"
	az "github.com/samsung-cnct/cma-aks/pkg/util/azureutil"
)

const (
	MockBadCreateClusterErrorMessage     = "test create cluster error"
	MockBadGetClusterErrorMessage        = "test get cluster error"
	MockBadDeleteClusterErrorMessage     = "test delete cluster error"
	MockBadListClustersErrorMessage      = "test list clusters error"
	MockBadClusterClientErrorMessage     = "test get cluster client error"
	MockBadClusterNodeCountErrorMessage  = "test get cluster node count error"
	MockBadGetClusterUpgradeErrorMessage = "test get cluster upgrade error"
	MockBadScaleClusterErrorMessage      = "test scale cluster error"
	MockBadUpgradeClusterErrorMessage    = "test upgrade cluster error"
)

var (
	MockGoodAKS = az.AKS{
		Client: containerservice.ManagedClustersClient{},
	}
	MockGoodClusterClientOutput = az.ClusterClientOutput{}

	MockGoodCreateClusterOutput = az.CreateClusterOutput{
		Status: "testCreating",
	}
	MockGoodGetClusterOutput = az.GetClusterOutput{
		Cluster: az.ClusterDetailItem{
			ID:         "testClusterID",
			Name:       "testClusterName",
			Status:     "testStatus",
			Kubeconfig: "testKubeConfig",
		},
	}
	MockGoodDeleteClusterOutput = az.DeleteClusterOutput{
		Status: "testDeleting",
	}
	MockGoodClusterNodeCountOutput = az.ClusterNodeCountOutput{
		Agent: az.Agent{
			Name:  to.StringPtr("testClusterName"),
			Count: to.Int32Ptr(int32(2)),
			Type:  "testType",
		},
	}
	MockGoodGetClusterUpgradeOutput = az.GetClusterUpgradeOutput{
		Upgrades: []string{},
	}
	MockGoodListClusterOutput = az.ListClusterOutput{
		Clusters: []containerservice.ManagedCluster{
			{
				ID:   to.StringPtr("Test1ID"),
				Name: to.StringPtr("Test1Name"),
			},
			{
				ID:   to.StringPtr("Test2ID"),
				Name: to.StringPtr("Test2Name"),
			},
		},
	}
	MockGoodScaleClusterOutput = az.ScaleClusterOutput{
		Status: "testScalingStatus",
	}
	MockGoodUpgradeClusterOutput = az.UpgradeClusterOutput{
		Status: "testUpgradeStatus",
	}

	MockBadAKS = az.AKS{
		Client: containerservice.ManagedClustersClient{},
	}

	MockBadCreateClusterError     = fmt.Errorf(MockBadCreateClusterErrorMessage)
	MockBadGetClusterError        = fmt.Errorf(MockBadGetClusterErrorMessage)
	MockBadDeleteClusterError     = fmt.Errorf(MockBadDeleteClusterErrorMessage)
	MockBadListClustersError      = fmt.Errorf(MockBadListClustersErrorMessage)
	MockBadClusterClientError     = fmt.Errorf(MockBadClusterClientErrorMessage)
	MockBadClusterNodeCountError  = fmt.Errorf(MockBadClusterNodeCountErrorMessage)
	MockBadGetClusterUpgradeError = fmt.Errorf(MockBadGetClusterUpgradeErrorMessage)
	MockBadScaleClusterError      = fmt.Errorf(MockBadScaleClusterErrorMessage)
	MockBadUpgradeClusterError    = fmt.Errorf(MockBadUpgradeClusterErrorMessage)
)

type MockGoodCMAAKS struct{}

func (m *MockGoodCMAAKS) SetClient(client containerservice.ManagedClustersClient) *az.AKS {
	return &MockGoodAKS
}

func (m *MockGoodCMAAKS) GetClusterClient(input az.ClusterClientInput) (az.ClusterClientOutput, error) {
	return MockGoodClusterClientOutput, nil
}

func (m *MockGoodCMAAKS) CreateCluster(input az.CreateClusterInput) (az.CreateClusterOutput, error) {
	return MockGoodCreateClusterOutput, nil
}

func (m *MockGoodCMAAKS) GetCluster(input az.GetClusterInput) (az.GetClusterOutput, error) {
	return MockGoodGetClusterOutput, nil
}

func (m *MockGoodCMAAKS) DeleteCluster(input az.DeleteClusterInput) (az.DeleteClusterOutput, error) {
	return MockGoodDeleteClusterOutput, nil
}

func (m *MockGoodCMAAKS) GetClusterNodeCount(input az.ClusterNodeCountInput) (az.ClusterNodeCountOutput, error) {
	return MockGoodClusterNodeCountOutput, nil
}

func (m *MockGoodCMAAKS) GetClusterUpgrades(input az.GetClusterUpgradeInput) (az.GetClusterUpgradeOutput, error) {
	return MockGoodGetClusterUpgradeOutput, nil
}

func (m *MockGoodCMAAKS) ListClusters(input az.ListClusterInput) (az.ListClusterOutput, error) {
	return MockGoodListClusterOutput, nil
}

func (m *MockGoodCMAAKS) ScaleClusterNodeCount(input az.ScaleClusterInput) (az.ScaleClusterOutput, error) {
	return MockGoodScaleClusterOutput, nil
}

func (m *MockGoodCMAAKS) UpgradeCluster(input az.UpgradeClusterInput) (az.UpgradeClusterOutput, error) {
	return MockGoodUpgradeClusterOutput, nil
}

type MockBadCMAAKS struct{}

func (m *MockBadCMAAKS) SetClient(client containerservice.ManagedClustersClient) *az.AKS {
	return &MockBadAKS
}

func (m *MockBadCMAAKS) CreateCluster(input az.CreateClusterInput) (az.CreateClusterOutput, error) {
	return az.CreateClusterOutput{}, MockBadCreateClusterError
}

func (m *MockBadCMAAKS) GetCluster(input az.GetClusterInput) (az.GetClusterOutput, error) {
	return az.GetClusterOutput{}, MockBadGetClusterError
}

func (m *MockBadCMAAKS) DeleteCluster(input az.DeleteClusterInput) (az.DeleteClusterOutput, error) {
	return az.DeleteClusterOutput{}, MockBadDeleteClusterError
}

func (m *MockBadCMAAKS) ListClusters(input az.ListClusterInput) (az.ListClusterOutput, error) {
	return az.ListClusterOutput{}, MockBadListClustersError
}

func (m *MockBadCMAAKS) GetClusterClient(input az.ClusterClientInput) (az.ClusterClientOutput, error) {
	return az.ClusterClientOutput{}, MockBadClusterClientError
}

func (m *MockBadCMAAKS) GetClusterNodeCount(input az.ClusterNodeCountInput) (az.ClusterNodeCountOutput, error) {
	return az.ClusterNodeCountOutput{}, MockBadClusterNodeCountError
}

func (m *MockBadCMAAKS) GetClusterUpgrades(input az.GetClusterUpgradeInput) (az.GetClusterUpgradeOutput, error) {
	return az.GetClusterUpgradeOutput{}, MockBadGetClusterUpgradeError
}

func (m *MockBadCMAAKS) ScaleClusterNodeCount(input az.ScaleClusterInput) (az.ScaleClusterOutput, error) {
	return az.ScaleClusterOutput{}, MockBadScaleClusterError
}

func (m *MockBadCMAAKS) UpgradeCluster(input az.UpgradeClusterInput) (az.UpgradeClusterOutput, error) {
	return az.UpgradeClusterOutput{}, MockBadUpgradeClusterError
}
