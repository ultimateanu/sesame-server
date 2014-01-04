package filesystem

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func GetAllVideoFiles(p string) ([]Video, error) {
	//fmt.Println("Trying to get all in :", p)
	files, err := ioutil.ReadDir(p)
	if err != nil {
		return nil, err
	}

	videos, err := GetVideoFiles(p)
	if err != nil {
		return nil, err
	}
	//fmt.Println("Initial: ", videos)

	for _, file := range files {
		if file.IsDir() {
			//fmt.Println("Getting : ", file.Name())
			v, err := GetAllVideoFiles(path.Join(p, file.Name()))
			if err != nil {
				return nil, err
			}
			//fmt.Println("GOT : ", v)
			videos = append(videos, v...)
		}
	}

	return videos, nil
}

func GetVideoFiles(p string) ([]Video, error) {
	files, err := ioutil.ReadDir(p)
	if err != nil {
		return nil, err
	}

	videos := make([]Video, 0, 10)
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".mp4") {
			videos = append(videos, Video{file.Name(), path.Join(p, file.Name()), file.Size()})
		}
	}

	return videos, nil
}

func VideoFilter(path string, info os.FileInfo, err error) error {
	fmt.Println("Walked at " + path)
	return nil
}
