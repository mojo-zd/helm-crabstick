package manager

import (
	chget "github.com/mojo-zd/helm-crabstick/pkg/helm/chart/action/get"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/manager/kube"
	rldo "github.com/mojo-zd/helm-crabstick/pkg/helm/release/action/do"
	rlget "github.com/mojo-zd/helm-crabstick/pkg/helm/release/action/get"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
)

type appManager struct {
	ChartGetter   chget.Getter
	ReleaseDoer   rldo.Doer
	ReleaseGetter rlget.Getter
}

func NewAppManager(cfg config.Config) *appManager {
	client, err := getClient(cfg)
	if err != nil {
		logrus.Errorf("manager init failed", err)
		return nil
	}

	mgr := &appManager{
		ChartGetter:   chget.NewGetter(cfg),
		ReleaseGetter: rlget.NewGetter(cfg, client, kube.NewApiManager(client)),
		ReleaseDoer:   rldo.NewDoer(client, cfg),
	}

	return mgr
}

func getClient(cfg config.Config) (kubernetes.Interface, error) {
	restcfg, err := cfg.ConfigFlags().ToRESTConfig()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(restcfg)
}
