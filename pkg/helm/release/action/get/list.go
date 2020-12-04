package get

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/storage"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/release"
)

// List list all release by condition
func (g *getter) List(namespace string, opts util.ListOptions) ([]*release.Release, error) {
	cfg := storage.ActionConfiguration(*g.cluster, g.config, namespace)
	client := action.NewList(cfg)
	client.SetStateMask()
	releases, err := client.Run()
	if err != nil {
		logrus.Errorf("can't list release, err:%s", err.Error())
		return nil, err
	}
	return filterWithOpts(releases, opts), nil
}

func filterWithOpts(releases []*release.Release, opts util.ListOptions) []*release.Release {
	if opts.Annotation == nil || len(opts.Annotation) == 0 {
		return releases
	}
	result := make([]*release.Release, 0)
	for _, rls := range releases {
		if rls.Chart != nil && rls.Chart.Metadata != nil {
			acp := copymap(opts.Annotation)
			for k, v := range rls.Chart.Metadata.Annotations {
				if val, ok := acp[k]; ok && val == v {
					delete(acp, k)
				}
			}

			if len(acp) == 0 {
				result = append(result, rls)
			}
		}
	}
	return result
}

func copymap(m map[string]string) map[string]string {
	if m == nil {
		return m
	}
	r := make(map[string]string)
	for k, v := range m {
		r[k] = v
	}
	return r
}
