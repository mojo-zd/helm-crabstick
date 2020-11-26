package release

import (
	"github.com/kataras/iris/v12/context"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/manager"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/parser/release"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (r *releaseRouter) requestID(ctx context.Context) string {
	return ctx.Params().Get("id")
}

func (r *releaseRouter) releases(ctx context.Context) error {
	releases, err := manager.NewAppManager(r.cfg).ReleaseGetter.List("", util.ListOptions{})
	if err != nil {
		return err
	}
	_, err = ctx.JSON(release.ToReleases(releases))
	return err
}

func (r *releaseRouter) release(ctx context.Context) error {
	name := ctx.URLParam("name")
	namespace := ctx.URLParam("namespace")
	mgr := manager.NewAppManager(r.cfg)
	rls, err := mgr.ReleaseGetter.Get(name, namespace)
	if err != nil {
		return err
	}

	resources := mgr.ReleaseGetter.Resources(name, namespace, rls, v1.ListOptions{})
	_, err = ctx.JSON(release.Profound{Release: rls, Resource: resources})
	return err
}
