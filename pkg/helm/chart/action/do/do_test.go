package do

import (
	"fmt"
	"os"
	"testing"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
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
	chartName = "apache"
	namespace = "default"
)

func TestInstall(t *testing.T) {
	setting := util.NewSetting(conf)
	config, _ := setting.RESTClientGetter().ToRESTConfig()
	cli, err := kubernetes.NewForConfig(config)
	doer := NewDoer(conf, cli)
	rls, err := doer.Install(chartName, "bn", namespace, DoerOptions{Annotation: map[string]string{"author": "mojo"}})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rls.Name, rls.Chart.Metadata.Annotations["author"])
}
