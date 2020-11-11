package resources

import "k8s.io/apimachinery/pkg/labels"

var defLabelSelector = labels.NewSelector()

type ApiKind string

const (
	Deployment = "deployments"
	Deamonset  = "deamonset"
)
