package get

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/manager/kube"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
	"helm.sh/helm/v3/pkg/release"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
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
	client  kubernetes.Interface
	config  config.Config
	manager *kube.ApiManager
}

// NewGetter ...
func NewGetter(config config.Config, client kubernetes.Interface, mgr *kube.ApiManager) Getter {
	return &getter{
		config:  config,
		client:  client,
		manager: mgr,
	}
}
