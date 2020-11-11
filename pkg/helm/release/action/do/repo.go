package do

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mojo-zd/helm-crabstick/pkg/util/index"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/mojo-zd/helm-crabstick/pkg/util/parser"
	"github.com/sirupsen/logrus"
)

// SynRepo synchronize repository information
func (d *doer) SynRepo(repo config.Repository) error {
	url, err := parser.ParseURL(repo.URL)
	if err != nil {
		logrus.Errorf("request url invalid err:%s", err.Error())
		return err
	}

	request, err := http.NewRequest(http.MethodGet, d.indexURL(url), nil)
	if err != nil {
		logrus.Errorf("new request failed err:%s", err.Error())
	}

	response, err := d.client.Do(request)
	if err != nil {
		logrus.Errorf("can't get repository index.yaml")
		return err
	}

	defer func() {
		response.Body.Close()
	}()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(index.GetIndexFile(repo.Name), data, 0755)
}

func (d *doer) indexURL(repoURL string) string {
	return fmt.Sprintf("%s/%s", repoURL, "index.yaml")
}
