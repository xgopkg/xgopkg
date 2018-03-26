GO_FILES=$(find . -iname '*.go' -type f | egrep -v '/vendor/')
gofmt -s -l $GO_FILES
# # go vet is the official Go static analyzer
go vet ./...
# #go test
XGP_CONFIG_DIR=$PWD/conf go test -v $(go list ./...)
gocyclo -over 19 $GO_FILES                                       # forbid code with huge functions
golint -set_exit_status $(go list ./...)                         # one last linter