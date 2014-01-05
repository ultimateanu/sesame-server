package filesystem

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func GetAllVideoFiles(p string) ([]*Video, error) {
	files, err := ioutil.ReadDir(p)
	if err != nil {
		return nil, err
	}

	videos, err := GetVideoFiles(p)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			v, err := GetAllVideoFiles(path.Join(p, file.Name()))
			if err != nil {
				return nil, err
			}
			videos = append(videos, v...)
		}
	}

	return videos, nil
}

func GetVideoFiles(p string) ([]*Video, error) {
	files, err := ioutil.ReadDir(p)
	if err != nil {
		return nil, err
	}

	videos := make([]*Video, 0, 10)
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".mp4") {
			videos = append(videos, &Video{file.Name(), path.Join(p, file.Name()), file.Size()})
		}
	}

	return videos, nil
}

func GetVideoFile(p string) (*Video, error) {
	fileinfo, err := os.Stat(p)
	if err != nil {
		return nil, err
	}

	if fileinfo.IsDir() {
		return nil, errors.New(p + " is a directory not a video file")
	}

	if strings.HasSuffix(fileinfo.Name(), ".mp4") {
		return &Video{fileinfo.Name(), p, fileinfo.Size()}, nil
	}
	return nil, errors.New(p + " is not a video file")
}
