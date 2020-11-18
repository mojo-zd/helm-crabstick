package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type serviceAccountHandler struct {
	client kubernetes.Interface
}

func NewServiceAccountHandler(client kubernetes.Interface) handler.ApiHandler {
	return &serviceAccountHandler{client: client}
}

func (s *serviceAccountHandler) Get(name, namespace string) (interface{}, error) {
	return s.client.CoreV1().ServiceAccounts(namespace).Get(context.Background(), name, v1.GetOptions{})
}

func (s *serviceAccountHandler) List(namespace string, opts v1.ListOptions) (interface{}, error) {
	return s.client.CoreV1().ServiceAccounts(namespace).List(context.Background(), opts)
}
