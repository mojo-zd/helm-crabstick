package get

import (
	"fmt"
	"os"
	"testing"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"helm.sh/helm/v3/pkg/action"
)

var (
	home = os.Getenv("HOME")
	conf = config.Config{
		Repository: &config.Repository{
			Name:  "bitnami",
			URL:   "https://charts.bitnami.com/bitnami",
			Cache: fmt.Sprintf("%s/%s", home, ".cache/helm"),
		},
		KubeConf: fmt.Sprintf("%s/%s", home, ".kube/config"),
	}
	g         = NewGetter(conf)
	chartName = "apache"
)

func TestList(t *testing.T) {
	charts := g.List()
	for _, chart := range charts {
		t.Log(chart.Name, chart.Version, chart.AppVersion, chart.Icon, chart.Description)
	}
}

func TestChartVersions(t *testing.T) {
	info := g.ChartVersion(chartName)
	for _, version := range info.Versions {
		t.Log(info.Name, version.Ver, version.Description)
	}
}

func TestChartShow(t *testing.T) {
	out := g.Show(chartName, "", action.ShowAll)
	t.Log(out)
}
