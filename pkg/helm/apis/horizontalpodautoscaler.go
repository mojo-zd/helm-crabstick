package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type horizontalPodAutoscalerHandler struct {
	client kubernetes.Interface
}

func NewHorizontalPodAutoscalerHandler(client kubernetes.Interface) handler.ApiHandler {
	return &horizontalPodAutoscalerHandler{client: client}
}

func (s *horizontalPodAutoscalerHandler) Get(name, namespace string) (interface{}, error) {
	return s.client.AutoscalingV1().HorizontalPodAutoscalers(namespace).Get(context.Background(), name, v1.GetOptions{})
}

func (s *horizontalPodAutoscalerHandler) List(namespace string, opts v1.ListOptions) (interface{}, error) {
	return s.client.AutoscalingV1().HorizontalPodAutoscalers(namespace).List(context.Background(), opts)
}
