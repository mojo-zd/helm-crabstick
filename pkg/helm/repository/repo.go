package repository

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	resty "github.com/go-resty/resty/v2"
	"github.com/gofrs/flock"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
	"github.com/mojo-zd/helm-crabstick/pkg/util/file"
	"github.com/mojo-zd/helm-crabstick/pkg/util/parser"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"helm.sh/helm/v3/pkg/getter"
	hr "helm.sh/helm/v3/pkg/repo"
)

// Config generate repositories.yaml
func (r *repo) Config() error {
	// Acquire a file lock for process synchronization
	fileLock := flock.New(strings.Replace(r.cfg.Repository.Config, filepath.Ext(r.cfg.Repository.Config), ".lock", 1))
	lockCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	locked, err := fileLock.TryLockContext(lockCtx, time.Second)
	if err == nil && locked {
		defer fileLock.Unlock()
	}
	if err != nil {
		return err
	}

	b, err := ioutil.ReadFile(file.GetConfig())
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	var f hr.File
	if err := yaml.Unmarshal(b, &f); err != nil {
		return err
	}

	c := hr.Entry{
		Name:                  r.cfg.Repository.Name,
		URL:                   r.cfg.Repository.URL,
		Username:              r.cfg.Repository.Username,
		Password:              r.cfg.Repository.Password,
		CertFile:              r.cfg.Repository.CertFile,
		KeyFile:               r.cfg.Repository.KeyFile,
		CAFile:                r.cfg.Repository.CaFile,
		InsecureSkipTLSverify: r.cfg.Repository.InsecureSkipTlsVerify,
	}

	chartRepo, err := hr.NewChartRepository(&c, getter.All(util.NewSetting(r.cfg)))
	if err != nil {
		return err
	}

	chartRepo.CachePath = r.cfg.Repository.Cache
	if r.cfg.Repository.Cache != "" {
		chartRepo.CachePath = r.cfg.Repository.Cache
	}
	if _, err := chartRepo.DownloadIndexFile(); err != nil {
		return errors.Wrapf(err, "looks like %q is not a valid chart repository or cannot be reached", r.cfg.Repository.URL)
	}

	f.Update(&c)

	if err := f.WriteFile(r.cfg.Repository.Config, 0644); err != nil {
		return err
	}
	return nil
}

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

	req := resty.New().R()
	req.SetBasicAuth(r.cfg.Repository.Username, r.cfg.Repository.Password)
	resp, err := req.Get(r.indexURL(url))
	if err != nil {
		logrus.Errorf("new request failed err:%s", err.Error())
	}
	logrus.Infof("start sync repository[%s] index...", r.indexURL(url))
	if resp.StatusCode() >= http.StatusBadRequest {
		return errors.New(fmt.Sprintf("can't get repository[%s] index file", r.cfg.Repository.Name))
	}
	return ioutil.WriteFile(file.GetIndexFile(rep.Name), resp.Body(), 0755)
}

func (r *repo) indexURL(repoURL string) string {
	return fmt.Sprintf("%s/%s", repoURL, "index.yaml")
}
