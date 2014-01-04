package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	panic(http.ListenAndServe(":8484", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr)
	fmt.Fprintf(w, `<html><head></head><body><h1>Bommarillu</h1><center>
  <video width="720" height="480" poster="static/bom.jpg" controls><source src="static/bom.mp4" type="video/mp4">Your browser does not support the video tag.</video>
  <br><br>
  
  </center></body></html>`)
}
