package types

// CreateOptions the options of create release
type CreateOptions struct {
	Name      string      `json:"name"`
	Version   string      `json:"version"`
	Chart     string      `json:"chart"`
	Namespace string      `json:"namespace"`
	Options   DoerOptions `json:"options"`
	Values    string      `json:"values"`
}

// DoerOptions ...
type DoerOptions struct {
	Annotation map[string]string `json:"annotation"`
}

// UpgradeOptions ...
type UpgradeOptions struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Chart     string `json:"chart"`
	Version   string `json:"version"`
	Values    string `json:"values"`
}
