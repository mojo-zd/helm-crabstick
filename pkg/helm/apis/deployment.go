package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type deployHandler struct {
	client kubernetes.Interface
}

func NewDeployHandler(client kubernetes.Interface) handler.ApiHandler {
	return &deployHandler{client: client}
}

func (d *deployHandler) Get(name, namespace string) (interface{}, error) {
	return d.client.AppsV1().Deployments(namespace).Get(context.Background(), name, v1.GetOptions{})
}

func (d *deployHandler) List(namespace string, opts v1.ListOptions) (interface{}, error) {
	return d.client.AppsV1().Deployments(namespace).List(context.Background(), v1.ListOptions{})
}
