package repository

import (
	"errors"
	"net/http"
	"time"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
)

var (
	repoSettingErr = errors.New("not found helm repository's setting, please set")
)

const timeout = time.Minute

type RepoHandler interface {
	CacheIndex() error
}

type repo struct {
	cfg    *config.Config
	client *http.Client
}

func NewRepo(cfg *config.Config) RepoHandler {
	return &repo{
		cfg:    cfg,
		client: util.NewHttpClient(timeout),
	}
}
