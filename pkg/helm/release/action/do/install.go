package do

import (
	"io"
	"os"

	stg "github.com/mojo-zd/helm-crabstick/pkg/helm/storage"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/types"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/downloader"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/release"
)

// Install install chart
func (d *doer) Install(createOpts types.CreateOptions) (*release.Release, error) {
	setting := util.NewSetting(d.cfg)
	cfg := stg.ActionConfiguration(*d.cluster, d.cfg, createOpts.Namespace)
	install := action.NewInstall(cfg)
	install.Namespace = createOpts.Namespace
	chartOpts := action.ChartPathOptions{Version: createOpts.Version}
	install.ChartPathOptions = chartOpts
	chartObj, err := d.installPre(install, setting, os.Stdout, createOpts.Name, createOpts.Chart)
	if createOpts.Options.Annotation != nil {
		for key, value := range createOpts.Options.Annotation {
			chartObj.Metadata.Annotations[key] = value
		}
	}
	if err != nil {
		return nil, err
	}

	val, err := util.GetValues(createOpts.Values)
	if err != nil {
		return nil, err
	}
	return install.Run(chartObj, val)
}

func (d *doer) installPre(
	client *action.Install,
	settings *cli.EnvSettings,
	out io.Writer,
	args ...string) (*chart.Chart, error) {
	name, chartName, err := client.NameAndChart(args)
	if err != nil {
		return nil, err
	}

	cp, err := client.ChartPathOptions.LocateChart(chartName, settings)
	if err != nil {
		logrus.Errorf("locate chart failed, err:%s", err.Error())
		return nil, err
	}
	logrus.Debugf("chart location [%s]", cp)
	client.ReleaseName = name
	providers := getter.All(settings)

	// Check chart dependencies to make sure all are present in /charts
	chartRequested, err := loader.Load(cp)
	if err != nil {
		logrus.Errorf("load chart failed, err:%s", err.Error())
		return nil, err
	}

	if err := checkIfInstallable(chartRequested); err != nil {
		logrus.Errorf("check installable failed, err:%s", err.Error())
		return nil, err
	}

	if chartRequested.Metadata.Deprecated {
		logrus.Warn("This chart is deprecated")
	}

	if req := chartRequested.Metadata.Dependencies; req != nil {
		// If CheckDependencies returns an error, we have unfulfilled dependencies.
		// As of Helm 2.4.0, this is treated as a stopping condition:
		// https://github.com/helm/helm/issues/2209
		if err := action.CheckDependencies(chartRequested, req); err != nil {
			if client.DependencyUpdate {
				man := &downloader.Manager{
					Out:              out,
					ChartPath:        cp,
					Keyring:          client.ChartPathOptions.Keyring,
					SkipUpdate:       false,
					Getters:          providers,
					RepositoryConfig: settings.RepositoryConfig,
					RepositoryCache:  settings.RepositoryCache,
					Debug:            settings.Debug,
				}
				if err := man.Update(); err != nil {
					return nil, err
				}
				// Reload the chart with the updated Chart.lock file.
				if chartRequested, err = loader.Load(cp); err != nil {
					return nil, errors.Wrap(err, "failed reloading chart after repo update")
				}
			} else {
				return nil, err
			}
		}
	}
	return chartRequested, nil
}

// checkIfInstallable validates if a chart can be installed
// Application chart type is only installable
func checkIfInstallable(ch *chart.Chart) error {
	switch ch.Metadata.Type {
	case "", "application":
		return nil
	}
	return errors.Errorf("%s charts are not installable", ch.Metadata.Type)
}
