package manager

import (
	"context"

	coreV1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ResourceManager struct {
	client kubernetes.Interface
}

// NewResourceMgr get resource manager instance
func NewResourceMgr(client kubernetes.Interface) *ResourceManager {
	return &ResourceManager{client: client}
}

// ListConfigMap list config map by condition
func (mgr *ResourceManager) ListConfigMap(namespace string, opts v1.ListOptions) (*coreV1.ConfigMapList, error) {
	return mgr.client.CoreV1().ConfigMaps(namespace).List(context.Background(), opts)
}
