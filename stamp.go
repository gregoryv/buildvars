// Package stamp parses build information from git repository
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
	Package          string
	Revision         string
	ChangelogVersion string
}

func NewStamp() *Stamp {
	return &Stamp{
		Package: "main",
		Revision: "unknown",
		ChangelogVersion: "unknown",
	}
}

// NewGoTemplate returns a go source template
func NewGoTemplate() *template.Template {
	return goSource
}

// Revision returns the short revision for HEAD
func Revision(repoRoot string) (string, error) {
	revision, err := exec.Command("git", "-C", repoRoot, "rev-parse", "--short", "HEAD").CombinedOutput()
	if err != nil {
		return "unknown", fmt.Errorf("%s: %s", revision, err)
	}
	return strings.TrimSpace(string(revision)), nil
}
