package main

import (
	"github.com/mojo-zd/helm-crabstick/api/server"
	"github.com/mojo-zd/helm-crabstick/api/server/router"
	"github.com/mojo-zd/helm-crabstick/api/server/router/chart"
	"github.com/mojo-zd/helm-crabstick/api/server/router/release"
	"github.com/mojo-zd/helm-crabstick/cmd/crabstick/config"
	"github.com/sirupsen/logrus"
)

func main() {
	srv := server.New(newAPIConfig())
	srv.InitRouter(routes()...)
	serveAPIWait := make(chan error)

	go srv.Wait(serveAPIWait)
	// Wait for serve API to complete
	errAPI := <-serveAPIWait
	if errAPI != nil {
		logrus.Warn(errAPI)
	}
}

func routes() []router.Router {
	routers := []router.Router{
		chart.NewRouter(),
		release.NewRouter(),
	}
	return routers
}

func newAPIConfig() *server.Config {
	cfg := config.GetConfig()
	return &server.Config{
		Address:    cfg.Address,
		Middleware: cfg.Middlewares,
		JWTSecret:  cfg.SecretKey,
	}
}
