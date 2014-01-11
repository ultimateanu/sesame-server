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
	http.HandleFunc("/dupfiles/", dupfileHandler)

	err = http.ListenAndServe(":"+strconv.Itoa(port), nil)
	return err
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) <= len("/files/") {
		fmt.Fprintf(w, filesIndexHtml)
		return
	}

	fileName := r.URL.Path[len("/files/"):]
	if !store.Contains(fileName) {
		http.NotFound(w, r)
	} else if store.Count(fileName) == 1 {
		ServeFile(w, r, fileName, 0)
	} else {
		//Lazy loading
		if dupIndexHtml[fileName] == "" {
			html, err := store.GetDupIndexPage(fileName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			dupIndexHtml[fileName] = html
		}

		fmt.Fprintf(w, dupIndexHtml[fileName])
	}
}

func dupfileHandler(w http.ResponseWriter, r *http.Request) {
	resource := r.URL.Path[len("/dupfiles/"):]
	strMatches := validDupPath.FindStringSubmatch(resource)
	if strMatches == nil {
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
	fileinfo, err := f.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeContent(w, r, fileinfo.Name(), fileinfo.ModTime(), f)
}

//TODO: add content-disposition header to start download with proper name
