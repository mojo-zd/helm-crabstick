package util

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"helm.sh/helm/v3/pkg/cli"
)

// NewSetting return helm cli setting
func NewSetting(config config.Config) *cli.EnvSettings {
	debug := false
	if config.LogLevel == "debug" {
		debug = true
	}

	setting := cli.New()
	setting.KubeConfig = config.KubeConf
	setting.KubeContext = config.KubeContext
	setting.KubeToken = config.KubeToken
	setting.RepositoryConfig = config.ConfigDir
	setting.RepositoryCache = config.CacheDir
	setting.MaxHistory = config.MaxHistory
	setting.Debug = debug

	return setting
}
