package get

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
		KubeConf: fmt.Sprintf("%s/.kube/config", home),
		CacheDir: fmt.Sprintf("%s/.cache/helm", home),
	}
	namespace   = "demo"
	releaseName = "bn"
)

func TestReleaseList(t *testing.T) {
	restConf, err := conf.ConfigFlags().ToRESTConfig()
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
	release, err := NewGetter(conf, client, nil).Get(releaseName, namespace)
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
	getter := NewGetter(conf, client, nil)
	rels, err := getter.List(namespace, util.ListOptions{})
	if err != nil {
		t.Error(err)
		return
	}

	for _, rel := range rels {
		t.Logf("release [%s], kind: %s", rel.Name, getter.Kind(rel.Name, namespace))
	}
}

func TestHistory(t *testing.T) {
	client, err := getConfAndClient()
	if err != nil {
		t.Fatal(err)
	}
	getter := NewGetter(conf, client, nil)
	out, err := getter.History(releaseName, namespace)
	if err != nil {
		t.Error(err)
		return
	}

	for _, re := range out {
		t.Log(re.Chart, re.Description, re.AppVersion, re.Status, re.Revision, re.Updated)
	}
}

func getConfAndClient() (kubernetes.Interface, error) {
	restConf, _ := conf.ConfigFlags().ToRESTConfig()
	client, err := kubernetes.NewForConfig(restConf)
	if err != nil {
		return nil, err
	}

	return client, nil
}
