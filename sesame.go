package main

import (
	"fmt"
	"github.com/ultimateanu/sesame-server/filesystem"
	"github.com/ultimateanu/sesame-server/server"
	"log"
	"runtime"
)

var (
	port         int
	filesAndDirs []string
	fileFilters  []filesystem.FileFilter
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	parseArguments()
}

func main() {
	files, _ := filesystem.ScanDirs(filesAndDirs, fileFilters)

	localIp, err := server.GetLocalIp()
	if err != nil {
		log.Fatalln("no local ip address detected")
	}
	for _, ip := range localIp {
		fmt.Printf("Serving files at http://%s:%d\n", ip, port)
	}

	server.ServeFiles(port, files)
}
