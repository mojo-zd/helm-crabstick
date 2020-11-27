package do

import (
	"io"
	"os"

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
	"helm.sh/helm/v3/pkg/kube"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/storage"
	"helm.sh/helm/v3/pkg/storage/driver"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/rest"
)

// Install install chart
func (d *doer) Install(createOpts types.CreateOptions) (*release.Release, error) {
	setting := util.NewSetting(d.cfg)
	cfg := d.buildActionConfiguration(createOpts.Namespace)
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

func (d *doer) buildActionConfiguration(namespace string) *action.Configuration {
	secrets := driver.NewSecrets(d.client.CoreV1().Secrets(namespace))
	secrets.Log = logrus.Infof
	store := storage.Init(secrets)

	actionConfig := new(action.Configuration)
	config, _ := d.cfg.ConfigFlags().ToRESTConfig()
	restClientGetter := NewConfigFlagsFromCluster(namespace, config)
	actionConfig.RESTClientGetter = restClientGetter
	actionConfig.KubeClient = kube.New(restClientGetter)
	actionConfig.Releases = store
	actionConfig.Log = logrus.Infof
	return actionConfig
}

// NewConfigFlagsFromCluster returns ConfigFlags with default values set from within cluster.
func NewConfigFlagsFromCluster(namespace string, clusterConfig *rest.Config) *genericclioptions.ConfigFlags {
	impersonateGroup := []string{}
	insecure := false

	// CertFile and KeyFile must be nil for the BearerToken to be used for authentication and authorization instead of the pod's service account.
	return &genericclioptions.ConfigFlags{
		Insecure:         &insecure,
		Timeout:          stringptr("0"),
		Namespace:        stringptr(namespace),
		APIServer:        stringptr(clusterConfig.Host),
		CAFile:           stringptr(clusterConfig.CAFile),
		BearerToken:      stringptr(clusterConfig.BearerToken),
		ImpersonateGroup: &impersonateGroup,
	}
}

func stringptr(val string) *string {
	return &val
}
