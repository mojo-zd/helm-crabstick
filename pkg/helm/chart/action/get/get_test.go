package get

import "testing"

var g = NewGetter("bitmina")

func TestList(t *testing.T) {
	charts := g.List()
	for _, chart := range charts {
		t.Log(chart.Name, chart.Version, chart.AppVersion, chart.Icon, chart.Description)
	}
}

func TestChartVersions(t *testing.T) {
	info := g.ChartVersion("apache")
	for _, version := range info.Versions {
		t.Log(info.Name, version.Ver, version.Description)
	}
}
