package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type clusterRoleBindingHandler struct {
	client kubernetes.Interface
}

func NewClusterRoleBindingHandler(client kubernetes.Interface) handler.ApiHandler {
	return &clusterRoleBindingHandler{client: client}
}

func (s *clusterRoleBindingHandler) Get(name, _ string) (interface{}, error) {
	return s.client.RbacV1().ClusterRoleBindings().Get(context.Background(), name, v1.GetOptions{})
}

func (s *clusterRoleBindingHandler) List(_ string, opts v1.ListOptions) (interface{}, error) {
	return s.client.RbacV1().ClusterRoleBindings().List(context.Background(), opts)
}
