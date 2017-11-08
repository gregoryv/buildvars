package stamp

import (
	"fmt"
	"os/exec"
	"strings"
	"text/template"
)

var (
	t *template.Template
)

func init() {
	t = template.Must(template.New("main").Parse(
		`package {{.Package}}

const (
	Revision = "{{.Revision}}"
)
`))
}

type Stamp struct {
	Package  string
	Revision string
}

func GoTemplate() *template.Template {
	return t
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
