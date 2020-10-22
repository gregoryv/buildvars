#!/bin/bash

command=stamp
out=dist/$command

mkdir -p dist
rm -rf dist/*
go build -o $out ./cmd/$command
upx $out
cp CHANGELOG.md dist/
# Get the version from the binary
version=`$out -v`
distname=$command-$version
rm -f $distname.tgz
mv dist $distname
tar -c $distname | gzip - > $distname.tgz
rm -rf $distname
tar tvfz $distname.tgz
