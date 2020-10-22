#!/bin/bash -e
path=$1
dir=$(dirname "$path")
filename=$(basename "$path")
extension="${filename##*.}"
nameonly="${filename%.*}"

case $extension in
    go)
        goimports -w $path
        ;;
esac

go generate ./...
go install github.com/gregoryv/stamp/cmd/stamp
go test -coverprofile /tmp/stamp.out .
uncover /tmp/stamp.out

