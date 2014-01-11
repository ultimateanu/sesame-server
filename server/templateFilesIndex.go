package server

import (
	"bytes"
	"github.com/ultimateanu/sesame-server/filesystem"
	"html/template"
)

const (
	anuTemplateSimple = `<!DOCTYPE html>
<html>
  <head>
    <title>Sesame Server Files</title>
  </head>
  <body>
    <h1>Files Index</h1>
    <ul>{{range $fileName, $files := .NameMap}}{{if $files | multiple}}{{range $i, $f := $files}}
      <li><a href="/dupfiles/{{$i}}/{{$fileName}}">{{$f.Name}}</a></li>{{end}}{{else}}{{range $i, $f := $files}}
      <li><a href="/files/{{$fileName}}">{{$f.Name}}</a></li>{{end}}{{end}}{{end}}
    </ul>
  </body>
</html>`
)

var (
	t *template.Template = template.Must(template.New("index").
		Funcs(template.FuncMap{"multiple": multipleEntries}).
		Parse(anuTemplateSimple))

	multipleEntries = func(files []*filesystem.File) bool {
		if len(files) > 1 {
			return true
		}
		return false
	}
)

func (s *Store) GetFilesIndexPage() (string, error) {
	var htmlIndex bytes.Buffer
	err := t.Execute(&htmlIndex, store)
	if err != nil {
		return "", err
	}
	return htmlIndex.String(), nil
}
