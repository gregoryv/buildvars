# README

stamp parses out build information to embed into your binary

[![Build Status](https://travis-ci.org/gregoryv/stamp.svg?branch=master)](https://travis-ci.org/gregoryv/stamp)

## Usage

Example main.go

    //go:generate go install github.com/gregoryv/stamp/cmd/stamp
	//go:generate stamp -go stamp.go
    package main
	
	import (
		"github.com/gregoryv/stamp"
		"flag"
	)
	
	func init() {
	    stamp.InitFlags()
	}
	
	func main() {
		flag.Parse()
	    if stamp.Show {
	        stamp.Print()
		    os.Exit(0)
		}
		if stamp.Verbose {
			stamp.PrintDetails()
			os.Exit(0)
		}
		//...
	}
		

Then generate with

    go generate .
	go build .
