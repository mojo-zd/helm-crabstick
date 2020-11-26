package do

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"helm.sh/helm/v3/pkg/release"
	"k8s.io/client-go/kubernetes"
)

type DoerOptions struct {
	Annotation map[string]string
}

// Doer include install、uninstall operator
type Doer interface {
	// Install install chart
	Install(chartName, name, namespace, valueString string, opts DoerOptions) (*release.Release, error)

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
