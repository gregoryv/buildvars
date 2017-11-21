package stamp

import (
	"fmt"
	"os/exec"
	"strings"
	"text/template"
)

var (
	goSource *template.Template
)

func init() {
	goSource = template.Must(template.New("main").Parse(
		`package {{.Package}}

import "github.com/gregoryv/stamp"

func init() {
    s := &stamp.Stamp{
	    Package: "{{.Package}}",
	    Revision: "{{.Revision}}",
	    ChangelogVersion: "{{.ChangelogVersion}}",
    }
    stamp.Use(s)
}

`))
}

// Stamp collects identifying information about a software library
type Stamp struct {
	Package  string
	Revision string
	ChangelogVersion string
}

// GoTemplate returns a go source template
func GoTemplate() *template.Template {
	return goSource
}

func Parse(repoRoot string) (build *Stamp, err error) {
	var revision []byte
	// todo refactor into Revisioner interface
	revision, err = exec.Command("git", "-C", repoRoot, "rev-parse", "HEAD").CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s: %s", revision, err)
	}
	build = &Stamp{
		Package:  "main",
		Revision: strings.TrimSpace(string(revision)),
	}
	return
}
