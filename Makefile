SHELL := /bin/bash

code/check:
	@gofmt -l `find . -type f -name '*.go' -not -path "./vendor/*"`

code/fmt:
	@gofmt -w `find . -type f -name '*.go' -not -path "./vendor/*"`

.PHONY: code/fmt code

test/clean:
	@go clean -testcache

test/unit:
	@go test -v ./kojiclient/.../

test/e2e:
	@go test -v ./test/e2e/.../

test: test/clean test/unit test/e2e

.PHONY: test/clean test/unit test/e2e test

build/clean:
	@rm -rf _build

build/prepare:
	@mkdir -p _build/bin

build/koji-controller: build/prepare
	@mkdir -p _build/bin
	@go build -v -o _build/bin/koji-controller koji-controller/main.go

build: build/prepare build/koji-controller

.PHONY: build/clean build/prepare build/koji-k8s-controller build
