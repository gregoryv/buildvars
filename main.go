package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

var (
	t   *template.Template
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

type Stamp struct {
	Package  string
	Revision string
}

func NewStamp() (build *Stamp, err error) {
	var revision []byte
	revision, err = exec.Command("git", "rev-parse", "HEAD").CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s: %s", revision, err)
	}
	build = &Stamp{
		Package:  "main",
		Revision: strings.TrimSpace(string(revision)),
	}
	return
}

func Generate(out io.Writer) error {
	m, err := NewStamp()
	if err != nil {
		return err
	}
	err = t.Execute(out, m)
	return err
}

func main() {
	flag.Parse()
	fh := os.Stdout
	var err error
	er := log.New(os.Stderr, "", 0)

	if out != "" {
		fh, err = os.Create(out)
		if err != nil {
			er.Printf("Failed to create %q: %s", out, err)
			os.Exit(1)
		}
		defer fh.Close()
	}

	if err := Generate(fh); err != nil {
		er.Printf("Failed to generate build: %s", err)
		os.Exit(1)
	}
}
