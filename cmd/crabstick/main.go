package main

import (
	"github.com/mojo-zd/helm-crabstick/api/server"
	"github.com/sirupsen/logrus"
)

func main() {
	initialize(newAppConfig())
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
