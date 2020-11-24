package util

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/cli"
)

// NewSetting return helm cli setting
func NewSetting(config config.Config) *cli.EnvSettings {
	debug := false
	if config.LogLevel == "debug" {
		debug = true
	}

	setting := cli.New()
	if config.Repository == nil {
		logrus.Error("not found repository config")
		return setting
	}

	setting.KubeConfig = config.KubeConf
	setting.KubeContext = config.KubeContext
	setting.KubeToken = config.KubeToken
	setting.RepositoryCache = config.Repository.Cache
	setting.RepositoryConfig = config.Repository.Config
	setting.MaxHistory = config.Repository.MaxHistory
	setting.Debug = debug

	return setting
}
