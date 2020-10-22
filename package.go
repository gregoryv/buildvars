package stamp

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

var (
	DefaultStamp = &Stamp{}
	exit         = os.Exit
	goSource     *template.Template
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

// Regiters -v and -vv flags on default stamp
func InitFlags() {
	DefaultStamp.InitFlags(flag.CommandLine)
}

func Print() {
	DefaultStamp.WriteTo(os.Stdout)
}

func PrintDetails() {
	DefaultStamp.WriteTo(os.Stdout)
}

// AsFlagged shows information according to flags and exits with code 0
func AsFlagged() {
	if DefaultStamp.Show || DefaultStamp.Verbose {
		DefaultStamp.WriteTo(os.Stdout)
		exit(0)
	}
}
