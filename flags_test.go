package stamp

import (
	"flag"
	"testing"
)

// Used in examples
var s *Stamp

func init() {
	s = &Stamp{
		Revision:         "rev",
		ChangelogVersion: "1.0.2",
	}
	Use(s)
}

func ExamplePrint() {
	Print()
	//output: 1.0.2
}

func ExamplePrintDetails() {
	PrintDetails()
	//output: 1.0.2-rev
}

func TestInitFlags(t *testing.T) {
	InitFlags()
	data := []struct {
		name string
	}{
		{"v"},
		{"vv"},
	}
	for _, d := range data {
		if flag.Lookup(d.name) == nil {
			t.Errorf("InitFlags() should register -%s flag", d.name)
		}
	}
}

func TestAsFlagged(t *testing.T) {
	AsFlagged()
	exit = func(code int) {
		if code != 0 {
			t.Errorf("Should exit with code 0: %v", code)
		}
	}
	Show = true
	AsFlagged()
	Show = false
	Verbose = true
	AsFlagged()
}
