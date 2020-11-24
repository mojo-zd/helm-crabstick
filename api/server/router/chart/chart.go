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
		router.NewRoute(http.MethodGet, "/charts", c.getCharts),
	}
}
