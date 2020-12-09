package service

import (
	"github.com/mojo-zd/helm-crabstick/data/query"
	db "github.com/mojo-zd/helm-crabstick/data/types"
	"github.com/mojo-zd/helm-crabstick/pkg/auth"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/manager"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/parser"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/types"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
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
	rel, err := manager.NewAppManager(cfg, &cluster).ReleaseDoer.Install(createOpts)
	if err != nil {
		return rel, err
	}
	release = buildRelease(token, createOpts)
	release.ClusterUUID = cluster.UUID
	if err = r.releaseDao.Create(release); err != nil {
		return nil, err
	}
	return rel, nil
}

// Delete delete release from cluster
func (r *ReleaseService) Delete(cfg config.Config, cluster mg.Cluster, name, namespace string) error {
	if _, err := manager.NewAppManager(cfg, &cluster).ReleaseDoer.Delete(name, namespace); err != nil {
		return err
	}

	return r.releaseDao.DeleteBathes(
		&db.Release{},
		map[string]interface{}{
			"name":         name,
			"namespace":    namespace,
			"cluster_uuid": cluster.UUID,
		})
}

// List return user's release
func (r *ReleaseService) List(
	cfg config.Config,
	token auth.Token,
	cluster mg.Cluster) ([]*parser.Release, error) {
	releases, err := r.releaseDao.List(
		&db.Release{
			ProjectID:   token.Project.ID,
			DomainID:    token.Project.Domain.ID,
			CreatorID:   token.User.ID,
			ClusterUUID: cluster.UUID,
		},
	)
	if err != nil {
		return []*parser.Release{}, err
	}

	rls, err := manager.NewAppManager(cfg, &cluster).ReleaseGetter.List("", util.ListOptions{})
	if err != nil {
		return []*parser.Release{}, err
	}

	out := []*release.Release{}
	for _, r := range releases {
		for _, rel := range rls {
			if rel.Name == r.Name {
				out = append(out, rel)
				break
			}
		}
	}
	return parser.ToReleases(out), nil
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
