package get

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/cache"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"helm.sh/helm/v3/pkg/action"
)

const CategoryKey = "category"

type Getter interface {
	// List get repository's all chart
	List(category string) ChartVersions
	// Charts find chart information with chart name from repository
	// it will find the spec chart if version assigned
	ChartVersion(chartName string) ChartInfo

	// Show show chart detail information
	Show(name, version string, output action.ShowOutputFormat) string

	// Versions list all version of chart
	Versions(name string) ChartVersions

	// Catalog get all catalog of chart
	Category() Category
}

type ChartVersions []*ChartVersion
type Versions []*Version

type ChartVersion struct {
	Name        string
	Version     string
	AppVersion  string
	Description string
	Icon        string
}

type ChartInfo struct {
	Name     string
	Versions Versions
}

type Version struct {
	Ver         string
	Description string
}

type getter struct {
	cache    cache.IndexCache
	repoName string
	cfg      config.Config
}

func NewGetter(conf config.Config) Getter {
	return &getter{
		cache:    cache.NewIndexCache(conf.Repository.Name),
		repoName: conf.Repository.Name,
		cfg:      conf,
	}
}

func (charts ChartVersions) Len() int {
	return len(charts)
}

func (charts ChartVersions) Less(i, j int) bool {
	return charts[i].Name < charts[j].Name
}

func (charts ChartVersions) Swap(i, j int) {
	charts[i], charts[j] = charts[j], charts[i]
}

func (version Versions) Len() int {
	return len(version)
}

func (version Versions) Less(i, j int) bool {
	return version[i].Ver > version[j].Ver
}

func (version Versions) Swap(i, j int) {
	version[i], version[j] = version[j], version[i]
}
