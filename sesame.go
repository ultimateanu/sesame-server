package main

import (
	"github.com/ultimateanu/sesame-server/filesystem"
	"github.com/ultimateanu/sesame-server/server"
)

var (
	port         int
	filesAndDirs []string
	fileFilters  []filesystem.FileFilter
)

func init() {
	parseArguments()
	printLocalIp()
}

func main() {
	files, _ := filesystem.ScanDirs(filesAndDirs, fileFilters)
	server.ServeFiles(port, files)
}
