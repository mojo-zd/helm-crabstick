package storage

import (
	"helm.sh/helm/v3/pkg/storage"
	"helm.sh/helm/v3/pkg/storage/driver"
)

// StoreBackend get helm storage backend
func (s *store) StoreBackend(backend StoreBackend) *storage.Storage {
	var st driver.Driver
	switch backend {
	case SecretBackend:
		st = driver.NewSecrets(s.client.CoreV1().Secrets(s.namespace))
	case ConfigMapBackend:
		st = driver.NewConfigMaps(s.client.CoreV1().ConfigMaps(s.namespace))
	default:
		st = driver.NewSecrets(s.client.CoreV1().Secrets(s.namespace))
	}
	return storage.Init(st)
}
