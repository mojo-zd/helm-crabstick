package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type persistentVolumeClaimHandler struct {
	client kubernetes.Interface
}

func NewPersistentVolumeClaimHandler(client kubernetes.Interface) handler.ApiHandler {
	return &persistentVolumeClaimHandler{client: client}
}

func (s *persistentVolumeClaimHandler) Get(name, namespace string) (interface{}, error) {
	return s.client.CoreV1().PersistentVolumeClaims(namespace).Get(context.Background(), name, v1.GetOptions{})
}

func (s *persistentVolumeClaimHandler) List(namespace string, opts v1.ListOptions) (interface{}, error) {
	return s.client.CoreV1().PersistentVolumeClaims(namespace).List(context.Background(), opts)
}
