package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ingressHandler struct {
	client kubernetes.Interface
}

func NewIngressHandler(client kubernetes.Interface) handler.ApiHandler {
	return &ingressHandler{client: client}
}

func (s *ingressHandler) Get(name, namespace string) (interface{}, error) {
	return s.client.ExtensionsV1beta1().Ingresses(namespace).Get(context.Background(), name, v1.GetOptions{})
}

func (s *ingressHandler) List(namespace string, opts v1.ListOptions) (interface{}, error) {
	return s.client.ExtensionsV1beta1().Ingresses(namespace).List(context.Background(), opts)
}
