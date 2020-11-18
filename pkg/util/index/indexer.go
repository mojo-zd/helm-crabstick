package index

import (
	"fmt"

	"github.com/mojo-zd/helm-crabstick/pkg/util/path"
)

func GetIndexFile(name string) string {
	return fmt.Sprintf("%s/%s-index.yaml", path.GetRepoCacheDir(), name)
}
