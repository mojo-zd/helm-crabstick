package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type persistentVolumeHandler struct {
	client kubernetes.Interface
}

func NewPersistentVolumeHandler(client kubernetes.Interface) handler.ApiHandler {
	return &persistentVolumeHandler{client: client}
}

func (s *persistentVolumeHandler) Get(name, _ string) (interface{}, error) {
	return s.client.CoreV1().PersistentVolumes().Get(context.Background(), name, v1.GetOptions{})
}

func (s *persistentVolumeHandler) List(_ string, opts v1.ListOptions) (interface{}, error) {
	return s.client.CoreV1().PersistentVolumes().List(context.Background(), opts)
}
