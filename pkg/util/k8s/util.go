package k8sutil

import (
	"fmt"
	"io/ioutil"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"

	"github.com/juju/loggo"
	log "github.com/samsung-cnct/cma-aks/pkg/util"
)

var (
	logger loggo.Logger
)

func SetLogger() {
	logger = log.GetModuleLogger("pkg.util.k8sutil", loggo.INFO)
}

func logErrorAndExit(err error) {
	logger.Criticalf("error: %s", err)
	os.Exit(1)
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

// SetKubeConfig sets kubeconfig string to file and returns a kubernetes client config
func SetKubeConfig(name string, kubeconfig string) (*rest.Config, error){
	file, err := ioutil.TempFile(os.TempDir(), name + "-kubeconfig")
	if err != nil {
		return nil, fmt.Errorf("err cerating tempdir: %v", err)
	}
	// delete the file after
	defer os.Remove(file.Name())

	file.WriteString(kubeconfig)

	clusterConfig, err := clientcmd.BuildConfigFromFlags("", file.Name())
	if err != nil {
		return nil, fmt.Errorf("could not load kubeconfig for cluster: %v", err)
	}

	return clusterConfig, nil
}
