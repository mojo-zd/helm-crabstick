package harbor

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/sirupsen/logrus"
)

const (
	Base        = "/api/v2.0"
	ProjectPath = "projects"
)

type Project struct {
	Name       string    `json:"project_name"`
	ProjectId  int64     `json:"project_id"`
	ChartCount int64     `json:"chart_count"`
	Metadata   *Metadata `json:"metadata"`
	Deleted    bool      `json:"deleted"`
}

type Metadata struct {
	Public string `json:"public"`
}

type Value struct {
	Path  string
	Query map[string]string
}

func doRequest(method string, value Value, body io.Reader, repository config.Repository) ([]byte, error) {
	client := &http.Client{}
	request, err := newRequest(method, value, body, repository)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		logrus.Errorf("execute http do request failed, err:%s", err.Error())
		return nil, err
	}
	defer response.Body.Close()

	out, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Errorf("read http body failed, err:%s", err.Error())
		return nil, err
	}
	return out, nil
}

func newRequest(method string, value Value, body io.Reader, repository config.Repository) (*http.Request, error) {
	u, err := url.Parse(repository.URL)
	if err != nil {
		logrus.Errorf("parse url[%s] failed,err:%s", repository.URL, err.Error())
		return nil, err
	}
	u.Path = path.Join(u.Path, value.Path)
	u.RawQuery = values(value.Query).Encode()
	request, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		logrus.Errorf("new request[%s] failed, err:%s", u.String(), err.Error())
		return nil, err
	}
	request.SetBasicAuth(repository.Username, repository.Password)
	request.Header.Set("Content-Type", "application/json")
	return request, nil
}

func values(val map[string]string) url.Values {
	value := url.Values{}
	if val == nil {
		return value
	}
	for k, v := range val {
		value.Add(k, v)
	}
	return value
}
