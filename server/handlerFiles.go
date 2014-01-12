package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func fileHandler(w http.ResponseWriter, r *http.Request) {
	// show files index page
	if len(r.URL.Path) <= len("/files/") {
		fmt.Fprintf(w, filesIndexHtml)
		return
	}

	// 0 files match, error
	fileName := r.URL.Path[len("/files/"):]
	if !store.Contains(fileName) {
		http.NotFound(w, r)
		return
	}

	// 1 file matches, serve it
	if store.Count(fileName) == 1 {
		ServeFile(w, r, fileName, 0)
		return
	}

	// multiple files match, show all
	_, idExists := r.URL.Query()["id"]
	if !idExists {
		if dupIndexHtml[fileName] == "" {
			html, err := store.GetDupIndexPage(fileName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			dupIndexHtml[fileName] = html
		}
		fmt.Fprintf(w, dupIndexHtml[fileName])
		return
	}

	// multiple files match & id specified, serve it
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	ServeFile(w, r, fileName, id)
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
	fileinfo, err := f.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.URL.Query().Get("dl") == "1" {
		w.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, file.Name))
	}
	http.ServeContent(w, r, fileinfo.Name(), fileinfo.ModTime(), f)
}
