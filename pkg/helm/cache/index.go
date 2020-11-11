package cache

import (
	"fmt"

	"github.com/mojo-zd/helm-crabstick/pkg/util/path"
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

func (cache *indexCache) LoadIndex() (*repo.IndexFile, error) {
	indexFile := fmt.Sprintf("%s/%s-index.yaml", path.GetRepositoryCacheDir(), cache.repoName)
	return repo.LoadIndexFile(indexFile)
}
