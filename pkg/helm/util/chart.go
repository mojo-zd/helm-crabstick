package util

import (
	"errors"

	"github.com/sirupsen/logrus"

	"sigs.k8s.io/yaml"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"helm.sh/helm/v3/pkg/action"
)

// LoadChartOptions new ChartPathOptions
func LoadChartOptions(config config.Config) (action.ChartPathOptions, error) {
	if config.Repository == nil {
		return action.ChartPathOptions{}, errors.New("not found repository information")
	}

	return action.ChartPathOptions{
		CaFile:                config.Repository.CaFile,
		CertFile:              config.Repository.CertFile,
		KeyFile:               config.Repository.KeyFile,
		InsecureSkipTLSverify: config.Repository.InsecureSkipTlsVerify,
		Username:              config.Repository.Username,
		Password:              config.Repository.Password,
		RepoURL:               config.Repository.URL,
	}, nil
}

func GetValues(value string) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	err := yaml.Unmarshal([]byte(value), &out)
	if err != nil {
		logrus.Errorf("unmarshal values failed, err:%s", err.Error())
		return out, err
	}
	return out, err
}
