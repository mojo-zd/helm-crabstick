package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// NewConfig ...
func NewConfig() *Config {
	conf := &Config{}
	bindAttrs(conf)
	return conf
}
func (c *Config) ConfigFlags(namespace string) *genericclioptions.ConfigFlags {
	return &genericclioptions.ConfigFlags{
		Namespace:   &namespace,
		KubeConfig:  &c.KubeConf,
		BearerToken: &c.KubeToken,
		Context:     &c.KubeContext,
	}
}

func bindAttrs(conf *Config) {
	viper.AutomaticEnv()
	// ./conf/config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
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
