#!/bin/bash
GOPATH=$HOME
go generate ./...
go install github.com/gregoryv/stamp/cmd/stamp
go test -coverprofile /tmp/c.out .
uncover /tmp/c.out

#cat cmd/stamp/stamp.go
#rm cmd/stamp/stamp.go
