package release

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
	"helm.sh/helm/v3/pkg/release"
)

// Release short info of release
type Release struct {
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	Revision     int    `json:"revision"`
	UpdateAt     string `json:"updateAt"`
	Status       string `json:"status"`
	Chart        string `json:"chart"`
	AppVersion   string `json:"appVersion"`
	ChartVersion string `json:"chartVersion"`
	Description  string `json:"description"`
}

// Profound include release and kuberentes resources
type Profound struct {
	Release  *release.Release `json:"release"`
	Resource interface{}      `json:"resource"` // kubernetes resources
}

// ToReleases convert release.Release to brief
func ToReleases(releases []*release.Release) []*Release {
	rls := []*Release{}
	for _, r := range releases {
		rls = append(rls, CopyToRelease(r))
	}
	return rls
}

// CopyToRelease copy release.Relase attr to Release
func CopyToRelease(release *release.Release) *Release {
	var updated string
	if last := release.Info.LastDeployed; !last.IsZero() {
		updated = last.Format(util.FormatStripingYMDHMS)
	}

	return &Release{
		Name:         release.Name,
		Namespace:    release.Namespace,
		Revision:     release.Version,
		UpdateAt:     updated,
		Status:       string(release.Info.Status),
		Chart:        release.Chart.Name(),
		AppVersion:   release.Chart.AppVersion(),
		ChartVersion: release.Chart.Metadata.Version,
		Description:  release.Info.Description,
	}
}
