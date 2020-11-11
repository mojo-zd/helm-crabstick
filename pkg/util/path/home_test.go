package path

import "testing"

func TestPath(t *testing.T) {
	repoCachePath := GetRepositoryCacheDir()
	cacheHome := GetCacheHome()
	if err := MkRepositoryCacheDir(); err != nil {
		t.Fatal(err)
	}
	t.Log(repoCachePath, cacheHome)
}
