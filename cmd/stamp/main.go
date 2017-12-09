//Command stamp generates go source code with build information.
//
//go:generate stamp -go build_stamp.go -clfile ../../CHANGELOG.md
package main

import (
	"flag"
	"fmt"
	"github.com/gregoryv/stamp"
	"log"
	"os"
)

var (
	out    = ""
	clfile = "CHANGELOG.md"
	help   = false
)

func init() {
	flag.BoolVar(&help, "h", help, "Print this help and exit")
	flag.StringVar(&out, "go", out, "Write Go file, defaults to stdout")
	flag.StringVar(&clfile, "clfile", clfile,
		"Changelog to parse for version, keepachangelog format")
	stamp.InitFlags()
}

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	stamp.AsFlagged()

	if help {
		Usage()
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
	m := stamp.NewStamp()
	if m.Revision, err = stamp.Revision("."); err != nil {
		er.Fatalf("Failed to generate build: %s", err)
	}
	// Set version from changelog
	m.ParseChangelog(clfile)

	// Write go code
	if err = stamp.NewGoTemplate().Execute(fh, m); err != nil {
		er.Fatalf("Failed to write go source: %s", err)
	}
}
