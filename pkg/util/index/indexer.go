package index

import (
	"fmt"
	pt "path"

	"github.com/mojo-zd/helm-crabstick/pkg/util/path"
)

func GetIndexFile(name string) string {
	return pt.Join(path.GetRepoCacheDir(), fmt.Sprintf("%s-index.yaml", name))
}
