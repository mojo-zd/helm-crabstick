package get

import (
	"regexp"
	"strings"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kindReg = `kind:\s(\w+)`
)

// Kind  list all resource's kind of release e.g deploy、ingress、service ...
func (g *getter) Kind(name, namespace string) []string {
	out := []string{}
	release, err := g.Get(name, namespace)
	if err != nil {
		logrus.Errorf("not found release [%s] and namespace is [%s], err:%s", name, namespace, err.Error())
		return out
	}
	reg, err := regexp.Compile(kindReg)
	if err != nil {
		logrus.Errorf("regexp compile failed, err:%s", err.Error())
		return out
	}
	out = reg.FindAllString(release.Manifest, -1)
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
	logrus.Debugf("all kind of release [%s]:%+v", name, kinds)
	return kinds
}

// Resources get resources from kubernetes
func (g *getter) Resources(name, namespace string, opts v1.ListOptions) map[util.KubeKind]interface{} {
	result := make(map[util.KubeKind]interface{})
	kinds := g.Kind(name, namespace)
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
