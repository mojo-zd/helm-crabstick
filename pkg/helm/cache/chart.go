package cache

import "helm.sh/helm/v3/pkg/chart/loader"

type ChartCache interface {
	Manifest(name string)
}

type chartCache struct {
	repoName string
}

func NewChartCache(repoName string) ChartCache {
	return &chartCache{repoName: repoName}
}

func (cache *chartCache) Manifest(name string) {

}

func (cache *chartCache) loadChart(name string) {
	loader.Load()
}

func (cache *chartCache) getChartURL(name string) string {
	idxCache := NewIndexCache(cache.repoName)
	file, err := idxCache.LoadIndex()
	if err != nil {

	}
}
