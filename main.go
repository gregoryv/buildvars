package main

import (
	"os"
	"text/template"
)

var t *template.Template

func init() {
	var err error
	t, err = template.New("main").Parse(
	`package {{.Package}}

const (
	Version = "{{.Version}}"
)
`)
	if err != nil {
		panic(err)
	}
}

type Build struct {
	Package string
	Version string
}

func main() {
	m := Build{
		Package: "main",
		Version: "1.0",
	}
	err := t.Execute(os.Stdout, m)
	if err != nil {
		panic(err)
	}
}
