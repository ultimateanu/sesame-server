package server

import (
	"fmt"
	"github.com/ultimateanu/sesame-server/filesystem"
	"net/http"
	"regexp"
	"strconv"
)

var (
	store          *Store
	filesIndexHtml string
	dupIndexHtml   map[string]string = make(map[string]string)
	validDupPath                     = regexp.MustCompile(`^([\d]+)/(.*)$`)
)

func ServeFiles(port int, files []*filesystem.File) (err error) {
	store = MakeStore(files)
	filesIndexHtml, err = store.GetFilesIndexPage()
	if err != nil {
		return err
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/files/", fileHandler)

	err = http.ListenAndServe(":"+strconv.Itoa(port), nil)
	return err
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}
