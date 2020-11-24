package release

import (
	"net/http"

	"github.com/mojo-zd/helm-crabstick/api/server/router"
)

type releaseRouter struct {
	routes []router.Route
}

func NewRouter() router.Router {
	return &releaseRouter{}
}

func (r *releaseRouter) Routes() []router.Route {
	return []router.Route{
		router.NewRoute(http.MethodGet, "/releases", r.getReleases),
	}
}
