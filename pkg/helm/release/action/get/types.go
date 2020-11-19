package get

import helmtime "helm.sh/helm/v3/pkg/time"

const (
	MaxHistory = 256
)

type releaseInfo struct {
	Revision    int           `json:"revision"`
	Updated     helmtime.Time `json:"updated"`
	Status      string        `json:"status"`
	Chart       string        `json:"chart"`
	AppVersion  string        `json:"app_version"`
	Description string        `json:"description"`
}

type ReleaseHistory []releaseInfo
