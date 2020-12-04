package get

import (
	"fmt"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/storage"
	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/releaseutil"
)

func (g *getter) History(name, namespace string) (ReleaseHistory, error) {
	configuration := storage.ActionConfiguration(*g.cluster, g.config, namespace)
	client := action.NewHistory(configuration)
	client.Max = MaxHistory
	return g.history(client, name)
}

func (g *getter) history(client *action.History, name string) (ReleaseHistory, error) {
	var releaseHistory ReleaseHistory
	hist, err := client.Run(name)
	if err != nil {
		return releaseHistory, err
	}
	releaseutil.Reverse(hist, releaseutil.SortByRevision)

	var rels []*release.Release
	for i := 0; i < min(len(hist), client.Max); i++ {
		rels = append(rels, hist[i])
	}

	if len(rels) == 0 {
		logrus.Warnf("not found release[%s] history record", name)
		return releaseHistory, nil
	}
	return getReleaseHistory(rels), nil
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getReleaseHistory(rls []*release.Release) (history ReleaseHistory) {
	for i := len(rls) - 1; i >= 0; i-- {
		r := rls[i]
		c := formatChartname(r.Chart)
		s := r.Info.Status.String()
		v := r.Version
		d := r.Info.Description
		a := formatAppVersion(r.Chart)

		rInfo := releaseInfo{
			Revision:    v,
			Status:      s,
			Chart:       c,
			AppVersion:  a,
			Description: d,
		}
		if !r.Info.LastDeployed.IsZero() {
			rInfo.Updated = r.Info.LastDeployed

		}
		history = append(history, rInfo)
	}

	return history
}

func formatChartname(c *chart.Chart) string {
	if c == nil || c.Metadata == nil {
		// This is an edge case that has happened in prod, though we don't
		// know how: https://github.com/helm/helm/issues/1347
		return "MISSING"
	}
	return fmt.Sprintf("%s-%s", c.Name(), c.Metadata.Version)
}

func formatAppVersion(c *chart.Chart) string {
	if c == nil || c.Metadata == nil {
		// This is an edge case that has happened in prod, though we don't
		// know how: https://github.com/helm/helm/issues/1347
		return "MISSING"
	}
	return c.AppVersion()
}
