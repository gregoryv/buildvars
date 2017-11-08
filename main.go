package main

import (
	"os"
	"os/exec"
	"text/template"
	"strings"
)

var t *template.Template

func init() {
	var err error
	t, err = template.New("main").Parse(
	`package {{.Package}}

const (
	Revision = "{{.Revision}}"
)
`)
	if err != nil {
		panic(err)
	}
}

type Build struct {
	Package string
	Revision string
}

func main() {
	revision, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		print(err.Error())
		os.Exit(1)
	}
	m := Build{
		Package: "main",
		Revision: strings.TrimSpace(string(revision)),
	}
	err = t.Execute(os.Stdout, m)
	if err != nil {
		panic(err)
	}
}
