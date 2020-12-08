package chart

import (
	"net/http"

	"github.com/mojo-zd/helm-crabstick/api/server/router"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
)

type chartRouter struct {
	routes []router.Route
	cfg    config.Config
}

func NewRouter(cfg config.Config) router.Router {
	return &chartRouter{cfg: cfg}
}

func (c *chartRouter) Routes() []router.Route {
	return []router.Route{
		router.NewRoute(http.MethodGet, "/charts", c.charts),
		router.NewRoute(http.MethodGet, "/charts/{name}", c.show),
		router.NewRoute(http.MethodGet, "/charts/{name}/versions", c.versions),
		router.NewRoute(http.MethodGet, "/charts/category", c.category),
	}
}
