package get

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/storage"
	ac "helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/release"
)

// Get get release
func (g *getter) Get(name, namespace string) (*release.Release, error) {
	cfg := storage.ActionConfiguration(*g.cluster, g.config, namespace)
	get := ac.NewGet(cfg)
	return get.Run(name)
}
