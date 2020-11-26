package get

func (g *getter) Versions(name string) ChartVersions {
	charts := ChartVersions{}
	file, err := g.cache.LoadIndex()
	if err != nil {
		return charts
	}
	for _, chartVersion := range file.Entries {
		for _, chart := range chartVersion {
			if chart.Name == name {
				charts = append(charts, &ChartVersion{
					Name:        chart.Name,
					Version:     chart.Version,
					AppVersion:  chart.AppVersion,
					Description: chart.Description,
					Icon:        chart.Icon,
					Home:        chart.Home,
					Maintainers: chart.Maintainers,
					Sources:     chart.Sources,
				})
			}
		}
	}
	return charts
}
