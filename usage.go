package main

import (
	"github.com/docopt/docopt.go"
	"log"
	"strconv"
)

    //sesame-server [--port=<p>] [videos music photos] <directory or file>... 
func parseArguments() {
	usage := `Sesame Server

Usage:
    sesame-server [options] [videos music photos] <directory or file>... 

Options:
    -p --port=<p>    Port to serve on [default: 8080].
    -h --help        Show this screen.
    --version        Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, "Sesame Server 0.1", true)
	port, err = strconv.Atoi(arguments["--port"].(string))

	log.Println(arguments)

	if err != nil {
		log.Fatalln("Please specify a valid port")
	}
	filesAndDirs = arguments["<directory or file>"].([]string)
}
