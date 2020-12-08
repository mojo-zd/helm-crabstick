package main

import (
	"github.com/mojo-zd/helm-crabstick/api/server"
	"github.com/mojo-zd/helm-crabstick/pkg/manager"
	"github.com/sirupsen/logrus"
)

func main() {
	initialize()
	apiConf := newAPIConfig()
	srv := server.New(apiConf)
	srv.InitRouter(routes(manager.NewClusterManager(apiConf.KeystoneAddr))...)

	serveAPIWait := make(chan error)
	go srv.Wait(serveAPIWait)
	// Wait for serve API to complete
	errAPI := <-serveAPIWait
	if errAPI != nil {
		logrus.Warn(errAPI)
	}
}
