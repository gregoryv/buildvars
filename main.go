package main

import (
	"io"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"flag"
)

var (
	t *template.Template
	out = ""
)

func init() {
	t = template.Must(template.New("main").Parse(
		`package {{.Package}}

const (
	Revision = "{{.Revision}}"
)
`))
	flag.StringVar(&out, "o", out, "Write to file, defaults to stdout")
}

type Build struct {
	Package  string
	Revision string
}

func Generate(out io.Writer) error {
	revision, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		return err
	}
	m := Build{
		Package:  "main",
		Revision: strings.TrimSpace(string(revision)),
	}
	err = t.Execute(out, m)
	return err
}

func main() {
	flag.Parse()
	fh := os.Stdout
	var err error
	if out != "" {
		fh, err = os.Create(out)
		if err != nil {
			print(err.Error())
			os.Exit(1)
		}
		defer fh.Close()
	}
	
	if err := Generate(fh); err != nil {
		os.Exit(1)
	}
}
