package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type jobHandler struct {
	client kubernetes.Interface
}

func NewJobHandler(client kubernetes.Interface) handler.ApiHandler {
	return &jobHandler{client: client}
}

func (s *jobHandler) Get(name, namespace string) (interface{}, error) {
	return s.client.BatchV1().Jobs(namespace).Get(context.Background(), name, v1.GetOptions{})
}

func (s *jobHandler) List(namespace string, opts v1.ListOptions) (interface{}, error) {
	return s.client.BatchV1().Jobs(namespace).List(context.Background(), opts)
}
