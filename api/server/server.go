package server

import (
	"strings"
	"time"

	uuid "github.com/iris-contrib/go.uuid"
	iris "github.com/kataras/iris/v12"
	context "github.com/kataras/iris/v12/context"
	"github.com/mojo-zd/helm-crabstick/api/server/httputils"
	"github.com/mojo-zd/helm-crabstick/api/server/middleware"
	"github.com/mojo-zd/helm-crabstick/api/server/router"
	"github.com/sirupsen/logrus"
)

// Config provides the configuration for the API server
type Config struct {
	Address      string
	Middleware   []string
	JWTSecret    string
	KeystoneAddr string
}

// Server contains instance details for the server
type Server struct {
	cfg         *Config
	routers     []router.Router
	app         *iris.Application
	middlewares []middleware.Middleware
}

//New return a new server
func New(cfg *Config) *Server {
	return &Server{
		app: iris.New(),
		cfg: cfg,
	}
}

func (s *Server) InitRouter(routers ...router.Router) {
	s.routers = append(s.routers, routers...)
	for _, m := range s.cfg.Middleware {
		switch m {
		case "jwt":
			s.UseMiddleware(middleware.NewJwtRequestMiddleware(s.cfg.JWTSecret))
		}
	}
	s.UseMiddleware(middleware.NewLogRequestMiddleware())
	s.createMux()
}

// UseMiddleware appends a new middleware to the request chain.
// This needs to be called before the API routes are configured.
func (s *Server) UseMiddleware(m middleware.Middleware) {
	s.middlewares = append(s.middlewares, m)
}

// Wait blocks the sermakeIrisHandlerver goroutine until it exits.
// It sends an error meSprintfge if there is any error during
// the API execution.
func (s *Server) Wait(waitChan chan error) {
	if err := s.serveAPI(); err != nil {
		logrus.Errorf("error: %v", err)
		waitChan <- err
		return
	}
	waitChan <- nil
}

func (s *Server) createMux() {
	logrus.Info("registering middleware")
	s.app.Use(s.makeIrisHandler(s.handlerWithGlobalMiddleware))
	logrus.Info("registering routers")
	for _, apiRouter := range s.routers {
		for _, r := range apiRouter.Routes() {
			f := s.makeIrisHandler(r.Handler())
			logrus.Debugf("registering route %s, %s", r.Method(), r.Path())
			s.app.Handle(r.Method(), r.Path(), f)
		}
	}
}

func (s *Server) makeIrisHandler(handler httputils.APIFunc) func(ctx context.Context) {
	return func(ctx context.Context) {
		requestID := uuid.Must(uuid.NewV4()).String()
		defer func() {
			if err := recover(); err != nil {
				logrus.Errorf("[Panic]-[%s] Handler for %s %s returned error: %v", requestID, ctx.Method(), ctx.Path(), err)
				ctx.JSON(httputils.NewFailedResponse(requestID, "thread panic"))
			}
		}()
		ctx.Params().Set("id", requestID)

		logrus.Infof("[Info %s]-[%s] %s %s", time.Now().Format("2006-01-02 15:04:05"), requestID, ctx.Method(), ctx.Path())
		if err := handler(ctx); err != nil {
			key := "not found"
			logrus.Errorf("[Error]-[%s] %s", requestID, err.Error())
			if found := strings.Contains(err.Error(), key); !found {
				ctx.JSON(httputils.NewFailedResponse(requestID, err.Error()))
			} else {
				ctx.JSON(httputils.NewSuccessResponse(requestID, err.Error()))
			}
			return
		}
		ctx.Next()
	}
}

func (s *Server) serveAPI() error {
	chError := make(chan error)
	go func() {
		var err error
		logrus.Infof("server listen on %s", s.cfg.Address)
		err = s.app.Run(iris.Addr(s.cfg.Address), iris.WithPathEscape)
		chError <- err
	}()
	err := <-chError
	return err
}
