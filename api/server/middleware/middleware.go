package middleware

import "github.com/kataras/iris/v12/context"

// Middleware is an interface to allow the use of ordinary functions as api
// Any struct that has the appropriate signature can be registered as a middleware.
type Middleware interface {
	WrapHandler(ctx context.Context) error
}
