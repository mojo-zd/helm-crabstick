package chart

import (
	"github.com/kataras/iris/v12/context"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/manager"
)

func (c *chartRouter) requestID(ctx context.Context) string {
	return ctx.Params().Get("id")
}

func (c *chartRouter) getCharts(ctx context.Context) error {
	ctx.JSON(manager.NewAppManager(c.cfg).ChartGetter.List())
	return nil
}
