package main

import (
	"github.com/docopt/docopt.go"
	"log"
	"strconv"
)

const (
	usage = `Sesame Server

Usage:
    sesame-server [options] [video audio image] <directory or file>... 

Options:
    -p --port=<p>    Port to serve on [default: 8080].
    -h --help        Show this screen.
    --version        Show version.`
)

func parseArguments() {
	/*
			usage := `Sesame Server

		Usage:
		    sesame-server [options] [video audio image] <directory or file>...

		Options:
		    -p --port=<p>    Port to serve on [default: 8080].
		    -h --help        Show this screen.
		    --version        Show version.`

	*/

	arguments, _ := docopt.Parse(usage, nil, true, "Sesame Server 0.1", true)

	port, err = strconv.Atoi(arguments["--port"].(string))
	if err != nil {
		log.Fatalln("Please specify a valid port")
	}

	videoFiles = arguments["video"].(bool)
	audioFiles = arguments["audio"].(bool)
	imageFiles = arguments["image"].(bool)

	filesAndDirs = arguments["<directory or file>"].([]string)
}
