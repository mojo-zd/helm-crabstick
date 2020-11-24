package middleware

import (
	"github.com/kataras/iris/v12/context"
	"github.com/sirupsen/logrus"
)

type LogRequestMiddleware struct {
}

func NewLogRequestMiddleware() Middleware {
	return &LogRequestMiddleware{}
}

func (middleware *LogRequestMiddleware) WrapHandler(ctx context.Context) error {
	logrus.Infof("request path:[%s], method:[%s] param[%s]", ctx.Path(), ctx.Method(), ctx.URLParams())
	return nil
}
