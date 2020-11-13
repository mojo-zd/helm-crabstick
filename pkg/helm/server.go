package helm

import (
	"net/http"

	"github.com/mojo-zd/helm-crabstick/pkg/util/path"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/release/action/do"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/release/action/get"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/resources"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Server interface {
}

type server struct {
	cfg        *config.Config
	kubeClient kubernetes.Interface
	httpClient *http.Client
	resource   resources.Resource
	doer       do.Doer
	getter     get.Getter
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
	path.MkRepositoryCacheDir() // prepare cache directory
	srv.httpClient = &http.Client{}

	config, err := clientcmd.BuildConfigFromFlags("", srv.cfg.KubeConf)
	if err != nil {
		return err
	}
	if srv.kubeClient, err = kubernetes.NewForConfig(config); err != nil {
		return err
	}

	srv.resource = resources.NewResource(srv.kubeClient)
	return err
}
