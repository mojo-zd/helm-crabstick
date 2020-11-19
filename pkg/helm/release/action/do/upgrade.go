package do

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/storage"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/release"
)

func (d *doer) Upgrade(release, chart, version, values, namespace string) (*release.Release, error) {
	configuration := storage.ActionConfiguration(d.client, d.cfg, namespace)
	client := action.NewUpgrade(configuration)
	client.Namespace = namespace
	return d.runUpgrade(release, chart, version, values, client)
}

func (d *doer) runUpgrade(release, chart, version, values string, client *action.Upgrade) (*release.Release, error) {
	setting := util.NewSetting(d.cfg)
	if client.Version == "" && client.Devel {
		logrus.Debug("setting version to >0.0.0-0")
		client.Version = ">0.0.0-0"
	}
	chartOption, err := util.LoadChartOptions(d.cfg)
	if err != nil {
		logrus.Errorf("instance load chart options failed, err:%s", err.Error())
		return nil, err
	}
	chartOption.Version = version
	client.ChartPathOptions = chartOption
	cp, err := client.ChartPathOptions.LocateChart(chart, setting)
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

	vals, err := util.GetValues(values)
	if err != nil {
		return nil, err
	}
	rel, err := client.Run(release, ch, vals)
	if err != nil {
		return nil, errors.Wrap(err, "UPGRADE FAILED")
	}
	return rel, nil
}
