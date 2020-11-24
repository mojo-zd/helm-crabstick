package chart

import (
	"net/http"

	"github.com/mojo-zd/helm-crabstick/api/server/router"
)

type chartRouter struct {
	routes []router.Route
}

func NewRouter() router.Router {
	return &chartRouter{}
}

func (c *chartRouter) Routes() []router.Route {
	return []router.Route{
		router.NewRoute(http.MethodGet, "/charts", c.getCharts),
	}
}
