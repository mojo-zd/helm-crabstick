package manager

import (
	chget "github.com/mojo-zd/helm-crabstick/pkg/helm/chart/action/get"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/manager/kube"
	rldo "github.com/mojo-zd/helm-crabstick/pkg/helm/release/action/do"
	rlget "github.com/mojo-zd/helm-crabstick/pkg/helm/release/action/get"
	"github.com/mojo-zd/helm-crabstick/pkg/manager"
)

type appManager struct {
	ChartGetter   chget.Getter
	ReleaseDoer   rldo.Doer
	ReleaseGetter rlget.Getter
}

// NewChartManager only support chart operator
func NewChartManager(cfg config.Config) *appManager {
	return &appManager{ChartGetter: chget.NewGetter(cfg)}
}

// NewAppManager support all operator of chart、release
func NewAppManager(cfg config.Config, cluster *manager.Cluster) *appManager {
	return &appManager{
		ChartGetter:   chget.NewGetter(cfg),
		ReleaseGetter: rlget.NewGetter(cfg, cluster, kube.NewApiManager(cluster.Client)),
		ReleaseDoer:   rldo.NewDoer(cluster, cfg),
	}
}
