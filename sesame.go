package main

import (
	"fmt"
	"github.com/ultimateanu/sesame-server/filesystem"
	"log"
	//"net/http"
	//"os"
	"github.com/docopt/docopt.go"
)

func main() {
	usage := `Sesame Server

  Usage:
  sesame-server <directory or file>...

  Options:
  -h --help     Show this screen.
  --version     Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, "Sesame Server 0.1", false)

	videos := make([]*filesystem.Video, 0, 10)
	for _, fileOrDir := range arguments["<directory or file>"].([]string) {
		if filesystem.IsFile(fileOrDir) {
			video, err := filesystem.GetVideoFile(fileOrDir)
			if err != nil {
				log.Print(err)
			} else {
				videos = append(videos, video)
			}
		} else if filesystem.IsDir(fileOrDir) {
			vids, err := filesystem.GetAllVideoFiles(fileOrDir)
			if err != nil {
				log.Print(err)
			} else {
				videos = append(videos, vids...)
			}
		} else {
			log.Print(fileOrDir + " is invalid")
		}
	}

	for _, v := range videos {
		fmt.Println(*v)
	}
}
