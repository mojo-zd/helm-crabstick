package action

type GetOps string

const (
	GetAll      GetOps = "all"
	GetHooks    GetOps = "hooks"
	GetManifest GetOps = "manifest"
	GetNotes    GetOps = "notes"
	GetValues   GetOps = "values"
)

type DoerOps struct {
	CreateNamespace string
	Namespace       string // namespace of release
	Description     string // custom description for release
}
