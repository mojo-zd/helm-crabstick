package types

// the instance of application
type Release struct {
	ID        int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	IsAdmin   bool   `json:"is_admin"`  // map to keystone is_admin
	Project   string `json:"project"`   // map to keystone project.name
	ProjectID string `json:"projectId"` // map to keystone project.id
	Domain    string `json:"domain"`    // map to keystone project.domain.name
	DomainID  string `json:"domainId"`  // map to keystone project.domain.id
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Creator   string `json:"creator"`   // map to keystone user
	CreatorID string `json:"creatorId"` // map to keystone user.id
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}
