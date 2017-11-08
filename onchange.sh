set -e
GOPATH=$HOME
go generate ./...
go test -cover -coverprofile /tmp/c.out .
BROWSER=firefox go tool cover -html /tmp/c.out
