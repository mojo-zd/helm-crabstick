package chart

import (
	"github.com/kataras/iris/v12/context"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/manager"
	"github.com/mojo-zd/helm-crabstick/pkg/util/page"
	"helm.sh/helm/v3/pkg/action"
)

func (c *chartRouter) requestID(ctx context.Context) string {
	return ctx.Params().Get("id")
}

func (c *chartRouter) charts(ctx context.Context) error {
	category := ctx.URLParam("category")
	current := ctx.URLParamInt64Default("current", page.DefCurrent)
	pageSize := ctx.URLParamInt64Default("pageSize", page.DefSize)

	charts := manager.NewChartManager(c.cfg).ChartGetter.List(category)
	_, err := ctx.JSON(page.NewPagination(charts, pageSize, current))
	return err
}

func (c *chartRouter) versions(ctx context.Context) error {
	chart := ctx.Params().GetDecoded("name")
	_, err := ctx.JSON(manager.NewChartManager(c.cfg).ChartGetter.Versions(chart))
	return err
}

func (c *chartRouter) category(ctx context.Context) error {
	_, err := ctx.JSON(manager.NewChartManager(c.cfg).ChartGetter.Category())
	return err
}

func (c *chartRouter) show(ctx context.Context) error {
	name := ctx.Params().GetDecoded("name")
	_, err := ctx.JSON(manager.NewChartManager(c.cfg).
		ChartGetter.Show(name, "", action.ShowReadme))
	return err
}
