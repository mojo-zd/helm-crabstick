package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type serviceHandler struct {
	client kubernetes.Interface
}

func NewServiceHandler(client kubernetes.Interface) handler.ApiHandler {
	return &serviceHandler{client: client}
}

func (s *serviceHandler) Get(name, namespace string) (interface{}, error) {
	return s.client.CoreV1().Services(namespace).Get(context.Background(), name, v1.GetOptions{})
}

func (s *serviceHandler) List(namespace string, opts v1.ListOptions) (interface{}, error) {
	return s.client.CoreV1().Services(namespace).List(context.Background(), v1.ListOptions{})
}
