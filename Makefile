SHELL := /bin/bash

code/check:
	@gofmt -l `find . -type f -name '*.go' -not -path "./vendor/*"`

code/fmt:
	@gofmt -w `find . -type f -name '*.go' -not -path "./vendor/*"`

.PHONY: code/fmt code

test/clean:
	@go clean -testcache

test/unit:
	@go test -v ./pkg/.../

test/e2e:
	@go test -v ./test/e2e/.../

test: test/clean test/unit test/e2e

.PHONY: test/clean test/unit test/e2e  test
