package filesystem

import (
	"os"
	"strings"
)

type FileFilter func(os.FileInfo) bool

func AllFiles(fileInfo os.FileInfo) bool {
	return !fileInfo.IsDir()
}

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
