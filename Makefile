PROJECT = github.com/IQXI/go_copy

all: get linter build

build: copy_file.go
	go build copy_file.go

linter:
	golangci-lint run
get:
    go get -u $(PROJECT)