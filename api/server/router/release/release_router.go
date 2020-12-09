package release

import (
	"github.com/kataras/iris/v12/context"
	"github.com/mojo-zd/helm-crabstick/pkg/auth"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/manager"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/parser"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/types"
	mg "github.com/mojo-zd/helm-crabstick/pkg/manager"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	defNamespace = "default"
)

func (r *releaseRouter) requestID(ctx context.Context) string {
	return ctx.Params().Get("id")
}

func (r *releaseRouter) releases(ctx context.Context) error {
	cluster, err := r.getCluster(ctx)
	if err != nil {
		return err
	}

	token, err := r.token(ctx)
	if err != nil {
		return err
	}

	releases, err := r.releaseService.List(r.cfg, token, cluster)
	if err != nil {
		return err
	}
	_, err = ctx.JSON(releases)
	return err
}

func (r *releaseRouter) release(ctx context.Context) error {
	name := ctx.Params().Get("name")
	namespace := ctx.URLParamDefault("namespace", defNamespace)
	cluster, err := r.getCluster(ctx)
	if err != nil {
		return err
	}
	mgr := manager.NewAppManager(r.cfg, &cluster)
	rls, err := mgr.ReleaseGetter.Get(name, namespace)
	if err != nil {
		return err
	}

	resources := mgr.ReleaseGetter.Resources(name, namespace, rls, v1.ListOptions{})
	_, err = ctx.JSON(parser.Profound{Release: rls, Resource: resources})
	return err
}

func (r *releaseRouter) install(ctx context.Context) error {
	cluster, err := r.getCluster(ctx)
	if err != nil {
		return err
	}
	token, err := r.token(ctx)
	if err != nil {
		return err
	}
	createOpts := types.CreateOptions{}
	if err := ctx.ReadJSON(&createOpts); err != nil {
		return err
	}

	rls, err := r.releaseService.Create(r.cfg, cluster, token, createOpts)
	if err != nil {
		return err
	}
	_, err = ctx.JSON(rls)
	return err
}

func (r *releaseRouter) uninstall(ctx context.Context) error {
	cluster, err := r.getCluster(ctx)
	if err != nil {
		return err
	}
	name := ctx.Params().Get("name")
	namespace := ctx.URLParamDefault("namespace", defNamespace)
	err = r.releaseService.Delete(r.cfg, cluster, name, namespace)
	if err != nil {
		return err
	}
	_, err = ctx.JSON(map[string]interface{}{})
	return err
}

func (r *releaseRouter) upgrade(ctx context.Context) error {
	cli, err := r.getCluster(ctx)
	if err != nil {
		return err
	}
	opts := types.UpgradeOptions{}
	err = ctx.ReadJSON(&opts)
	if err != nil {
		return err
	}
	rls, err := manager.NewAppManager(r.cfg, &cli).ReleaseDoer.Upgrade(opts)
	if err != nil {
		return err
	}
	_, err = ctx.JSON(rls)
	return err
}

func (r *releaseRouter) history(ctx context.Context) error {
	cluster, err := r.getCluster(ctx)
	if err != nil {
		return err
	}
	name := ctx.Params().Get("name")
	namespace := ctx.URLParamDefault("namespace", defNamespace)
	history, err := manager.NewAppManager(r.cfg, &cluster).ReleaseGetter.History(name, namespace)
	if err != nil {
		return err
	}
	_, err = ctx.JSON(history)
	return err
}

func (r *releaseRouter) getCluster(ctx context.Context) (mg.Cluster, error) {
	cluster := ctx.Params().Get("cluster_uuid")
	token := ctx.GetHeader("TOKEN")
	return r.clusterMgr.Client(cluster, token)
}

func (r *releaseRouter) token(ctx context.Context) (auth.Token, error) {
	return r.clusterMgr.Token(ctx.GetHeader("TOKEN"))
}
