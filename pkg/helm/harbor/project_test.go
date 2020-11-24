package harbor

import (
	"testing"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
)

var (
	cfg = config.Config{
		Harbor: &config.Harbor{
			Username: "admin",
			Password: "Harbor12345",
			URL:      "http://118.31.50.65:10080",
		},
	}
)

func TestProjectGet(t *testing.T) {
	handler := NewProjectHandler(*cfg.Harbor)
	projects, _ := handler.Get("library")
	for _, project := range projects {
		t.Log(project.Name, project.ProjectId, project.Metadata.Public)
	}
}

func TestProjectCreate(t *testing.T) {
	handler := NewProjectHandler(*cfg.Harbor)
	if err := handler.Create(&Project{Name: "locast11", Metadata: &Metadata{Public: "false"}}); err != nil {
		t.Fatal(err)
		return
	}
}
