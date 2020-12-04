package config

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// 参考 https://github.com/spf13/viper#unmarshaling
type Config struct {
	Address       string `mapstructure:"address"` //address for server
	SocketAddress string `mapstructure:"socketAddress"`
	TokenExpired  int    `mapstructure:"expired"`   //token for expired
	SecretKey     string `mapstructure:"secretkey"` // secret for token

	Middlewares []string `mapstructure:"middlewares"` //middleware

	Redis struct { //for redis
		Address  string `mapstructure:"address"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
	} `mapstructure:"redis"`

	Registry struct {
		Address string `mapstructure:"host"`
	} `mapstructure:"registry"` // registry center

	Repository struct {
		Name           string `mapstructure:"name"`
		URL            string `mapstructure:"url"`
		Username       string `mapstructure:"username"`
		Password       string `mapstructure:"password"`
		MaxHistory     string `mapstructure:"maxHistory"`
		StorageBackend string `mapstructure:"storageBackend"`
	}
	Auth struct {
		URL string `mapstructure:"url"`
	}
	RunMode string `mapstructure:"runMode"` // dev or prod
}

var (
	cfg *Config
)

func init() {
	viper.AutomaticEnv() // 绑定环境变量
	// 环境变量分隔符使用"_"
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	// ./conf/config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}

	cfg = &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		logrus.Fatal("config.yaml invalid", err)
	}

	if cfg.RunMode == "dev" {
		logrus.SetReportCaller(true)
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.Infof("init application with config: %+v", cfg)
}

// GetConfig ...
func GetConfig() *Config {
	return cfg
}
