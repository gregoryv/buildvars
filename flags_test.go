package stamp

import (
	"flag"
	"testing"
)

func Test_InUse(t *testing.T) {
	s := InUse()
	if s == nil {
		t.Error("InUse should always return a Stamp, got nil")
	}
}

func ExamplePrint() {
	s := &Stamp{
		Revision:         "rev",
		ChangelogVersion: "1.0.2",
		Show:             true,
	}
	Use(s)
	Print()
	//output: 1.0.2
}

func ExamplePrintDetails() {
	s := &Stamp{
		Revision:         "6e6046c",
		ChangelogVersion: "1.0.0",
		Verbose:          true,
	}
	Use(s)
	PrintDetails()
	//output: 1.0.0-6e6046c
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
