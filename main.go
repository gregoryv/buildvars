//go:generate go install github.com/gregoryv/stamp
//go:generate stamp -o stamp.go
package main

import (
	"flag"
	"fmt"
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

func main() {
	flag.Parse()
	fh := os.Stdout
	var err error
	er := log.New(os.Stderr, "", 0)

	if out != "" {
		if fh, err = os.Create(out); err != nil {
			er.Fatalf("Failed to create %q: %s", out, err)
		}
		defer fh.Close()
	}
	var m *Stamp
	if m, err = NewStamp(); err != nil {
		er.Fatalf("Failed to generate build: %s", err)
	}
	if err = t.Execute(fh, m); err != nil {
		er.Fatalf("Failed to write go source: %s", err)
	}
}
