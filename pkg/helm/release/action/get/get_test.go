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
	namespace = "aaa"
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
	getter := NewGetter(conf, client, nil)
	releases, err := getter.List(namespace, util.ListOptions{Annotation: map[string]string{"author": "mojo"}})
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, release := range releases {
		t.Log(release.Name, release.Chart.Metadata.Annotations)
	}
}

func TestReleaseGet(t *testing.T) {
	client, err := getConfAndClient()
	if err != nil {
		t.Fatal(err)
	}
	release, err := NewGetter(conf, client, nil).Get("mn", namespace)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(release.Manifest)
}

func TestReleaseKind(t *testing.T) {
	client, err := getConfAndClient()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(NewGetter(conf, client, nil).Kind("mn", namespace))
}

func getConfAndClient() (kubernetes.Interface, error) {
	restConf, _ := conf.ConfigFlags("aaa").ToRESTConfig()
	client, err := kubernetes.NewForConfig(restConf)
	if err != nil {
		return nil, err
	}

	return client, nil
}
