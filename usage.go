package main

import (
	"fmt"
	"github.com/docopt/docopt.go"
	"github.com/dustin/go-humanize"
	"github.com/ultimateanu/sesame-server/filesystem"
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
    -m --min=<m>     Minimum file size (ex: 50b, 10kib)
    -M --max=<M>     Maximum file size (ex: 100mb, 4gb)
    -s --sysfiles    Include hidden files.
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

	fmt.Println(arguments)

	port, err = strconv.Atoi(arguments["--port"].(string))
	if err != nil {
		log.Fatalln("Please specify a valid port")
	}

	fileFilters = append(fileFilters, filesystem.AllFiles)

	if !arguments["--sysfiles"].(bool) {
		fileFilters = append(fileFilters, filesystem.IgnoreSystemFiles)
	}

	if arguments["--min"] != nil {
		minSize, err := humanize.ParseBytes(arguments["--min"].(string))
		if err != nil {
			log.Fatalln("Please enter a valid minimum file size. (ex: --min 2mb)")
		}
		fileFilters = append(fileFilters, filesystem.MinFilter(minSize))
	}

	if arguments["--max"] != nil {
		maxSize, err := humanize.ParseBytes(arguments["--max"].(string))
		if err != nil {
			log.Fatalln("Please enter a valid maximum file size. (ex: --max 4gb)")
		}
		fileFilters = append(fileFilters, filesystem.MaxFilter(maxSize))
	}

	videoFiles = arguments["video"].(bool)
	audioFiles = arguments["audio"].(bool)
	imageFiles = arguments["image"].(bool)

	filesAndDirs = arguments["<directory or file>"].([]string)
}
