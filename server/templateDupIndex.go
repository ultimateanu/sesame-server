package server

import (
	"bytes"
	"html/template"
)

const (
	dupIndexTemplate = `<!DOCTYPE html>
<html>
  <head>
    <title>Sesame Server</title>
  </head>
  <body>
    <h1>Multiple Files</h1>
    <table border="1" cellspacing="0" cellpadding="4">
      <tr>
        <th></th>
        <th>File</th>
        <th>Size</th>
      </tr>{{range $i, $f := .}}
      <tr>
        <td>
          <form action="/files/{{$f.Name | urlsafe}}">
            <input type="hidden" name="id" value="{{$i}}" /> 
            <input type="hidden" name="dl" value="1" /> 
            <input type="submit" value="download">
          </form>
        </td>
        <td><a href="/files/{{$f.Name | urlsafe}}?id={{$i}}">{{$f.Name}}</a></td>
        <td align=right>{{$f.Size | humanize}}</td>
      </tr>{{end}}
    </table>
    <br><br><hr>
    <center>Sesame Server<br>© 2014 Anu Bandi. All Rights Reserved.</center>
  </body>
</html>`
)

var (
	tee *template.Template = template.Must(template.New("index").
		Funcs(template.FuncMap{"urlsafe": UrlSafe, "humanize": humanizeSize}).
		Parse(dupIndexTemplate))
)

func (s *Store) GetDupIndexPage(fileName string) (string, error) {
	var htmlIndex bytes.Buffer
	err := tee.Execute(&htmlIndex, store.Get(fileName))
	if err != nil {
		return "", err
	}
	return htmlIndex.String(), nil
}
