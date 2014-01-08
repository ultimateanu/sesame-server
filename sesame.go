package main

import (
	"fmt"
	"github.com/ultimateanu/sesame-server/filesystem"
	"github.com/ultimateanu/sesame-server/server"
	"io/ioutil"
	"log"
)

var (
	port         int
	filesAndDirs []string
	err          error
)

func main() {
	parseArguments()

	videos := make([]*filesystem.Video, 0, 10)
	for _, fileOrDir := range filesAndDirs {
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

	tempdir, err := ioutil.TempDir(".", "temp_sesame_videos_")
	if err != nil {
		log.Fatalln("cannot create temp directory:", err)
	}

	err = filesystem.SymlinkVideos(videos, tempdir)
	if err != nil {
		log.Fatalln("symlink failed:", err)
	}

	/*
		ht := server.GenerateMainPage(videos, tempdir)
		fmt.Println(ht)
	*/

	/*
		for _, v := range videos {
			fmt.Println(*v)
		}
	*/

	localIp, err := server.GetLocalIp()
	if err != nil {
		log.Fatalln("no local ip address detected")
	}
	for _, ip := range localIp {
		fmt.Printf("Serving videos at http://%s:%d\n", ip, port)
	}

	server.ServeVideos(videos, tempdir, port)
}
