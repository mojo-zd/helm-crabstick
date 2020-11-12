package template

import (
	"errors"
	"strconv"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
)

type Placeholder string

const (
	CHARTNAME   = "<CHARTNAME>"
	SERVICENAME = "<SERVICENAME>"
	INDEX       = "<INDEX>"
)

type ResourceType string

const (
	Deploy    = "deployments"
	Deamonset = "deamonset"
)

type Values struct {
	ChartName string
	Services  []Service
}

type Service struct {
	Name string
	Type string
}

type Render interface {
	Replacer() []string
	HelperReplacer() string
}

type render struct {
	values Values
}

func NewRender(values string) (Render, error) {
	r := &render{}
	vv, err := r.fetchBasic(values)
	if err != nil {
		return nil, err
	}
	r.values = vv
	return r, nil
}

// Replacer handler of kuernetes resources
func (r *render) Replacer() []string {
	result := []string{}
	for index, service := range r.values.Services {
		switch service.Type {
		case Deploy:
			o := strings.ReplaceAll(deployTemplate, CHARTNAME, r.values.ChartName)
			o = strings.ReplaceAll(o, SERVICENAME, service.Name)
			o = strings.ReplaceAll(o, INDEX, strconv.FormatInt(int64(index), 10))
			result = append(result, o)
		}
	}
	return result
}

func (r *render) HelperReplacer() string {
	return strings.ReplaceAll(helperTemplate, CHARTNAME, r.values.ChartName)
}

func (r *render) fetchBasic(values string) (Values, error) {
	result := Values{}
	out, err := r.strToMap(values)
	if err != nil {
		return result, err
	}
	chartName, ok := out["name"]
	if !ok {
		return result, errors.New("values invalid, values should match the format of example values.yaml")
	}
	result.ChartName = chartName.(string)
	services, ok := out["services"]
	if !ok {
		return result, errors.New("service node not found, at least 1 node")
	}

	vals, ok := services.([]interface{})
	if !ok {
		return result, errors.New("service")
	}

	// 遍历service层
	for _, svrs := range vals {
		ss, err := service(svrs.(map[string]interface{}))
		if err != nil {
			return result, err
		}
		result.Services = append(result.Services, ss)
	}
	return result, nil
}

func service(svr map[string]interface{}) (Service, error) {
	ss := Service{}
	for key, value := range svr {
		vv := value.(map[string]interface{})
		typ, ok := vv["type"]
		if !ok {
			return ss, errors.New("service define invalid")
		}
		ss.Name = key
		ss.Type = typ.(string)
	}
	return ss, nil
}

// strToMap convert string to map
func (r *render) strToMap(values string) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	err := yaml.Unmarshal([]byte(values), &out)
	if err != nil {
		logrus.Errorf("unmarshal failed, err:%s", err.Error())
		return out, err
	}
	return out, err
}
