package storage

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/kube"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

// ActionConfiguration get helm configuration
func ActionConfiguration(client kubernetes.Interface, c config.Config, namespace string) *action.Configuration {
	config := new(action.Configuration)
	config.Log = logrus.Infof
	restClientGetter := &genericclioptions.ConfigFlags{
		KubeConfig:  &c.KubeConf,
		BearerToken: &c.KubeToken,
		Context:     &c.KubeContext,
	}
	config.RESTClientGetter = restClientGetter
	config.KubeClient = kube.New(restClientGetter)
	config.Releases = NewStorage(namespace, client).StoreBackend(StoreBackend(c.StorageBackend))
	return config
}
