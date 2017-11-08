# README

stamp parses out build information to embed into your binary

[![Build Status](https://travis-ci.org/gregoryv/stamp.svg?branch=master)](https://travis-ci.org/gregoryv/stamp)

## Usage

Add to one of your main files

    //go:generate go install github.com/gregoryv/stamp/cmd/stamp
	//go:generate stamp -o stamp.go
    package main
	...

Then generate with

    go generate .
	go build .
