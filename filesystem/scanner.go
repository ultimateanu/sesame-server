package filesystem

import (
	"os"
	"path/filepath"
)

func ScanDirs(paths []string, filters []FileFilter) ([]*File, error) {
	allfiles := make([]*File, 0)
	for _, p := range paths {
		files, err := ScanDir(p, filters)
		if err != nil {
			return nil, err
		}
		allfiles = append(allfiles, files...)
	}
	return allfiles, nil
}

func ScanDir(p string, filters []FileFilter) ([]*File, error) {
	var fileInfos []*File
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		includeFile := true
		for _, filter := range filters {
			if !filter(info) {
				includeFile = false
			}
		}
		if includeFile {
			fileInfos = append(fileInfos, MakeFile(info, path))
		}

		return nil
	}

	err := filepath.Walk(p, walkFunc)
	if err != nil {
		return nil, err
	}
	return fileInfos, nil
}
