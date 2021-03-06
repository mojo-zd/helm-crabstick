package harbor

import (
	"testing"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
)

var (
	cfg = config.Config{
		Repository: &config.Repository{
			Username: "admin",
			Password: "Harbor12345",
			URL:      "http://118.31.50.65:10080",
		},
	}
)

func TestProjectGet(t *testing.T) {
	handler := NewProjectHandler(*cfg.Repository)
	projects, _ := handler.Get("library")
	for _, project := range projects {
		t.Log(project.Name, project.ProjectId, project.Metadata.Public)
	}
}

func TestProjectCreate(t *testing.T) {
	handler := NewProjectHandler(*cfg.Repository)
	if err := handler.Create(&Project{Name: "locast11", Metadata: &Metadata{Public: "false"}}); err != nil {
		t.Fatal(err)
		return
	}
}
