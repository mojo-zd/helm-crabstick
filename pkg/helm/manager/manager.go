package manager

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
	"k8s.io/client-go/kubernetes"
)

type ApiManager struct {
	client    kubernetes.Interface
	resources map[util.KubeKind]handler.ApiHandler
}

// NewApiManager expose kubernetes resources api
func NewApiManager(client kubernetes.Interface) *ApiManager {
	manager := &ApiManager{
		client:    client,
		resources: make(map[util.KubeKind]handler.ApiHandler),
	}
	manager.registry()
	return manager
}

func (m *ApiManager) registry() {
	m.resources[util.ConfigMap] = apis.NewConfigMapHandler(m.client)
	m.resources[util.DaemonSet] = apis.NewDaemonSetHandler(m.client)
	m.resources[util.Deploy] = apis.NewDeployHandler(m.client)
	m.resources[util.Secret] = apis.NewSecretHandler(m.client)
	m.resources[util.Service] = apis.NewServiceHandler(m.client)
	m.resources[util.ServiceAccount] = apis.NewServiceAccountHandler(m.client)
}

func (m *ApiManager) GetResources(kind util.KubeKind) handler.ApiHandler {
	return m.resources[kind]
}
