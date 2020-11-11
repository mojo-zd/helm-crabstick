package do

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"helm.sh/helm/v3/pkg/release"
	"k8s.io/client-go/kubernetes"
)

type DoerOptions struct {
	Annotation map[string]string
}

type Doer interface {
	Install(chart, name, namespace string, opts DoerOptions) (*release.Release, error)
}

type doer struct {
	config config.Config
	client kubernetes.Interface
}

func NewDoer(config config.Config, client kubernetes.Interface) Doer {
	return &doer{config: config, client: client}
}
