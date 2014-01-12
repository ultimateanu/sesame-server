package server

import (
	"bytes"
	"github.com/dustin/go-humanize"
	"html/template"
)

const (
	fileIndexTemplate = `<!DOCTYPE html>
<html>
  <head>
    <title>Sesame Server</title>
  </head>
  <body>
    <h1>Files Index</h1>
    <table border="1" cellspacing="0" cellpadding="6">
      <tr>
        <th></th>
        <th>File</th>
        <th>Size</th>
      </tr>{{range $fileName, $files := .NameMap}}{{range $i, $f := $files}}
      <tr>
        <td>
          <form action="/files/{{$fileName}}">
            <input type="hidden" name="id" value="{{$i}}" /> 
            <input type="hidden" name="dl" value="1" /> 
            <input type="submit" value="download">
          </form>
        </td>
        <td><a href="/files/{{$fileName}}?id={{$i}}">{{$f.Name}}</a></td>
        <td align=right>{{$f.Size | humanize}}</td>
      </tr>{{end}}{{end}}
    </table>
    <br><br><hr>
    <center>Sesame Server<br>© 2014 Anu Bandi. All Rights Reserved.</center>
  </body>
</html>`
)

var (
	humanizeSize = func(size int64) string {
		return humanize.Bytes(uint64(size))
	}

	t *template.Template = template.Must(template.New("index").
		Funcs(template.FuncMap{"humanize": humanizeSize}).
		Parse(fileIndexTemplate))
)

func (s *Store) GetFilesIndexPage() (string, error) {
	var htmlIndex bytes.Buffer
	err := t.Execute(&htmlIndex, store)
	if err != nil {
		return "", err
	}
	return htmlIndex.String(), nil
}
