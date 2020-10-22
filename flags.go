package stamp

import (
	"flag"
	"fmt"
	"os"
)

var (
	Show         = false
	Verbose      = false
	DefaultStamp = &Stamp{}
	exit         = os.Exit
)

// Use sets the stamp to use when printing details
func Use(stamp *Stamp) {
	DefaultStamp = stamp
}

func InUse() *Stamp {
	return DefaultStamp
}

// Regiters -v and -vv flags
func InitFlags() {
	flag.BoolVar(&Show, "v", Show, "Print version and exit")
	flag.BoolVar(&Verbose, "vv", Verbose, "Print version with details and exit")
}

func Print() {
	fmt.Print(DefaultStamp.ChangelogVersion)
}

func PrintDetails() {
	fmt.Printf("%s-%s", DefaultStamp.ChangelogVersion, DefaultStamp.Revision)
}

// AsFlagged shows information according to flags and exits with code 0
func AsFlagged() {
	if Show {
		Print()
		exit(0)
	}
	if Verbose {
		PrintDetails()
		exit(0)
	}
}
