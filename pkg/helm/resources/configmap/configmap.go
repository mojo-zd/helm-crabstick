package configmap

import (
	"context"

	coreV1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Resource interface {
	ListConfigMap(namespace string, opts v1.ListOptions) (*coreV1.ConfigMapList, error)
}

type resource struct {
	client kubernetes.Interface
}

// ListConfigMap ...
func (r *resource) ListConfigMap(namespace string, opts v1.ListOptions) (*coreV1.ConfigMapList, error) {
	return r.client.CoreV1().ConfigMaps(namespace).List(context.Background(), opts)
}
