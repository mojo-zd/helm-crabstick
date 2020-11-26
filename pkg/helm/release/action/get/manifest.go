package get

import (
	"regexp"
	"strings"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/release"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kindReg = `kind:\s(\w+)`
)

// Kind  list all resource's kind of release e.g deploy、ingress、service ...
func (g *getter) Kind(rls *release.Release) []string {
	out := []string{}
	reg, err := regexp.Compile(kindReg)
	if err != nil {
		logrus.Errorf("regexp compile failed, err:%s", err.Error())
		return out
	}
	out = reg.FindAllString(rls.Manifest, -1)
	kinds := []string{}
	single := make(map[string]bool)
	for _, val := range out {
		if _, ok := single[val]; ok {
			continue
		}
		single[val] = true
		args := strings.Split(val, ":")
		if len(args) != 2 {
			logrus.Warnf("skip kind[%s]", val)
			continue
		}
		kinds = append(kinds, strings.TrimSpace(args[1]))
	}
	logrus.Debugf("all kind of release [%s]:%+v", rls.Name, kinds)
	return kinds
}

// Resources get resources from kubernetes
func (g *getter) Resources(name, namespace string, rls *release.Release, opts v1.ListOptions) map[util.KubeKind]interface{} {
	result := make(map[util.KubeKind]interface{})
	kinds := g.Kind(rls)
	for _, kind := range kinds {
		out, err := g.manager.GetResources(util.KubeKind(kind)).List(
			namespace,
			opts,
		)
		if err != nil {
			logrus.Warnf("can't get [%s] resources of release [%s], err:%s", kind, name, err.Error())
			continue
		}
		result[util.KubeKind(kind)] = out
	}
	return result
}
