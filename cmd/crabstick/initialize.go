package main

import (
	"github.com/mojo-zd/helm-crabstick/api/server"
	"github.com/mojo-zd/helm-crabstick/api/server/router"
	"github.com/mojo-zd/helm-crabstick/api/server/router/chart"
	"github.com/mojo-zd/helm-crabstick/api/server/router/release"
	"github.com/mojo-zd/helm-crabstick/cmd/crabstick/config"
	appconf "github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/repository"
	"github.com/mojo-zd/helm-crabstick/pkg/util/file"
)

func initialize(cfg appconf.Config) {
	// prepare helm dir e.g. cache dir、config dir
	file.CreateHelmDirIfNotExist()
	// cache repository index.yaml
	if config.GetConfig().RunMode == "dev" {
		if !file.RepoIndexExist(cfg.Repository.Name) {
			repository.NewRepo(cfg).CacheIndex()
		}
		return
	}
	repository.NewRepo(cfg).CacheIndex()
}

func routes() []router.Router {
	cfg := newAppConfig()
	routers := []router.Router{
		chart.NewRouter(cfg),
		release.NewRouter(cfg),
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

func newAppConfig() appconf.Config {
	cfg := config.GetConfig()
	return appconf.Config{
		Repository: &appconf.Repository{
			Name:     cfg.Repository.Name,
			URL:      cfg.Repository.URL,
			Username: cfg.Repository.Username,
			Password: cfg.Repository.Password,
			Cache:    file.GetCacheDir(),
		},
	}
}
