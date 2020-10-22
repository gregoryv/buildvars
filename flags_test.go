package stamp

import (
	"flag"
	"os"
	"testing"
)

func ExamplePrint() {
	DefaultStamp = &Stamp{
		Revision:         "rev",
		ChangelogVersion: "1.0.2",
		Show:             true,
	}
	Print()
	//output: 1.0.2
}

func ExamplePrintDetails() {
	DefaultStamp = &Stamp{
		Revision:         "6e6046c",
		ChangelogVersion: "1.0.0",
		Verbose:          true,
	}
	PrintDetails()
	//output: 1.0.0-6e6046c
}

func ExampleStamp_WriteTo() {
	s := &Stamp{
		Revision:         "6e6046c",
		ChangelogVersion: "1.0.0",
		Verbose:          true,
	}
	s.WriteTo(os.Stdout)
	//output: 1.0.0-6e6046c
}

func ExampleStamp_WriteTo_noOutput() {
	s := &Stamp{
		Revision:         "6e6046c",
		ChangelogVersion: "1.0.0",
	}
	s.WriteTo(os.Stdout)
	//output:
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
	AsFlagged()
	DefaultStamp.Show = true
	AsFlagged()
	DefaultStamp.Show = false
	DefaultStamp.Verbose = true
	AsFlagged()
}
