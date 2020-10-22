package stamp

import (
	"flag"
	"os"
)

var (
	DefaultStamp = &Stamp{}
	exit         = os.Exit
)

// Regiters -v and -vv flags
func InitFlags() {
	DefaultStamp.InitFlags(flag.CommandLine)
}

func Print() {
	DefaultStamp.WriteTo(os.Stdout)
}

func PrintDetails() {
	DefaultStamp.WriteTo(os.Stdout)
}

// AsFlagged shows information according to flags and exits with code 0
func AsFlagged() {
	if DefaultStamp.Show || DefaultStamp.Verbose {
		DefaultStamp.WriteTo(os.Stdout)
		exit(0)
	}
}
