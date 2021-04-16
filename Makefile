GO = go
GO_BUILD = $(GO) build
GO_TEST := $(GO) test -timeout 600s

VERSION = 3.4.2
COMMIT = $(shell git rev-parse --short HEAD)
PROJECT = github.com/hsqlu/comet-backup

export TRAFFIC_CONTROLLER = $(shell pwd)
export GOARCH=amd64

TARGET_DIR = target
BUILD_TIME := $(shell date +%Y%m%d_%H%M%S%z)
LD_FLAGS :=
.PHONY: all build test fmt clean

all: build

build: $(eval SHELL:=/bin/bash)
	$(GO_BUILD) -ldflags "$(LD_FLAGS)" -o $(TARGET_DIR)/traffic-controller

test:
	go clean -testcache
ifeq ($(origin VERBOSE),undefined)
	$(GO_TEST) ./...
else
	$(GO_TEST) -v ./...
endif

clean:
	rm -rf ${TARGET_DIR}

fmt:
	$(GO) fmt ./...