package main

import (
	"fmt"
	"github.com/ultimateanu/sesame-server/filesystem"
	"github.com/ultimateanu/sesame-server/server"
	"log"
)

var (
	port         int
	filesAndDirs []string
	err          error
	videoFiles   bool
	audioFiles   bool
	imageFiles   bool
	validExt     []string
	videoExt     []string = []string{"mp4", "avi", "mkv", "wmv"}
	audioExt     []string = []string{"mp3", "wma", "aac"}
	imageExt     []string = []string{"jpg", "jpeg", "png"}
)

func main() {
	parseArguments()
	files, _ := filesystem.ExtractDirs(filesAndDirs)

	if videoFiles {
		validExt = append(validExt, videoExt...)
	}
	if audioFiles {
		validExt = append(validExt, audioExt...)
	}
	if imageFiles {
		validExt = append(validExt, imageExt...)
	}

	if len(validExt) != 0 {
		files = filesystem.Filter(files, filesystem.FileExtension(validExt))
	}

	//log.Println(server.MakeStore(files))

	localIp, err := server.GetLocalIp()
	if err != nil {
		log.Fatalln("no local ip address detected")
	}
	for _, ip := range localIp {
		fmt.Printf("Serving files at http://%s:%d\n", ip, port)
	}

	server.ServeVideos(port, files)
}
