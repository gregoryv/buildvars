//go:generate stamp -go stamp.go -clfile ../../CHANGELOG.md
package main

import (
	"flag"
	"github.com/gregoryv/stamp"
	"io/ioutil"
	"log"
	"os"
)

var (
	out    = ""
	clfile = "CHANGELOG.md"
)

func init() {
	flag.StringVar(&out, "go", out, "Write Go file, defaults to stdout")
	flag.StringVar(&clfile, "clfile", clfile, "Changelog to parse for version, keepachangelog format")
	stamp.InitFlags()
}

func main() {
	flag.Parse()
	// Use the stamp our selves
	if stamp.Show {
		stamp.Print()
		os.Exit(0)
	}
	if stamp.Verbose {
		stamp.PrintDetails()
		os.Exit(0)
	}

	var err error
	fh := os.Stdout
	// Errors are written to stderr
	er := log.New(os.Stderr, "", 0)

	if out != "" {
		if fh, err = os.Create(out); err != nil {
			er.Fatalf("Failed to create %q: %s", out, err)
		}
		defer fh.Close()
	}
	// Create initial stamp by parsing repository for current revision
	var m *stamp.Stamp
	if m, err = stamp.Parse("."); err != nil {
		er.Fatalf("Failed to generate build: %s", err)
	}
	// Set version from changelog
	content, err := ioutil.ReadFile(clfile)
	if err != nil {
		er.Fatal(err)
	}
	changelog := stamp.NewChangelog(content)
	m.ChangelogVersion, err = changelog.Version()
	if err != nil {
		er.Fatal(err)
	}

	// Write go code
	if err = stamp.GoTemplate().Execute(fh, m); err != nil {
		er.Fatalf("Failed to write go source: %s", err)
	}
}
