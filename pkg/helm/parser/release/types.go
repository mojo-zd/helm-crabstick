package release

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"

	"helm.sh/helm/v3/pkg/release"
)

// Release short info of release
type Release struct {
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
	Revision   int    `json:"revision"`
	UpdateAt   string `json:"updateAt"`
	Status     string `json:"status"`
	Chart      string `json:"chart"`
	AppVersion string `json:"appVersion"`
}

// Profound more detail information of release
type Profound struct {
	Release
}

func ToReleases(releases []*release.Release) []*Release {
	rls := []*Release{}
	for _, release := range releases {
		rls = append(rls, CopyToRelease(release))
	}
	return rls
}

func CopyToRelease(release *release.Release) *Release {
	var updated string
	if last := release.Info.LastDeployed; !last.IsZero() {
		updated = last.Format(util.FormatStripingYMDHMS)
	}

	return &Release{
		Name:       release.Name,
		Namespace:  release.Namespace,
		Revision:   release.Version,
		UpdateAt:   updated,
		Status:     string(release.Info.Status),
		Chart:      release.Chart.Name(),
		AppVersion: release.Chart.AppVersion(),
	}
}

func CopyToProfound(release *release.Release) {

}
