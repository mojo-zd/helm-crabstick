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
		KubeConf: fmt.Sprintf("%s/%s", home, ".kube/config"),
		CacheDir: fmt.Sprintf("%s/%s", home, ".cache/helm"),
	}
	chartName   = "apache"
	namespace   = "demo"
	releaseName = "bn"
)

func TestInstall(t *testing.T) {
	doer := NewDoer(conf, getClient(t))
	rls, err := doer.Install(chartName, releaseName, namespace, DoerOptions{Annotation: map[string]string{"author": "mojo"}})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rls.Name, rls.Chart.Metadata.Annotations["author"])
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
