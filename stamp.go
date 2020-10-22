// Package stamp parses build information from git repository and
// changelog.md
package stamp

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
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
	stamp.DefaultStamp = &stamp.Stamp{
		Package:          "{{.Package}}",
		Revision:         "{{.Revision}}",
		ChangelogVersion: "{{.ChangelogVersion}}",
	}
}
`))
}

func NewStamp() *Stamp {
	return &Stamp{
		Package:          "main",
		Revision:         "unknown",
		ChangelogVersion: "unknown",
	}
}

// Stamp collects identifying information about a software library
type Stamp struct {
	Package          string
	Revision         string
	ChangelogVersion string

	Show    bool
	Verbose bool
}

// Regiters -v and -vv flags
func (me *Stamp) InitFlags(fs *flag.FlagSet) {
	fs.BoolVar(&me.Show, "v", me.Show, "Print version and exit")
	fs.BoolVar(&me.Verbose, "vv", me.Verbose,
		"Print version with details and exit",
	)
}

// WriteTo
func (me *Stamp) WriteTo(w io.Writer) (int64, error) {
	switch {
	case me.Show:
		n, err := fmt.Fprint(w, me.ChangelogVersion)
		return int64(n), err
	case me.Verbose:
		n, err := fmt.Fprintf(w, "%s-%s", me.ChangelogVersion, me.Revision)
		return int64(n), err
	}
	return 0, nil
}

// ParseChangelog sets ChangelogVersion of this stamp from the given file
func (s *Stamp) ParseChangelog(file string) (err error) {
	var content []byte
	if content, err = ioutil.ReadFile(file); err != nil {
		return
	}
	changelog := NewChangelog(content)
	if s.ChangelogVersion, err = changelog.Version(); err != nil {
		return
	}
	return
}

// NewGoTemplate returns a go source template
func NewGoTemplate() *template.Template {
	return goSource
}

// Revision returns the short revision for HEAD
func Revision(repoRoot string) (string, error) {
	revision, err := exec.Command("git", "-C", repoRoot, "rev-parse", "--short",
		"HEAD").CombinedOutput()
	if err != nil {
		return "unknown", fmt.Errorf("%s: %s", revision, err)
	}
	return strings.TrimSpace(string(revision)), nil
}
