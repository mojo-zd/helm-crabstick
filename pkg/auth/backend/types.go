package backend

// Services ...
type Services struct {
	Services []Service `json:"services"`
}

// Service keystone service entity
type Service struct {
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Type        string `json:"type"`
	ID          string `json:"id"`
	Name        string `json:"name"`
}

// Endpoints ...
type Endpoints struct {
	Endpoints []Endpoint `json:"endpoints"`
}

// Endpoint keystone endpoint entity
type Endpoint struct {
	ID        string `json:"id"`
	RegionID  string `json:"region_id"`
	URL       string `json:"url"`
	Region    string `json:"region"`
	Enabled   bool   `json:"enabled"`
	Interface string `json:"interface"`
	ServiceID string `json:"service_id"`
}

// Certificate keystone certificate entity
type Certificate struct {
	ClusterUUID string `json:"cluster_uuid"`
	CSR         string `json:"csr"`
	PEM         string `json:"pem"`
}
