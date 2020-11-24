package helm

import (
	"net/http"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/repository"
	"github.com/mojo-zd/helm-crabstick/pkg/util/file"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/release/action/do"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/release/action/get"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Server interface {
}

type server struct {
	cfg        *config.Config
	kubeClient kubernetes.Interface
	httpClient *http.Client
	doer       do.Doer
	getter     get.Getter
	repo       repository.RepoHandler
}

// NewServer return helm server operator
func NewServer(cfgPath string) (Server, error) {
	conf, err := config.NewConfig(cfgPath)
	if err != nil {
		return nil, err
	}

	srv := &server{
		cfg: conf,
	}

	srv.init()
	return srv, nil
}

func (srv *server) init() error {
	file.CreateHelmDirIfNotExist() // prepare cache directory
	srv.httpClient = &http.Client{}

	config, err := clientcmd.BuildConfigFromFlags("", srv.cfg.KubeConf)
	if err != nil {
		return err
	}
	if srv.kubeClient, err = kubernetes.NewForConfig(config); err != nil {
		return err
	}
	return err
}
