#!/bin/bash
GOPATH=$HOME
go generate ./...
go install github.com/gregoryv/stamp/cmd/stamp
go test -coverprofile /tmp/stamp.out .
uncover /tmp/stamp.out

