set -e
GOPATH=$HOME
go generate ./...
go test -cover -coverprofile /tmp/c.out .
BROWSER=surf go tool cover -html /tmp/c.out
