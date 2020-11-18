package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type roleBindingHandler struct {
	client kubernetes.Interface
}

func NewRoleBindingHandler(client kubernetes.Interface) handler.ApiHandler {
	return &roleBindingHandler{client: client}
}

func (s *roleBindingHandler) Get(name, namespace string) (interface{}, error) {
	return s.client.RbacV1().RoleBindings(namespace).Get(context.Background(), name, v1.GetOptions{})
}

func (s *roleBindingHandler) List(namespace string, opts v1.ListOptions) (interface{}, error) {
	return s.client.RbacV1().RoleBindings(namespace).List(context.Background(), opts)
}
