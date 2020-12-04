package release

import (
	"net/http"

	"github.com/mojo-zd/helm-crabstick/pkg/manager"

	"github.com/mojo-zd/helm-crabstick/api/server/router"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
)

type releaseRouter struct {
	routes     []router.Route
	cfg        config.Config
	clusterMgr manager.Manager
}

func NewRouter(cfg config.Config, clusterMgr manager.Manager) router.Router {
	return &releaseRouter{cfg: cfg, clusterMgr: clusterMgr}
}

func (r *releaseRouter) Routes() []router.Route {
	return []router.Route{
		router.NewRoute(http.MethodGet, "/clusters/{cluster_uuid}/releases", r.releases),
		router.NewRoute(http.MethodGet, "/clusters/{cluster_uuid}/releases/{name}", r.release),
		router.NewRoute(http.MethodPost, "/clusters/{cluster_uuid}/releases", r.install),
		router.NewRoute(http.MethodPut, "/clusters/{cluster_uuid}/releases/{name}", r.upgrade),
		router.NewRoute(http.MethodDelete, "/clusters/{cluster_uuid}/releases/{name}", r.uninstall),
		router.NewRoute(http.MethodGet, "/clusters/{cluster_uuid}/releases/{name}/history", r.history),
	}
}
