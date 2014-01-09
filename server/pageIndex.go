package server

import (
	"fmt"
	"github.com/dustin/go-humanize"
)

const (
	indexHtmlStart = `<!DOCTYPE html>
<html>
  <head>
    <title>Sesame Server Files</title>
  </head>
  <body>
    <h1>Files</h1>
    <ul>`
	indexHtmlEnd = `
    </ul>
  </body>
</html>`
)

func MakeFilesIndexPage(s Store) string {
	indexPage := indexHtmlStart
	for i, f := range *s {
		indexPage += fmt.Sprintf("\n\t<li>%s [<a href=\"%d\" download=\"%s\">%s</a>]</li>",
			Escape(f.Name), i, Escape(f.Name), Escape(humanize.Bytes(uint64(f.Size))))
	}
	indexPage += indexHtmlEnd

	return indexPage
}
