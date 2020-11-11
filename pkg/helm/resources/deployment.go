package resources

import (
	"context"
	appV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ListDeployment ...
func (r *resource) ListDeployment(namespace string, opts v1.ListOptions) (*appV1.DeploymentList, error) {
	return r.client.AppsV1().Deployments(namespace).List(
		context.Background(),
		opts,
	)
}
