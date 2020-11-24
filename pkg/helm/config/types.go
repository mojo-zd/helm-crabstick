package config

// Repository helm repository config
type Repository struct {
	Name                  string
	URL                   string
	Username              string
	Password              string
	CaFile                string
	CertFile              string
	KeyFile               string
	InsecureSkipTlsVerify bool
	Cache                 string // helm cache directory e.g. /root/.cache/helm
	Config                string // helm repository config directory e.g. /root/.config/helm
	MaxHistory            int
	StorageBackend        string // secrets、configmap
}

// Config ...
type Config struct {
	Repository  *Repository
	KubeConf    string
	KubeToken   string // option
	KubeContext string // option
	LogLevel    string
}
