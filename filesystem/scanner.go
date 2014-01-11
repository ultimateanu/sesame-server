package filesystem

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
)

func ExtractDirs(paths []string) ([]*File, error) {
	files := make([]*File, 0, 10)
	for _, p := range paths {
		f, err := ExtractDir(p)
		if err != nil {
			return nil, err
		}
		files = append(files, f...)
	}
	return files, nil
}

func ExtractDir(p string) (files []*File, err error) {
	files = make([]*File, 0, 10)
	if IsFile(p) {
		var file *File
		file, err = GetFile(p)
		files = append(files, file)
	} else if IsDir(p) {
		files, err = GetFiles(p)
	} else {
		err = errors.New(p + " is neither a file nor a directory")
	}

	if err != nil {
		return nil, err
	}
	return
}

func GetFiles(p string) ([]*File, error) {
	fileinfos, err := ioutil.ReadDir(p)
	if err != nil {
		return nil, err
	}

	files := make([]*File, 0, 10)
	for _, fileinfo := range fileinfos {
		filePath := path.Join(p, fileinfo.Name())
		if fileinfo.IsDir() {
			f, err := GetFiles(filePath)
			if err != nil {
				return nil, err
			}
			files = append(files, f...)
		} else {
			files = append(files, MakeFile(fileinfo, filePath))
		}
	}
	return files, nil
}

func GetFile(p string) (*File, error) {
	fileinfo, err := os.Stat(p)
	if err != nil {
		return nil, err
	}

	if fileinfo.IsDir() {
		return nil, errors.New(p + " is a directory")
	}

	return MakeFile(fileinfo, p), nil
}
