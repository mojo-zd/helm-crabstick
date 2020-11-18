package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type cronJobHandler struct {
	client kubernetes.Interface
}

func NewCronJobHandler(client kubernetes.Interface) handler.ApiHandler {
	return &cronJobHandler{client: client}
}

func (s *cronJobHandler) Get(name, namespace string) (interface{}, error) {
	return s.client.BatchV2alpha1().CronJobs(namespace).Get(context.Background(), name, v1.GetOptions{})
}

func (s *cronJobHandler) List(namespace string, opts v1.ListOptions) (interface{}, error) {
	return s.client.BatchV2alpha1().CronJobs(namespace).List(context.Background(), opts)
}
