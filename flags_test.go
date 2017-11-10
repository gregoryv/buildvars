package stamp

import (
	"flag"
	"testing"
)

func ExamplePrint() {
	s := &Stamp{
		Revision: "x",
	}
	Use(s)
	Print()
	//output: x
}

func ExamplePrintDetails() {
	s := &Stamp{
		Revision: "x",
	}
	Use(s)
	PrintDetails()
	//output: x
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
