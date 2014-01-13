package filesystem

import (
	"os"
	"strings"
)

type FileFilter func(os.FileInfo) bool

func AllFiles(fileInfo os.FileInfo) bool {
	return !fileInfo.IsDir()
}

//TODO: other system files?
func IgnoreSystemFiles(fi os.FileInfo) bool {
	return !strings.HasPrefix(fi.Name(), ".")
}

func MinFilter(minSize uint64) FileFilter {
	return func(fi os.FileInfo) bool {
		return uint64(fi.Size()) >= minSize
	}
}

func MaxFilter(maxSize uint64) FileFilter {
	return func(fi os.FileInfo) bool {
		return uint64(fi.Size()) <= maxSize
	}
}

func ExtensionFilter(extensions []string) FileFilter {
	return func(fi os.FileInfo) bool {
		for _, ext := range extensions {
			if strings.HasSuffix(fi.Name(), ext) {
				return true
			}
		}
		return false
	}
}
