package resources

import (
	appV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Resource kubernetes resources operator
type Resource interface {
	ListDeployment(namespace string, opts v1.ListOptions) (*appV1.DeploymentList, error)
	ListSecret(namespace string, opts v1.ListOptions) (*coreV1.SecretList, error)
	ListDaemonSet(namespace string, opts v1.ListOptions) (*appV1.DaemonSetList, error)
	ListConfigMap(namespace string, opts v1.ListOptions) (*coreV1.ConfigMapList, error)
}

type resource struct {
	client kubernetes.Interface
}

// NewResource ...
func NewResource(client kubernetes.Interface) Resource {
	return &resource{
		client: client,
	}
}
