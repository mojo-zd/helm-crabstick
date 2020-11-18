package repository

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mojo-zd/helm-crabstick/pkg/util/index"
	"github.com/mojo-zd/helm-crabstick/pkg/util/parser"
	"github.com/mojo-zd/helm-crabstick/pkg/util/path"
	"github.com/sirupsen/logrus"
)

func (r *repo) CacheIndex() error {
	rep := r.cfg.Repository
	if rep == nil {
		logrus.Error(repoSettingErr.Error())
		return repoSettingErr
	}

	url, err := parser.ParseURL(rep.URL)
	if err != nil {
		logrus.Errorf("request url invalid err:%s", err.Error())
		return err
	}

	request, err := http.NewRequest(http.MethodGet, r.indexURL(url), nil)
	if err != nil {
		logrus.Errorf("new request failed err:%s", err.Error())
	}
	logrus.Infof("start sync repository[%s] index...", url)
	response, err := r.client.Do(request)
	if err != nil {
		logrus.Errorf("can't get repository[%s] index.yaml", url)
		return err
	}

	defer func() {
		response.Body.Close()
	}()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	path.MkRepoCacheDirIfNotExist()
	return ioutil.WriteFile(index.GetIndexFile(rep.Name), data, 0755)
}

func (r *repo) indexURL(repoURL string) string {
	return fmt.Sprintf("%s/%s", repoURL, "index.yaml")
}
