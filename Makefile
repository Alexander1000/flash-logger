SOURCE_DIR=./src/flash-logger
GOPACKAGES=$(shell find $(SOURCE_DIR) -name '*.go' -not -path "$(SOURCE_DIR)/vendor/*" -exec dirname {} \; | uniq)

build:
	GOPATH=$(shell pwd) go build -o bin/flash-logger flash-logger/cmd/main

test:
	GOPATH=$(shell pwd) go test -v $(GOPACKAGES)
