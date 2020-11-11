package get

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/cache"
)

type Getter interface {
	// List get repository's all chart
	List() ChartVersions
	// Charts find chart information with chart name from repository
	// it will find the spec chart if version assigned
	ChartVersion(chartName string) ChartInfo
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
}

func NewGetter(repo string) Getter {
	return &getter{
		cache:    cache.NewIndexCache(repo),
		repoName: repo,
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
