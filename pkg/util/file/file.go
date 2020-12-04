package file

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	CacheHomeEnvVar  = "HELM_CACHE_HOME"
	ConfigHomeEnvVar = "HELM_CONFIG_HOME"
	lp               = "helm"
)

// RepoIndexExist check repo index file exist
func RepoIndexExist(repoName string) bool {
	_, err := os.Stat(GetIndexFile(repoName))
	return !os.IsNotExist(err)
}

// GetIndexFile get helm index file
func GetIndexFile(name string) string {
	return path.Join(GetCacheDir(), fmt.Sprintf("%s-index.yaml", name))
}

// GetCacheRepositoryDir repository's cache dir
func GetCacheDir() string {
	return path.Join(GetCacheHome(), "repository")
}

// GetCacheHome the cache path of helm
func GetCacheHome() string {
	cacheDir := os.Getenv(CacheHomeEnvVar)
	if strings.TrimSpace(cacheDir) == "" {
		cacheDir = path.Join(HomeDir(), ".cache", lp)
	}
	return cacheDir
}

// GetConfigDir get helm config directory
func GetConfigDir() string {
	configDir := os.Getenv(ConfigHomeEnvVar)
	if strings.TrimSpace(configDir) == "" {
		configDir = path.Join(HomeDir(), ".config", lp)
	}
	return configDir
}

// GetConfig return repositories.yaml's location
func GetConfig() string {
	return path.Join(GetConfigDir(), "repositories.yaml")
}

// CreateHelmDirIfNotExist prepare helm cache dir and repository config dir
func CreateHelmDirIfNotExist() error {
	configDir := GetConfigDir()
	cacheDir := GetCacheDir()
	if _, err := os.Stat(cacheDir); err != nil && os.IsNotExist(err) {
		if err = os.MkdirAll(cacheDir, 0755); err != nil {
			logrus.Errorf("make helm cache dir[%s] failed, err:%s", cacheDir, err.Error())
			return err
		}
	}

	if _, err := os.Stat(configDir); err != nil && os.IsNotExist(err) {
		if err = os.MkdirAll(configDir, 0755); err != nil {
			logrus.Errorf("make helm config dir[%s] failed, err:%s", configDir, err.Error())
			return err
		}
	}
	return nil
}

func HomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOME")
		homeDriveHomePath := ""
		if homeDrive, homePath := os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"); len(homeDrive) > 0 && len(homePath) > 0 {
			homeDriveHomePath = homeDrive + homePath
		}
		userProfile := os.Getenv("USERPROFILE")

		// Return first of %HOME%, %HOMEDRIVE%/%HOMEPATH%, %USERPROFILE% that contains a `.kube\config` file.
		// %HOMEDRIVE%/%HOMEPATH% is preferred over %USERPROFILE% for backwards-compatibility.
		for _, p := range []string{home, homeDriveHomePath, userProfile} {
			if len(p) == 0 {
				continue
			}
			if _, err := os.Stat(filepath.Join(p, ".kube", "config")); err != nil {
				continue
			}
			return p
		}

		firstSetPath := ""
		firstExistingPath := ""

		// Prefer %USERPROFILE% over %HOMEDRIVE%/%HOMEPATH% for compatibility with other auth-writing tools
		for _, p := range []string{home, userProfile, homeDriveHomePath} {
			if len(p) == 0 {
				continue
			}
			if len(firstSetPath) == 0 {
				// remember the first path that is set
				firstSetPath = p
			}
			info, err := os.Stat(p)
			if err != nil {
				continue
			}
			if len(firstExistingPath) == 0 {
				// remember the first path that exists
				firstExistingPath = p
			}
			if info.IsDir() && info.Mode().Perm()&(1<<(uint(7))) != 0 {
				// return first path that is writeable
				return p
			}
		}

		// If none are writeable, return first location that exists
		if len(firstExistingPath) > 0 {
			return firstExistingPath
		}

		// If none exist, return first location that is set
		if len(firstSetPath) > 0 {
			return firstSetPath
		}

		// We've got nothing
		return ""
	}
	return os.Getenv("HOME")
}
