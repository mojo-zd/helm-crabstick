package httputils

import "github.com/kataras/iris/v12/context"

type APIFunc func(ctx context.Context) error
