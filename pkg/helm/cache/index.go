package cache

import (
	"github.com/mojo-zd/helm-crabstick/pkg/util/file"
	"helm.sh/helm/v3/pkg/repo"
)

type IndexCache interface {
	LoadIndex() (*repo.IndexFile, error)
}

type indexCache struct {
	repoName string
}

func NewIndexCache(repo string) IndexCache {
	return &indexCache{repoName: repo}
}

// LoadIndex load repository index yaml
func (cache *indexCache) LoadIndex() (*repo.IndexFile, error) {
	return repo.LoadIndexFile(file.GetIndexFile(cache.repoName))
}
