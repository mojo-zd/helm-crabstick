package get

import "helm.sh/helm/v3/pkg/release"

// Status get release status and set chart to nil
func (g *getter) Status(name, namespace string) (*release.Release, error) {
	release, err := g.Get(name, namespace)
	if err != nil {
		return nil, err
	}
	release.Chart = nil
	return release, nil
}
