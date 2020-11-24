package config

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// NewConfig ...
func NewConfig(path string) (*Config, error) {
	conf := &Config{}
	binding(conf, path)
	err := conf.validate()
	return conf, err
}

// ConfigFlags ...
func (c *Config) ConfigFlags() *genericclioptions.ConfigFlags {
	return &genericclioptions.ConfigFlags{
		KubeConfig:  &c.KubeConf,
		BearerToken: &c.KubeToken,
		Context:     &c.KubeContext,
	}
}

func (c *Config) validate() error {
	if len(c.KubeConf) == 0 {
		return errors.New("please spec the kube config. e.g. $HOME/.kube/config")
	}
	if c.Repository == nil {
		return errors.New("not setting helm repository")
	}
	// set default storage backend
	if len(c.Repository.StorageBackend) == 0 {
		c.Repository.StorageBackend = "secrets"
	}
	return nil
}

func binding(conf *Config, path string) {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// set default path ./conf/config.yaml
	viper.AddConfigPath("conf")
	if len(path) != 0 {
		viper.AddConfigPath(path)
	}

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}

	if err := viper.Unmarshal(conf); err != nil {
		logrus.Fatal(err)
	}

	if conf.LogLevel == "debug" {
		logrus.SetReportCaller(true)
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.Infof("init application with config: %+v", conf)
}
