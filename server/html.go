package server

import (
	"github.com/ultimateanu/sesame-server/filesystem"
	"path/filepath"
)

const (
	htmlStart = "<html><head><center><h1>Sesame Server</h1></center></head><body><center>"
	htmlEnd   = "</center></body></html"
)

func GenerateMainPage(videos []*filesystem.File, dir string) (html string) {
	html = htmlStart
	for _, video := range videos {
		html += `<video width="720" height="480" preload="metadata" controls><source src="`
		html += filepath.Join(dir, video.Name)
		html += `" type="video/mp4">Your browser does not support the video tag.</video>`
		//html += "\n<h2>" + video.Name + "</h2>"
		html += "\n" + `<h2><a href="` + dir + "/" + video.Name + `">` + video.Name + "</a>"

		html += "<br><br><br><hr>"
	}
	html += "\n" + htmlEnd
	return
}
