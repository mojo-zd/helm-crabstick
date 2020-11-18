package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type resourceQuotaHandler struct {
	client kubernetes.Interface
}

func NewResourceQuotaHandler(client kubernetes.Interface) handler.ApiHandler {
	return &resourceQuotaHandler{client: client}
}

func (s *resourceQuotaHandler) Get(name, namespace string) (interface{}, error) {
	return s.client.CoreV1().ResourceQuotas(namespace).Get(context.Background(), name, v1.GetOptions{})
}

func (s *resourceQuotaHandler) List(namespace string, opts v1.ListOptions) (interface{}, error) {
	return s.client.CoreV1().ResourceQuotas(namespace).List(context.Background(), opts)
}
