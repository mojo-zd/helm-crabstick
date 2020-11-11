package get

import (
	"sort"

	"github.com/sirupsen/logrus"
)

// ChartVersion list versions of chart
func (g *getter) ChartVersion(chartName string) ChartInfo {
	info := ChartInfo{Name: chartName}
	versions := Versions{}
	file, err := g.cache.LoadIndex()
	if err != nil {
		logrus.Errorf("load repository[%s] index failed, err:%s", g.repoName, err.Error())
		return info
	}

	for _, chartVersions := range file.Entries {
		for _, version := range chartVersions {
			if version.Name == chartName {
				versions = append(versions, &Version{
					Ver:         version.Version,
					Description: version.Description,
				})
			}
		}
	}
	sort.Sort(versions)
	info.Versions = versions
	return info
}
