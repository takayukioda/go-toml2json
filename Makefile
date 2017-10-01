NAME := go-toml2json
REVISION := $(shell git rev-parse --short HEAD)

FILES := $$(find . -name '*.go' | grep -v vendor)

default: lint

# Install required tools
# dep: dependence verify tool
# golint: lint tool for go
setup:
	go get github.com/golang/dep/cmd/dep
	go get github.com/golang/lint/golint

# Install dependencies by dep
deps: setup
	dep ensure

lint: setup
	go vet $(FILES)
	golint $(FILES)

test: setup lint
	go test

fmt:
	go fmt $(FILES)

build: deps lint fmt
	go build

run: build
	./$(NAME)

.PHONY: setup deps lint fmt build run
