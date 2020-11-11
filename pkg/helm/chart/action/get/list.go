package get

import (
	"fmt"
	"sort"

	"helm.sh/helm/v3/pkg/repo"
)

// List get repository's all chart
func (g *getter) List() ChartVersions {
	charts := ChartVersions{}
	chartFilter := make(map[string]*repo.ChartVersion)
	file, err := g.cache.LoadIndex()
	if err != nil {
		return charts
	}

	// remove repeat and get the latest chart
	for _, chartVersion := range file.Entries {
		for _, chart := range chartVersion {
			if _, ok := chartFilter[chart.Name]; !ok {
				chartFilter[chart.Name] = chart
				charts = append(charts, &ChartVersion{
					Name:        fmt.Sprintf("%s/%s", g.repoName, chart.Name),
					AppVersion:  chart.AppVersion,
					Version:     chart.Version,
					Description: chart.Description,
					Icon:        chart.Icon,
				})
			}
		}
	}
	sort.Sort(charts)
	return charts
}
