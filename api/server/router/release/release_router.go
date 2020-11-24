package release

import "github.com/kataras/iris/v12/context"

func (r *releaseRouter) requestID(ctx context.Context) string {
	return ctx.Params().Get("id")
}

func (r *releaseRouter) getReleases(ctx context.Context) error {
	return nil
}
