package server

import (
	"fmt"
	"github.com/ultimateanu/sesame-server/filesystem"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

var (
	homePage  string
	filesPage string
	store     *Store
	validPath = regexp.MustCompile(`^([\d]+)/(.*)$`)
)

func ServeVideos(port int, files []*filesystem.File) error {
	store = MakeStore(files)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/files/", fileHandler)
	http.HandleFunc("/dupfiles/", dupfileHandler)

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	return err
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) <= len("/files/") {
		st := store.GetFilesIndexPage()
		fmt.Fprintf(w, st)
		return
	}

	fileName := r.URL.Path[len("/files/"):]
	if !store.Contains(fileName) {
		http.NotFound(w, r)
	} else if store.Count(fileName) == 1 {
		ServeFile(w, r, fileName, 0)
	} else {
		//TODO make
		//html := store.GetFilesIndexPage(fileName)
		html := ""
		fmt.Fprintf(w, html)
	}
}

func dupfileHandler(w http.ResponseWriter, r *http.Request) {
	resource := r.URL.Path[len("/dupfiles/"):]

	strMatches := validPath.FindStringSubmatch(resource)
	if strMatches == nil {
		fmt.Println("didn't match")
		http.NotFound(w, r)
		return
	}

	index, _ := strconv.Atoi(strMatches[1])
	ServeFile(w, r, strMatches[2], index)
}

func ServeFile(w http.ResponseWriter, r *http.Request, fileName string, index int) {
	file, err := store.GetFile(fileName, index)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	f, err := os.Open(file.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileinfo, _ := f.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeContent(w, r, fileinfo.Name(), fileinfo.ModTime(), f)
}
