package do

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"helm.sh/helm/v3/pkg/release"
	"k8s.io/client-go/kubernetes"
)

// Doer include install、uninstall operator
type Doer interface {
	// Delete uninstall release
	Delete(name, namespace string) (*release.UninstallReleaseResponse, error)

	// Upgrade upgrade release
	Upgrade(release, chart, version, values, namespace string) (*release.Release, error)
}

type doer struct {
	client kubernetes.Interface
	cfg    config.Config
}

// NewDoer ...
func NewDoer(client kubernetes.Interface, conf config.Config) Doer {
	return &doer{
		client: client,
		cfg:    conf,
	}
}
