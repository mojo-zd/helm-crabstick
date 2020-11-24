package harbor

import (
	"bytes"
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
	"github.com/sirupsen/logrus"
)

type ProjectHandler interface {
	Get(name string) ([]*Project, error)
	Create(project *Project) error
}

type projectHandler struct {
	repository config.Repository
}

func NewProjectHandler(repository config.Repository) ProjectHandler {
	return &projectHandler{
		repository: repository,
	}
}

func (p *projectHandler) Create(project *Project) error {
	pro, err := json.Marshal(project)
	if err != nil {
		logrus.Errorf("project[%+v] invalid, err:%s", project, err.Error())
		return err
	}

	if _, err = doRequest(
		http.MethodPost,
		Value{Path: path.Join(Base, ProjectPath)},
		bytes.NewBuffer(pro),
		p.repository,
	); err != nil {
		logrus.Errorf("create project[%s] failed, err:%s", string(pro), err.Error())
		return nil
	}
	return nil
}

func (p *projectHandler) Get(name string) ([]*Project, error) {
	pro := []*Project{}
	value := Value{
		Path: path.Join(Base, ProjectPath),
		Query: map[string]string{
			"page":      strconv.FormatInt(util.DefaultPageSize, 10),
			"page_size": strconv.FormatInt(util.DefaultPageSize, 10),
			"name":      name,
		},
	}
	out, err := doRequest(http.MethodGet, value, nil, p.repository)
	if err != nil {
		logrus.Errorf("get project[%s] failed, err:%s", name, err.Error())
		return nil, err
	}
	if err = json.Unmarshal(out, &pro); err != nil {
		logrus.Errorf("unmarshal project failed, err:%s", err.Error())
		return nil, err
	}
	return pro, nil
}
