package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type daemonSetHandler struct {
	client kubernetes.Interface
}

func NewDaemonSetHandler(client kubernetes.Interface) handler.ApiHandler {
	return &daemonSetHandler{client: client}
}

func (d *daemonSetHandler) Get(name, namespace string) (interface{}, error) {
	return d.client.AppsV1().DaemonSets(namespace).Get(context.Background(), name, v1.GetOptions{})
}

func (d *daemonSetHandler) List(namespace string, opts v1.ListOptions) (interface{}, error) {
	return d.client.AppsV1().DaemonSets(namespace).List(context.Background(), opts)
}
