package get

import (
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	kindReg = `kind:\s(\w+)`
)

// ReleaseKind  list the release's all kind of kubernetes
func (g *getter) ReleaseKind(name, namespace string) []string {
	out := []string{}
	release, err := g.Get(name, namespace)
	if err != nil {
		logrus.Errorf("get release failed, err:%s", err.Error())
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
	return kinds
}
