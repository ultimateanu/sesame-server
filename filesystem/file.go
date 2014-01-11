package filesystem

import (
	"os"
	"strings"
)

type File struct {
	Name string
	Path string
	Size int64
}

func MakeFile(fileinfo os.FileInfo, p string) *File {
	return &File{fileinfo.Name(), p, fileinfo.Size()}
}

func Filter(files []*File, fn func(*File) bool) []*File {
	filteredFiles := make([]*File, 0, 10)
	for _, f := range files {
		if fn(f) {
			filteredFiles = append(filteredFiles, f)
		}
	}
	return filteredFiles
}

func FileExtension(extensions []string) func(*File) bool {
	return func(f *File) bool {
		for _, ext := range extensions {
			if strings.HasSuffix(f.Name, ext) {
				return true
			}
		}
		return false
	}
}
