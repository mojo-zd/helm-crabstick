package do

import "helm.sh/helm/v3/pkg/release"

func (d *doer) Install(name, chart string) (*release.Release, error) {
	return nil, nil
}

func (d *doer) Uninstall(name string) error {
	return nil
}
