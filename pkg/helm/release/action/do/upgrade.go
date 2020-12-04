package do

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/storage"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/types"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/release"
)

func (d *doer) Upgrade(opts types.UpgradeOptions) (*release.Release, error) {
	configuration := storage.ActionConfiguration(*d.cluster, d.cfg, opts.Namespace)
	client := action.NewUpgrade(configuration)
	client.Namespace = opts.Namespace
	return d.runUpgrade(opts, client)
}

func (d *doer) runUpgrade(opts types.UpgradeOptions, client *action.Upgrade) (*release.Release, error) {
	setting := util.NewSetting(d.cfg)
	if client.Version == "" && client.Devel {
		logrus.Debug("setting version to >0.0.0-0")
		client.Version = ">0.0.0-0"
	}
	chartOption := action.ChartPathOptions{Version: opts.Version}
	client.ChartPathOptions = chartOption
	cp, err := client.ChartPathOptions.LocateChart(opts.Chart, setting)
	if err != nil {
		return nil, err
	}

	ch, err := loader.Load(cp)
	if err != nil {
		return nil, err
	}
	if req := ch.Metadata.Dependencies; req != nil {
		if err := action.CheckDependencies(ch, req); err != nil {
			return nil, err
		}
	}

	if ch.Metadata.Deprecated {
		logrus.Warn("This chart is deprecated")
	}

	vals, err := util.GetValues(opts.Values)
	if err != nil {
		return nil, err
	}
	rel, err := client.Run(opts.Name, ch, vals)
	if err != nil {
		return nil, errors.Wrap(err, "UPGRADE FAILED")
	}
	return rel, nil
}
