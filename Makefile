GO ?= go
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOFMT ?= gofmt "-s"
GO_VERSION=$(shell $(GO) version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)
VERSION := $(shell git describe --tags --dirty="-dev")
VERSION_PKG := github.com/yulefox/log

TEST_FOLDER := $(shell $(GO) list ./... | grep -v examples)
TEST_TAGS ?= ""

.PHONY: all
all: build

.PHONY: build
build:
	GOARCH=$(GOARCH) GOOS=$(GOOS) $(GO) build -gcflags=all="-N -l" -ldflags "-s -w -X '$(VERSION_PKG).Version=$(VERSION)'" ./...

.PHONY: test
test:
	echo "mode: count" > coverage.out
	for d in $(TEST_FOLDER); do \
		$(GO) test $(TEST_TAGS) -v -covermode=count -coverprofile=profile.out $$d > tmp.out; \
		cat tmp.out; \
		if grep -q "^--- FAIL" tmp.out; then \
			rm tmp.out; \
			exit 1; \
		elif grep -q "build failed" tmp.out; then \
			rm tmp.out; \
			exit 1; \
		elif grep -q "setup failed" tmp.out; then \
			rm tmp.out; \
			exit 1; \
		fi; \
		if [ -f profile.out ]; then \
			cat profile.out | grep -v "mode:" >> coverage.out; \
			rm profile.out; \
		fi; \
	done