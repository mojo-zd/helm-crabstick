package config

// Repository helm repository config
type Repository struct {
	Name                  string `mapstructure:"name"`
	URL                   string `mapstructure:"url"`
	Username              string `mapstructure:"username"`
	Password              string `mapstructure:"password"`
	CaFile                string `mapstructure:"caFile"`
	CertFile              string `mapstructure:"certFile"`
	KeyFile               string `mapstructure:"keyFile"`
	InsecureSkipTlsVerify bool   `mapstructure:"insecureSkipTlsVerify"`
}

// Config ...
type Config struct {
	Repository     *Repository `mapstructure:"repository"`
	Harbor         *Harbor     `mapstructure:"harbor"`
	KubeConf       string      `mapstructure:"kubeConf"`
	KubeToken      string      `mapstructure:"kubeToken"`   // option
	KubeContext    string      `mapstructure:"kubeContext"` // option
	CacheDir       string      `mapstructure:"cacheDir"`    // helm cache directory e.g. /root/.cache/helm
	ConfigDir      string      `mapstructure:"configDir"`   // helm config directory e.g. /root/.config/helm
	LogLevel       string      `mapstructure:"logLevel"`
	MaxHistory     int         `mapstructure:"maxHistory"`
	StorageBackend string      `mapstructure:"storageBackend"` // secrets、configmap
}

// Harbor information of harbor
type Harbor struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	URL      string `mapstructure:"url"`
}
