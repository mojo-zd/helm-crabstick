package do

import (
	"testing"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
)

var (
	repo = config.Repository{
		Name: "bitmina",
		URL:  "https://charts.bitnami.com/bitnami",
	}
)

func TestSyncRepo(t *testing.T) {
	doer := NewDoer()
	if err := doer.SynRepo(repo); err != nil {
		t.Fatal(err)
	}
}
