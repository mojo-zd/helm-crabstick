package service

import (
	"github.com/mojo-zd/helm-crabstick/data/query"
	db "github.com/mojo-zd/helm-crabstick/data/types"
	"github.com/mojo-zd/helm-crabstick/pkg/auth"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/manager"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/types"
	mg "github.com/mojo-zd/helm-crabstick/pkg/manager"
	"helm.sh/helm/v3/pkg/release"
)

// ReleaseService ...
type ReleaseService struct {
	releaseDao query.ReleaseDao
}

func NewReleaseService() *ReleaseService {
	return &ReleaseService{
		releaseDao: query.ReleaseDao{},
	}
}

// Create install chart
func (r *ReleaseService) Create(
	cfg config.Config,
	cluster mg.Cluster,
	token auth.Token,
	createOpts types.CreateOptions,
) (*release.Release, error) {
	var (
		err     error
		release *db.Release
	)
	defer func() {
		// rollback if err is not nil
		if err != nil {
			r.releaseDao.Delete(release)
		}
	}()

	release = buildRelease(token, createOpts)
	if err = r.releaseDao.Create(release); err != nil {
		return nil, err
	}
	return manager.NewAppManager(cfg, &cluster).ReleaseDoer.Install(createOpts)
}

func buildRelease(token auth.Token, createOpts types.CreateOptions) *db.Release {
	return &db.Release{
		IsAdmin:   token.IsAdmin,
		Project:   token.Project.Name,
		ProjectID: token.Project.ID,
		Domain:    token.Project.Domain.Name,
		DomainID:  token.Project.Domain.ID,
		Name:      createOpts.Name,
		Namespace: createOpts.Namespace,
		Creator:   token.User.Name,
		CreatorID: token.User.ID,
	}
}
