package azureutil_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	az "github.com/samsung-cnct/cma-aks/pkg/util/azureutil"
)

var _ = Describe("azureutil api functions", func() {

	var (
		goodClient az.ClientInterface
		badClient  az.ClientInterface
	)

	BeforeEach(func() {
		goodClient = &MockGoodCMAAKS{}
		badClient = &MockBadCMAAKS{}
	})

	Context("when creating a cluster", func() {
		var (
			err      error
			response az.CreateClusterOutput
		)
		Context("when the apiserver says it is successful", func() {
			BeforeEach(func() {
				response, err = goodClient.CreateCluster(az.CreateClusterInput{})
			})
			It("should return the apiserver message", func() {
				Expect(response.Status).To(Equal(MockGoodCreateClusterOutput.Status))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the apiserver returns an error", func() {
			BeforeEach(func() {
				response, err = badClient.CreateCluster(az.CreateClusterInput{})
			})
			It("should return an empty object", func() {
				Expect(response).To(Equal(az.CreateClusterOutput{}))
			})
			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
			It("should return apiserver error message", func() {
				Expect(err).To(Equal(MockBadCreateClusterError))
			})
		})
	})

	Context("when getting a cluster", func() {
		var (
			err      error
			response az.GetClusterOutput
		)
		Context("when the apiserver says it is successful", func() {
			BeforeEach(func() {
				response, err = goodClient.GetCluster(az.GetClusterInput{})
			})
			It("should return the apiserver message", func() {
				Expect(response.Cluster.ID).To(Equal(MockGoodGetClusterOutput.Cluster.ID))
				Expect(response.Cluster.Name).To(Equal(MockGoodGetClusterOutput.Cluster.Name))
				Expect(response.Cluster.Status).To(Equal(MockGoodGetClusterOutput.Cluster.Status))
				Expect(response.Cluster.Kubeconfig).To(Equal(MockGoodGetClusterOutput.Cluster.Kubeconfig))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the apiserver returns an error", func() {
			BeforeEach(func() {
				response, err = badClient.GetCluster(az.GetClusterInput{})
			})
			It("should return an empty object", func() {
				Expect(response).To(Equal(az.GetClusterOutput{}))
			})
			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
			It("should return apiserver error message", func() {
				Expect(err).To(Equal(MockBadGetClusterError))
			})
		})
	})

	Context("when deleting a cluster", func() {
		var (
			err      error
			response az.DeleteClusterOutput
		)
		Context("when the apiserver says it is successful", func() {
			BeforeEach(func() {
				response, err = goodClient.DeleteCluster(az.DeleteClusterInput{})
			})
			It("should return the apiserver message", func() {
				Expect(response.Status).To(Equal(MockGoodDeleteClusterOutput.Status))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the apiserver returns an error", func() {
			BeforeEach(func() {
				response, err = badClient.DeleteCluster(az.DeleteClusterInput{})
			})
			It("should return an empty object", func() {
				Expect(response).To(Equal(az.DeleteClusterOutput{}))
			})
			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
			It("should return apiserver error message", func() {
				Expect(err).To(Equal(MockBadDeleteClusterError))
			})
		})
	})

	Context("when getting a cluster list", func() {
		var (
			err      error
			response az.ListClusterOutput
		)
		Context("when the apiserver says it is successful", func() {
			BeforeEach(func() {
				response, err = goodClient.ListClusters(az.ListClusterInput{})
			})
			It("should return the apiserver message", func() {
				Expect(response.Clusters).To(HaveLen(2))
				Expect(response.Clusters[0].Name).To(BeEquivalentTo(MockGoodListClusterOutput.Clusters[0].Name))
				Expect(response.Clusters[1].ID).To(BeEquivalentTo(MockGoodListClusterOutput.Clusters[1].ID))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the apiserver returns an error", func() {
			BeforeEach(func() {
				response, err = badClient.ListClusters(az.ListClusterInput{})
			})
			It("should return an empty object", func() {
				Expect(response).To(Equal(az.ListClusterOutput{}))
			})
			It("should return apiserver error message", func() {
				Expect(err).To(Equal(MockBadListClustersError))
			})
		})
	})

	Context("when getting a cluster node count", func() {
		var (
			err      error
			response az.ClusterNodeCountOutput
		)
		Context("when the apiserver says it is successful", func() {
			BeforeEach(func() {
				response, err = goodClient.GetClusterNodeCount(az.ClusterNodeCountInput{})
			})
			It("should return the apiserver message", func() {
				Expect(response.Agent.Name).To(Equal(MockGoodClusterNodeCountOutput.Agent.Name))
				Expect(response.Agent.Count).To(Equal(MockGoodClusterNodeCountOutput.Agent.Count))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the apiserver returns and error", func() {
			BeforeEach(func() {
				response, err = badClient.GetClusterNodeCount(az.ClusterNodeCountInput{})
			})
			It("should return an empty object", func() {
				Expect(response).To(Equal(az.ClusterNodeCountOutput{}))
			})
			It("should return apiserver error message", func() {
				Expect(err).To(Equal(MockBadClusterNodeCountError))
			})
		})
	})

	Context("when scaling a cluster up or down", func() {
		var (
			err      error
			response az.ScaleClusterOutput
		)
		Context("when the apiserver says it is successful", func() {
			BeforeEach(func() {
				response, err = goodClient.ScaleClusterNodeCount(az.ScaleClusterInput{})
			})
			It("should return the apiserver message", func() {
				Expect(response.Status).To(Equal(MockGoodScaleClusterOutput.Status))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the apiserver returns an error", func() {
			BeforeEach(func() {
				response, err = badClient.ScaleClusterNodeCount(az.ScaleClusterInput{})
			})
			It("should return an empty object", func() {
				Expect(response).To(Equal(az.ScaleClusterOutput{}))
			})
			It("should return apiserver error message", func() {
				Expect(err).To(Equal(MockBadScaleClusterError))
			})
		})
	})

	Context("when getting upgrades of a cluster", func() {
		var (
			err      error
			response az.GetClusterUpgradeOutput
		)
		Context("when the apiserver says it was successful", func() {
			BeforeEach(func() {
				response, err = goodClient.GetClusterUpgrades(az.GetClusterUpgradeInput{})
			})
			It("should return the apiserver message", func() {
				Expect(response.Upgrades).To(Equal(MockGoodGetClusterUpgradeOutput.Upgrades))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the apiserver returns an error", func() {
			BeforeEach(func() {
				response, err = badClient.GetClusterUpgrades(az.GetClusterUpgradeInput{})
			})
			It("should return an empty object", func() {
				Expect(response).To(Equal(az.GetClusterUpgradeOutput{}))
			})
			It("should return apiserver error message", func() {
				Expect(err).To(Equal(MockBadGetClusterUpgradeError))
			})
		})
	})

	Context("when upgrading a cluster", func() {
		var (
			err      error
			response az.UpgradeClusterOutput
		)
		Context("when the apiserver says it was successful", func() {
			BeforeEach(func() {
				response, err = goodClient.UpgradeCluster(az.UpgradeClusterInput{})
			})
			It("should return the apiserver message", func() {
				Expect(response.Status).To(Equal(MockGoodUpgradeClusterOutput.Status))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when the apiserver returns an error", func() {
			BeforeEach(func() {
				response, err = badClient.UpgradeCluster(az.UpgradeClusterInput{})
			})
			It("should return an empty object", func() {
				Expect(response).To(Equal(az.UpgradeClusterOutput{}))
			})
			It("should return apiserver error message", func() {
				Expect(err).To(Equal(MockBadUpgradeClusterError))
			})
		})
	})
})
