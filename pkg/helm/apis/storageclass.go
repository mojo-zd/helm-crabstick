package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type storageClassHandler struct {
	client kubernetes.Interface
}

func NewStorageClassHandler(client kubernetes.Interface) handler.ApiHandler {
	return &storageClassHandler{client: client}
}

func (s *storageClassHandler) Get(name, _ string) (interface{}, error) {
	return s.client.StorageV1().StorageClasses().Get(context.Background(), name, v1.GetOptions{})
}

func (s *storageClassHandler) List(_ string, opts v1.ListOptions) (interface{}, error) {
	return s.client.StorageV1().StorageClasses().List(context.Background(), opts)
}
