package server

import (
	"bytes"
	"github.com/ultimateanu/sesame-server/filesystem"
	"html/template"
)

const (
	anuTemplate = `
{{range $page, $name := .Pages}}
    <li><a href="{{$name}}/{{$page}}">{{$page}}</a></li>
{{end}}
`
	anuTemplateSimpleWorks = `
{{range $fileName, $files := .NameMap}}
    {{if $files | multiple}}
        {{range $i, $f := $files}}
            <a href="../dupfiles/{{$i}}/{{$fileName}}">{{$fileName}}</a>
        {{end}}
    {{else}}
        <a href="/{{$fileName}}">{{$fileName}}</a>
    {{end}}
{{end}}
`
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

/*

      <li><a href="/dupfiles/{{$i}}/{{$fileName}}">{{$fileName}}</a></li>{{end}}{{else}}
      <li><a href="/files/{{$fileName}}">{{$fileName}}</a></li>{{end}}{{end}}

      <li><a href="/dupfiles/{{$i}}/{{$fileName}}">{{$fileName}}</a></li>{{end}}{{else}}
      <li><a href="/files/{{$fileName}}">{{$fileName}}</a></li>{{end}}{{end}}

   {{range $i, $file := $files}}
    <a href="/{{$fileName}}">{{$fileName}}</a>
   {{end}}


   <li><a href="{{$name}}">{{len $file}}Bob</a></li>
   {{range $i, $f := $file}}
       <i>{{$i}} - {{$f.Path}}</i>
   {{end}}

*/
)

func (s *Store) GetFilesIndexPage() string {
	multipleEntries := func(files []*filesystem.File) bool {
		if len(files) > 1 {
			return true
		}
		return false
	}

	t := template.New("hello template")
	template.Must(t.Funcs(template.FuncMap{"multiple": multipleEntries}).Parse(anuTemplateSimple))

	//t, _ = t.Parse(anuTemplateSimple)
	var htmlIndex bytes.Buffer
	//TODO error
	//t.Execute(os.Stdout, store)
	t.Execute(&htmlIndex, store)

	return htmlIndex.String()
}
