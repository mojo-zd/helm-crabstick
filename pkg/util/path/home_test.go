package path

import "testing"

func TestPath(t *testing.T) {
	repoCachePath := GetRepoCacheDir()
	cacheHome := GetCacheHome()
	if err := MkRepoCacheDirIfNotExist(); err != nil {
		t.Fatal(err)
	}
	t.Log(repoCachePath, cacheHome)
}
