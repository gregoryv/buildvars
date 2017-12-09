package stamp

import (
	"flag"
	"testing"
)

func ExamplePrint() {
	s := &Stamp{
		Revision:         "rev",
		ChangelogVersion: "1.0.0",
	}
	Use(s)
	Print()
	//output: 1.0.0
}

func ExamplePrintDetails() {
	s := &Stamp{
		Revision:         "rev",
		ChangelogVersion: "1.0.0",
	}
	Use(s)
	PrintDetails()
	//output: 1.0.0-rev
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