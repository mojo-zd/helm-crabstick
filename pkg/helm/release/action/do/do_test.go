package do

import (
	"fmt"
	"os"
	"testing"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"k8s.io/client-go/kubernetes"
)

var (
	home = os.Getenv("HOME")
	conf = config.Config{
		Repository: &config.Repository{
			Name: "bitnami",
			URL:  "https://charts.bitnami.com/bitnami",
		},
		KubeConf: fmt.Sprintf("%s/.kube/config", home),
		CacheDir: fmt.Sprintf("%s/.cache/helm", home),
	}
	namespace   = "demo"
	releaseName = "bn"
	chart       = "apache"
	version     = "8.0.0"
)

func TestUpgrade(t *testing.T) {
	_, err := NewDoer(getClient(t), conf).Upgrade(releaseName, chart, version, namespace)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {
	doer := NewDoer(getClient(t), conf)
	if _, err := doer.Delete(releaseName, namespace); err != nil {
		t.Error("delete release failed", err)
		return
	}
	t.Log("delete success")
}

func getClient(t *testing.T) kubernetes.Interface {
	restConf, err := conf.ConfigFlags().ToRESTConfig()
	if err != nil {
		t.Fatal(err)
		return nil
	}
	client, err := kubernetes.NewForConfig(restConf)
	if err != nil {
		t.Fatal(err)
	}

	return client
}
