package get

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
	"helm.sh/helm/v3/pkg/release"
	"k8s.io/client-go/kubernetes"
)

// Getter helm list、 get operator
type Getter interface {
	List(namespace string, opts util.ListOptions) ([]*release.Release, error)
	Get(name, namespace string) (*release.Release, error)
	Status(name, namespace string) (*release.Release, error)
	ReleaseKind(name, namespace string) []string
}

type getter struct {
	client kubernetes.Interface
	config config.Config
}

// NewGetter ...
func NewGetter(config config.Config, client kubernetes.Interface) Getter {
	return &getter{config: config, client: client}
}
