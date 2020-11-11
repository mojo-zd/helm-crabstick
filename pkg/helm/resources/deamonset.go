package resources

import (
	"context"
	appV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ListDaemonSet ...
func (r *resource) ListDaemonSet(namespace string, opts v1.ListOptions) (*appV1.DaemonSetList, error) {
	return r.client.AppsV1().DaemonSets(namespace).List(context.Background(), opts)
}
