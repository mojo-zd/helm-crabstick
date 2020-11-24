package manager

import (
	chdo "github.com/mojo-zd/helm-crabstick/pkg/helm/chart/action/do"
	chget "github.com/mojo-zd/helm-crabstick/pkg/helm/chart/action/get"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	rldo "github.com/mojo-zd/helm-crabstick/pkg/helm/release/action/do"
	rlget "github.com/mojo-zd/helm-crabstick/pkg/helm/release/action/get"
)

type appManager struct {
	ChartDoer     chdo.Doer
	ChartGetter   chget.Getter
	ReleaseDoer   rldo.Doer
	ReleaseGetter rlget.Getter
}

func NewAppManager(cfg config.Config) *appManager {
	mgr := &appManager{
		ChartGetter: chget.NewGetter(cfg),
	}
	return mgr
}
