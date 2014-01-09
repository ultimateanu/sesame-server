package main

import (
	"github.com/docopt/docopt.go"
	"log"
	"strconv"
)

const (
	usage = `
Sesame Server

Usage:
    sesame-server [options] [video audio image] <directory or file>... 

Options:
    -p --port=<p>    Port to serve on [default: 8080].
    -h --help        Show this screen.
    --version        Show version.

Help:
    $ sesame-server /Users/anu/Documents/
        Serves all files in the Documents directory

    $ sesame-server video audio /Users/anu/Downloads/
        Serves video & audio files in the Downloads directory

    $ sesame-server -p 8888 /Users/anu/Documents/report.doc /Users/anu/Documents/report.pdf
        Serves report.doc & report.pdf on port 8888

    /files/ will have a simple list of all files being served
`
)

func parseArguments() {
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
