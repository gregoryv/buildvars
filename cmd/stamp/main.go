//go:generate go install github.com/gregoryv/stamp/cmd/stamp
//go:generate stamp -go stamp.go
package main

import (
	"github.com/gregoryv/stamp"
	"flag"
	"log"
	"os"
)

var (
	out = ""
)

func init() {
	flag.StringVar(&out, "go", out, "Write Go file, defaults to stdout")
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
	var m *stamp.Stamp
	if m, err = stamp.NewStamp(); err != nil {
		er.Fatalf("Failed to generate build: %s", err)
	}
	if err = stamp.GoTemplate().Execute(fh, m); err != nil {
		er.Fatalf("Failed to write go source: %s", err)
	}
}
