package server

import (
	"fmt"
	"github.com/ultimateanu/sesame-server/filesystem"
	"net/http"
	"strconv"
)

var html string

func ServeVideos(videos []*filesystem.Video, tempdir string, port int) error {
	html = GenerateMainPage(videos, tempdir)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/files/", fileHandler)
	http.HandleFunc("/"+tempdir+"/", fileHandler)

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	return err
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr)
	fmt.Fprintf(w, html)
}

func videoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
