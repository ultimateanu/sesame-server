package filesystem

import (
	"os"
)

func IsFile(p string) bool {
	fileinfo, err := os.Stat(p)
	if err == nil && !fileinfo.IsDir() {
		return true
	}
	return false
}

func IsDir(p string) bool {
	fileinfo, err := os.Stat(p)
	if err == nil && fileinfo.IsDir() {
		return true
	}
	return false
}
