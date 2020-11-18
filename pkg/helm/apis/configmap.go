package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type configMapHandler struct {
	client kubernetes.Interface
}

// NewConfigMapHandler return config map resource operator
func NewConfigMapHandler(client kubernetes.Interface) handler.ApiHandler {
	return &configMapHandler{client: client}
}

// Get ...
func (r *configMapHandler) Get(name, namespace string) (interface{}, error) {
	return r.client.CoreV1().ConfigMaps(namespace).Get(context.Background(), name, v1.GetOptions{})
}

// List ...
func (r *configMapHandler) List(namespace string, opts v1.ListOptions) (interface{}, error) {
	return r.client.CoreV1().ConfigMaps(namespace).List(context.Background(), opts)
}
