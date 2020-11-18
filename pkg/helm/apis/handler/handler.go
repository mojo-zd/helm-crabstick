package handler

import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// ApiHandler define common handler to get kubernetes resources
type ApiHandler interface {
	Get(name, namespace string) (interface{}, error)
	List(namespace string, opts v1.ListOptions) (interface{}, error)
}
