#!/bin/bash -e
GOPATH=$HOME
go generate ./...
go test -cover -coverprofile /tmp/c.out . 
go install github.com/gregoryv/stamp/cmd/stamp

