package release

import (
	"github.com/kataras/iris/v12/context"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/manager"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/parser/release"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
)

func (r *releaseRouter) requestID(ctx context.Context) string {
	return ctx.Params().Get("id")
}

func (r *releaseRouter) getReleases(ctx context.Context) error {
	releases, err := manager.NewAppManager(r.cfg).ReleaseGetter.List("", util.ListOptions{})
	if err != nil {
		return err
	}
	_, err = ctx.JSON(release.ToReleases(releases))
	return err
}
