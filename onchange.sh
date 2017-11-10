#!/bin/bash -e
GOPATH=$HOME
go generate ./...
go test -cover -coverprofile /tmp/c.out . > /tmp/test.out
grep "100.0%" /tmp/test.out || cat /tmp/test.out; BROWSER=surf go tool cover -html /tmp/c.out
