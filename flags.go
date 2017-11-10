package stamp

import (
	"flag"
)

var (
	Show    = false
	Verbose = false
	sp *Stamp
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
	print(sp.Revision)
}

func PrintDetails() {
	print(sp.Revision)
}
