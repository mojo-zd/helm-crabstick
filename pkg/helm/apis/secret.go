package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type secretHandler struct {
	client kubernetes.Interface
}

func NewSecretHandler(client kubernetes.Interface) handler.ApiHandler {
	return &secretHandler{client: client}
}

func (s *secretHandler) Get(name, namespace string) (interface{}, error) {
	return s.client.CoreV1().Secrets(namespace).Get(context.Background(), name, v1.GetOptions{})
}

func (s *secretHandler) List(namespace string, opts v1.ListOptions) (interface{}, error) {
	return s.client.CoreV1().Secrets(namespace).List(context.Background(), opts)
}
