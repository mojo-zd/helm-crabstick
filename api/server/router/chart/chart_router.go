package chart

import (
	"github.com/kataras/iris/v12/context"
	"github.com/sirupsen/logrus"
)

func (c *chartRouter) requestID(ctx context.Context) string {
	return ctx.Params().Get("id")
}

func (c *chartRouter) getCharts(ctx context.Context) error {
	ctx.JSON(map[string]interface{}{"name": "chartDemo", "version": "v0.12.0"})
	logrus.Info("call get charts api")
	return nil
}
