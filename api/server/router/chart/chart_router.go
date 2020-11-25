package chart

import (
	"github.com/kataras/iris/v12/context"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/manager"
	"helm.sh/helm/v3/pkg/action"
)

func (c *chartRouter) requestID(ctx context.Context) string {
	return ctx.Params().Get("id")
}

func (c *chartRouter) charts(ctx context.Context) error {
	category := ctx.URLParam("category")
	_, err := ctx.JSON(manager.NewAppManager(c.cfg).ChartGetter.List(category))
	return err
}

func (c *chartRouter) versions(ctx context.Context) error {
	chart := ctx.Params().Get("name")
	_, err := ctx.JSON(manager.NewAppManager(c.cfg).ChartGetter.Versions(chart))
	return err
}

func (c *chartRouter) category(ctx context.Context) error {
	_, err := ctx.JSON(manager.NewAppManager(c.cfg).ChartGetter.Category())
	return err
}

func (c *chartRouter) show(ctx context.Context) error {
	name := ctx.Params().Get("name")
	_, err := ctx.JSON(map[string]interface{}{
		"result": manager.NewAppManager(c.cfg).
			ChartGetter.
			Show(name, "", action.ShowAll),
	})
	return err
}
