package azureutil_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAzureutil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Azureutil Suite")
}
