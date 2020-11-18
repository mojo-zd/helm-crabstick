package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type roleHandler struct {
	client kubernetes.Interface
}

func NewRoleHandler(client kubernetes.Interface) handler.ApiHandler {
	return &roleHandler{client: client}
}

func (s *roleHandler) Get(name, namespace string) (interface{}, error) {
	return s.client.RbacV1().Roles(namespace).Get(context.Background(), name, v1.GetOptions{})
}

func (s *roleHandler) List(namespace string, opts v1.ListOptions) (interface{}, error) {
	return s.client.RbacV1().Roles(namespace).List(context.Background(), opts)
}
