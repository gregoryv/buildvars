#!/bin/bash
GOPATH=$HOME
go generate ./...
go test -cover -coverprofile /tmp/c.out .
go install github.com/gregoryv/stamp/cmd/stamp
#cat cmd/stamp/stamp.go
#rm cmd/stamp/stamp.go
