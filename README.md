# README

go-buildvars parses out build information to embed into your binary

[![Build Status](https://travis-ci.org/gregoryv/buildvars.svg?branch=master)](https://travis-ci.org/gregoryv/buildvars)

## Usage

Add to one of your main files

    //go:generate go install github.com/gregoryv/buildvars
	//go:generate buildvars -o buildvars.go
    package main
	...
	
Then generate with

    go generate .
	go build .
	

