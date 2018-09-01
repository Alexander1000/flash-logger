build:
	GOPATH=$(shell pwd) go build -o bin/flash-logger ./src/cmd/main
