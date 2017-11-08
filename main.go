package main

import (
	"os"
	"os/exec"
	"text/template"
	"strings"
	"io"
)

var t *template.Template

func init() {
	t = template.Must(template.New("main").Parse(
	`package {{.Package}}

const (
	Revision = "{{.Revision}}"
)
`))
}

type Build struct {
	Package string
	Revision string
}

func Generate(out io.Writer) error {
	revision, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		return err
	}
	m := Build{
		Package: "main",
		Revision: strings.TrimSpace(string(revision)),
	}
	err = t.Execute(out, m)
	return err
}

func main() {
	if err := Generate(os.Stdout); err != nil {
		os.Exit(1)
	}
}

