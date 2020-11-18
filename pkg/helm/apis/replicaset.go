package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type replicaSetHandler struct {
	client kubernetes.Interface
}

func NewReplicaSetHandler(client kubernetes.Interface) handler.ApiHandler {
	return &replicaSetHandler{client: client}
}

func (s *replicaSetHandler) Get(name, namespace string) (interface{}, error) {
	return s.client.AppsV1().ReplicaSets(namespace).Get(context.Background(), name, v1.GetOptions{})
}

func (s *replicaSetHandler) List(namespace string, opts v1.ListOptions) (interface{}, error) {
	return s.client.AppsV1().ReplicaSets(namespace).List(context.Background(), opts)
}
