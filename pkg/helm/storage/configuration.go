package storage

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/mojo-zd/helm-crabstick/pkg/manager"
	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/kube"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// ActionConfiguration get helm configuration
func ActionConfiguration(cluster manager.Cluster, c config.Config, namespace string) *action.Configuration {
	config := new(action.Configuration)
	config.Log = logrus.Infof
	n := ""
	restClientGetter := &genericclioptions.ConfigFlags{
		APIServer:  &cluster.ApiAddress,
		CAFile:     &cluster.CAFile,
		KeyFile:    &cluster.KeyFile,
		CertFile:   &cluster.CertFile,
		Namespace:  &namespace,
		KubeConfig: &n,
	}

	config.RESTClientGetter = restClientGetter
	config.KubeClient = kube.New(restClientGetter)
	config.Releases = NewStorage(namespace, cluster.Client).StoreBackend(StoreBackend(c.Repository.StorageBackend))
	return config
}
