set -e
GOPATH=$HOME
go generate ./...
go test -cover .
