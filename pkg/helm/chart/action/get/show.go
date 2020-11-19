package get

import (
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/action"
)

func (g *getter) Show(name, version string, output action.ShowOutputFormat) string {
	if output == "" {
		output = action.ShowAll
	}
	client := action.NewShow(output)
	client.OutputFormat = output
	out, _ := g.run(name, version, client)
	return out
}

func (g *getter) run(name, version string, client *action.Show) (string, error) {
	if client.Version == "" && client.Devel {
		logrus.Debug("setting version to >0.0.0-0")
		client.Version = ">0.0.0-0"
	}

	chartOption, err := util.LoadChartOptions(g.cfg)
	if err != nil {
		logrus.Errorf("can't instance chart options object, err:%s", err.Error())
		return "", err
	}
	chartOption.Version = version
	client.ChartPathOptions = chartOption
	cp, err := client.ChartPathOptions.LocateChart(name, util.NewSetting(g.cfg))
	if err != nil {
		logrus.Errorf("locate chart failed, err:%s", err.Error())
		return "", err
	}
	out, err := client.Run(cp)
	if err != nil {
		logrus.Errorf("can't show chart information, err:%s", err.Error())
		return "", err
	}
	return out, nil
}
