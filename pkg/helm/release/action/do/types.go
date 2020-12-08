package do

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/types"
	"github.com/mojo-zd/helm-crabstick/pkg/manager"
	"helm.sh/helm/v3/pkg/release"
)

// Doer include install、uninstall operator
type Doer interface {
	// Install install chart
	Install(createOpt types.CreateOptions) (*release.Release, error)

	// DeleteBathes uninstall release
	Delete(name, namespace string) (*release.UninstallReleaseResponse, error)

	// Upgrade upgrade release
	Upgrade(opts types.UpgradeOptions) (*release.Release, error)
}

type doer struct {
	cluster *manager.Cluster
	cfg     config.Config
}

// NewDoer ...
func NewDoer(cluster *manager.Cluster, conf config.Config) Doer {
	return &doer{
		cluster: cluster,
		cfg:     conf,
	}
}
