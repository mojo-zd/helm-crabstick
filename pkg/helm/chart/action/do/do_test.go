package do

import (
	"testing"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
	"k8s.io/client-go/kubernetes"
)

var conf = config.Config{
	Repository: &config.Repository{
		Name: "bitnami",
		URL:  "https://charts.bitnami.com/bitnami",
	},
	KubeConf: "/Users/mojo/.kube/config",
	CacheDir: "/Users/mojo/.cache/helm",
}

func TestInstall(t *testing.T) {
	setting := util.NewSetting(conf)
	config, _ := setting.RESTClientGetter().ToRESTConfig()
	cli, err := kubernetes.NewForConfig(config)
	doer := NewDoer(conf, cli)
	rls, err := doer.Install("apache", "bn", "default", DoerOptions{Annotation: map[string]string{"author": "mojo"}})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rls.Name, rls.Chart.Metadata.Annotations["author"])
}
