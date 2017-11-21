# README

[stamp](https://godoc.org/github.com/gregoryv/stamp) - Parses out build information to embed into your binary

[![Build Status](https://travis-ci.org/gregoryv/stamp.svg?branch=master)](https://travis-ci.org/gregoryv/stamp)

Normalize how version and build information makes it's way into your binaries.
Generates code that can be used to add flags

    -v    Print version and exit
    -vv
          Print version with details and exit

## Usage

Install

    go get github.com/gregoryv/stamp/...

Example main.go

	//go:generate stamp -go build_stamp.go -clfile changelog.md
    package main

	import (
		"github.com/gregoryv/stamp"
		"flag"
	)

	func init() {
		// Add -v and -vv flags
		stamp.InitFlags()
	}

	func main() {
		flag.Parse()
		stamp.AsFlagged()
		//...
	}


Then generate with

    go generate .
	go build .

## Details

stamp depends on git and that you have a CHANGELOG.md. The changelog is parsed for the lates
released version and assumes it follows http://keepachangelog.com/en/1.0.0/ format.
