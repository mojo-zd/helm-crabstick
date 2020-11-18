package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type statefulSetHandler struct {
	client kubernetes.Interface
}

func NewStatefulSetHandler(client kubernetes.Interface) handler.ApiHandler {
	return &statefulSetHandler{client: client}
}

func (s *statefulSetHandler) Get(name, namespace string) (interface{}, error) {
	return s.client.AppsV1().StatefulSets(namespace).Get(context.Background(), name, v1.GetOptions{})
}

func (s *statefulSetHandler) List(namespace string, opts v1.ListOptions) (interface{}, error) {
	return s.client.AppsV1().StatefulSets(namespace).List(context.Background(), opts)
}
