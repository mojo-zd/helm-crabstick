package do

import (
	"net/http"
	"time"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"

	"helm.sh/helm/v3/pkg/release"
)

const timeout = 1 * time.Minute

// Doer include install、uninstall operator
type Doer interface {
	Install(name, chart string) (*release.Release, error)
	Uninstall(name string) error
}

type doer struct {
	client *http.Client
}

// NewDoer ...
func NewDoer() Doer {
	return &doer{
		client: util.NewHttpClient(timeout),
	}
}
