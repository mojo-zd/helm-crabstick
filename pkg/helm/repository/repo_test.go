package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
)

var (
	home = os.Getenv("HOME")
	conf = config.Config{
		Repository: &config.Repository{
			Name:  "bitnami",
			URL:   "https://charts.bitnami.com/bitnami",
			Cache: fmt.Sprintf("%s/.cache/helm", home),
		},
		KubeConf: fmt.Sprintf("%s/.kube/config", home),
	}
)

func TestRepo(t *testing.T) {
	repo := NewRepo(conf)
	if err := repo.CacheIndex(); err != nil {
		t.Error(err.Error())
		return
	}
}
