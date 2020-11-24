package server

import "github.com/kataras/iris/v12/context"

func (s *Server) handlerWithGlobalMiddleware(ctx context.Context) error {
	for _, m := range s.middlewares {
		if err := m.WrapHandler(ctx); err != nil {
			return err
		}
	}
	return nil
}
