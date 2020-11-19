package do

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/storage"
	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/release"
)

func (d *doer) Delete(name, namespace string) (*release.UninstallReleaseResponse, error) {
	configuration := storage.ActionConfiguration(d.client, d.cfg, namespace)
	uninstall := action.NewUninstall(configuration)
	out, err := uninstall.Run(name)
	if err != nil {
		logrus.Errorf("uninstall release[%s] failed, err:%s", name, err.Error())
		return nil, err
	}
	return out, err
}
