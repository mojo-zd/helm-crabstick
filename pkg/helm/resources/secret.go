package resources

import (
	"context"
	coreV1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ListSecret ...
func (r *resource) ListSecret(namespace string, opts v1.ListOptions) (*coreV1.SecretList, error) {
	return r.client.CoreV1().Secrets(namespace).List(
		context.Background(),
		opts,
	)
}
