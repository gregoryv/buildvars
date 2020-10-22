// Package stamp parses build information from git repository and
// changelog.md
package stamp

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
)

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

// Registers -v and -vv flags on the given flagset
func (me *Stamp) InitFlags(fs *flag.FlagSet) {
	fs.BoolVar(&me.Show, "v", me.Show, "Print version and exit")
	fs.BoolVar(&me.Verbose, "vv", me.Verbose,
		"Print version with details and exit",
	)
}

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
func (me *Stamp) ParseChangelog(file string) (err error) {
	var content []byte
	if content, err = ioutil.ReadFile(file); err != nil {
		return
	}
	changelog := NewChangelog(content)
	if me.ChangelogVersion, err = changelog.Version(); err != nil {
		return
	}
	return
}
