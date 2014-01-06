package main

import (
	"fmt"
	"github.com/docopt/docopt.go"
	"github.com/ultimateanu/sesame-server/filesystem"
	"github.com/ultimateanu/sesame-server/server"
	"io/ioutil"
	"log"
	"strconv"
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

func parseArguments() {
	usage := `Sesame Server

Usage:
    sesame-server [--port=<p>] <directory or file>... 

Options:
    -h --help        Show this screen.
    -p --port=<p>    Port to serve on [default: 8080].
    --version        Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, "Sesame Server 0.1", true)
	port, err = strconv.Atoi(arguments["--port"].(string))

	if err != nil {
		log.Fatalln("Please specify a valid port")
	}
	filesAndDirs = arguments["<directory or file>"].([]string)
}
