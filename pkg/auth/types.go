package auth

const (
	HealthStatus = "HEALTHY"
)

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

type Cluster struct {
	UUID         string `json:"uuid"`
	Name         string `json:"name"`
	NodeCount    int    `json:"node_count"`
	MasterCount  int    `json:"master_count"`
	Status       string `json:"status"`
	HealthStatus string `json:"health_status"`
	ApiAddress   string `json:"api_address"`
}

type Token struct {
	IsAdmin bool    `json:"is_admin"`
	Project Project `json:"project"`
	User    User    `json:"user"`
}

type Project struct {
	Domain Domain `json:"domain"`
	ID     string `json:"id"`
	Name   string `json:"name"`
}

type Domain struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
