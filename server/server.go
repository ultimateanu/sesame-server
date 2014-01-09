package server

import (
	"fmt"
	"github.com/ultimateanu/sesame-server/filesystem"
	"net/http"
	"os"
	"strconv"
)

var (
	homePage  string
	filesPage string
	m         *map[int]*filesystem.File
)

func ServeVideos(port int, files []*filesystem.File) error {
	//html = GenerateMainPage(videos, tempdir)

	m = MakeStore(files)
	filesPage = MakeFilesIndexPage(m)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/files/", fileHandler)

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	return err
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.RemoteAddr)
	//fmt.Fprintf(w, html)
	fmt.Fprintf(w, "hello world")
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	resource := r.URL.Path[len("/files/"):]
	if resource == "" {
		fmt.Fprintf(w, filesPage)
		return
	}

	id, err := strconv.Atoi(resource)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	file := (*m)[id]
	if file == nil {
		http.NotFound(w, r)
		return
	}

	f, _ := os.Open(file.Path)
	fileinfo, _ := f.Stat()
	http.ServeContent(w, r, fileinfo.Name(), fileinfo.ModTime(), f)
}
