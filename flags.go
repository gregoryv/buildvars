package stamp

import (
	"flag"
	"fmt"
)

var (
	Show    = false
	Verbose = false
	sp      *Stamp
)

func init() {
	sp = &Stamp{}
}

func Use(s *Stamp) {
	sp = s
}

func InitFlags() {
	flag.BoolVar(&Show, "v", Show, "Print version and exit")
	flag.BoolVar(&Verbose, "vv", Verbose, "Print version with details and exit")
}

func Print() {
	fmt.Print(sp.ChangelogVersion)
}

func PrintDetails() {
	fmt.Printf("%s %s", sp.ChangelogVersion, sp.Revision)
}
