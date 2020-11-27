package types

type ReleaseCreateOptions struct {
	Name      string      `json:"name"`
	Version   string      `json:"version"`
	ChartName string      `json:"chartName"`
	Namespace string      `json:"namespace"`
	Options   DoerOptions `json:"options"`
	Values    string      `json:"values"`
}

type DoerOptions struct {
	Annotation map[string]string `json:"annotation"`
}
