package get

import (
	"testing"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
	"k8s.io/client-go/kubernetes"
)

var (
	conf = config.Config{
		Repository: &config.Repository{
			Name: "bitnami",
			URL:  "https://charts.bitnami.com/bitnami",
		},
		KubeConf: "/Users/mojo/.kube/config",
		CacheDir: "/Users/mojo/.cache/helm",
	}
	namespace = "default"
)

func TestReleaseList(t *testing.T) {
	restConf, err := conf.ConfigFlags(namespace).ToRESTConfig()
	if err != nil {
		t.Fatal(err)
		return
	}
	client, err := kubernetes.NewForConfig(restConf)
	if err != nil {
		t.Fatal(err)
	}
	getter := NewGetter(conf, client)
	releases, err := getter.List(namespace, util.ListOptions{Annotation: map[string]string{"author": "mojo"}})
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, release := range releases {
		t.Log(release.Name, release.Chart.Metadata.Annotations)
	}
}
