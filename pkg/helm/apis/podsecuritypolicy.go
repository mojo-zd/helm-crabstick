package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type podSecurityPolicyHandler struct {
	client kubernetes.Interface
}

func NewPodSecurityPolicyHandler(client kubernetes.Interface) handler.ApiHandler {
	return &podSecurityPolicyHandler{client: client}
}

func (s *podSecurityPolicyHandler) Get(name, _ string) (interface{}, error) {
	return s.client.PolicyV1beta1().PodSecurityPolicies().Get(context.Background(), name, v1.GetOptions{})
}

func (s *podSecurityPolicyHandler) List(_ string, opts v1.ListOptions) (interface{}, error) {
	return s.client.PolicyV1beta1().PodSecurityPolicies().List(context.Background(), opts)
}
