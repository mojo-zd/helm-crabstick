package apis

import (
	"context"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/apis/handler"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type networkPolicyHandler struct {
	client kubernetes.Interface
}

func NewNetworkPolicyHandler(client kubernetes.Interface) handler.ApiHandler {
	return &networkPolicyHandler{client: client}
}

func (s *networkPolicyHandler) Get(name, namespace string) (interface{}, error) {
	return s.client.NetworkingV1().NetworkPolicies(namespace).Get(context.Background(), name, v1.GetOptions{})
}

func (s *networkPolicyHandler) List(namespace string, opts v1.ListOptions) (interface{}, error) {
	return s.client.NetworkingV1().NetworkPolicies(namespace).List(context.Background(), opts)
}
