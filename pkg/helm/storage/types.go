package storage

import (
	"helm.sh/helm/v3/pkg/storage"
	"k8s.io/client-go/kubernetes"
)

type StoreBackend string

const (
	SecretBackend    = "secrets"
	ConfigMapBackend = "configmap"
)

type Store interface {
	StoreBackend(backend StoreBackend) *storage.Storage
}

type store struct {
	client    kubernetes.Interface
	namespace string
}

// NewStorage new storage instance
func NewStorage(namespace string, client kubernetes.Interface) Store {
	return &store{client: client, namespace: namespace}
}
