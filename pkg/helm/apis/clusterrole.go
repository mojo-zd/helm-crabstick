package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type clusterRoleHandler struct {
	client kubernetes.Interface
}

func NewClusterRoleHandler(client kubernetes.Interface) handler.ApiHandler {
	return &clusterRoleHandler{client: client}
}

func (s *clusterRoleHandler) Get(name, _ string) (interface{}, error) {
	return s.client.RbacV1().ClusterRoles().Get(context.Background(), name, v1.GetOptions{})
}

func (s *clusterRoleHandler) List(_ string, opts v1.ListOptions) (interface{}, error) {
	return s.client.RbacV1().ClusterRoles().List(context.Background(), opts)
}
