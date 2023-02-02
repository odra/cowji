SHELL := /bin/bash

code/check:
	@gofmt -l `find . -type f -name '*.go' -not -path "./vendor/*"`

code/fmt:
	@gofmt -w `find . -type f -name '*.go' -not -path "./vendor/*"`

.PHONY: code/fmt code

test/clean:
	@go clean -testcache

test/unit/kojiclient:
	@go test -v ./kojiclient/.../

test/unit/koji-controller:
	@go test -v ./koji-controller/.../

test/unit: test/unit/kojiclient test/unit/koji-controller

test/e2e:
	@go test -v ./test/e2e/.../

test: test/clean test/unit test/e2e

.PHONY: test
.PHONY: test/clean
.PHONY: test/unit
.PHONY: test/unit/koji-controller test/unit/kojiclient
.PHONY: test/e2e

build/clean:
	@rm -rf _build

build/prepare:
	@mkdir -p _build/bin

build/koji-controller: build/prepare
	@mkdir -p _build/bin
	@go build -v -o _build/bin/koji-controller koji-controller/main.go

build/kojid-cloud-scheduler: build/prepare
	@mkdir -p _build/bin
	@go build -v -o _build/bin/kojid-cloud-scheduler kojid-cloud-scheduler/main.go

build: build/prepare build/koji-controller build/kojid-cloud-scheduler


.PHONY: build
.PHONY: build/clean
.PHONY: build/prepare
.PHONY: build/koji-controller
