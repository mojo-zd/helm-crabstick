package get

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
	"github.com/mojo-zd/helm-crabstick/pkg/manager"
	"helm.sh/helm/v3/pkg/release"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Getter helm list、 get operator
type Getter interface {
	List(namespace string, opts util.ListOptions) ([]*release.Release, error)
	Get(name, namespace string) (*release.Release, error)
	Status(name, namespace string) (*release.Release, error)
	Kind(rls *release.Release) []string
	Resources(name, namespace string, rls *release.Release, opts v1.ListOptions) map[util.KubeKind]interface{}
	History(name, namespace string) (ReleaseHistory, error)
}

type getter struct {
	cluster *manager.Cluster
	config  config.Config
	manager *manager.ApiManager
}

// NewGetter ...
func NewGetter(config config.Config, cluster *manager.Cluster, mgr *manager.ApiManager) Getter {
	return &getter{
		config:  config,
		cluster: cluster,
		manager: mgr,
	}
}
