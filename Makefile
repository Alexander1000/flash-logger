build:
	GOPATH=$(shell pwd) go build -o bin/flash-logger flash-logger/cmd/main
