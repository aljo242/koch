#!/usr/bin/make -f
PACKAGES=$(shell go list ./...)
GOBIN ?= $(GOPATH)/bin
VERSION := $(shell echo $(shell git describe --tags 2> /dev/null || echo "dev-$(shell git describe --always)") | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
PROJECT_NAME = $(shell git remote get-url origin | xargs basename -s .git)
ARTIFACT_DIR := ./artifacts
BINDIR ?= $(GOPATH)/bin
DEMOAPP = ./demo/kochd

export GO111MODULE = on

###############################################################################
###                               Build                                     ###
###############################################################################

all: install lint test

install: go.sum
	@echo "--> Installing kochd"
	@go install -mod=readonly $(BUILD_FLAGS) $(DEMOAPP)

build: go.sum
	@echo "--> Building kochd"
	@go build -o $(ARTIFACT_DIR)/pylonsd -mod=readonly $(BUILD_FLAGS) $(DEMOAPP)

clean:
	@rm -rf $(ARTIFACT_DIR)

go.sum: go.mod
	@echo "Ensure dependencies have not been modified ..."
	@go mod verify
	@go mod tidy -go=1.17

.PHONY: install go.sum clean build

###############################################################################
###                                Testing                                  ###
###############################################################################

test: test-unit

test-unit:
	@VERSION=$(VERSION) go test -mod=readonly -v -timeout 30m $(PACKAGES)

COVER_FILE := coverage.txt
COVER_HTML_FILE := cover.html

test-cover:
	@VERSION=$(VERSION) go test -mod=readonly -v -timeout 30m -coverprofile=$(COVER_FILE) -covermode=atomic $(PACKAGES)
	@go tool cover -html=$(COVER_FILE) -o $(COVER_HTML_FILE)

bench:
	@VERSION=$(VERSION) go test -mod=readonly -v -timeout 30m -bench=. $(PACKAGES)

.PHONY: test test-unit test-race test-cover bench

###############################################################################
###                                Linting                                  ###
###############################################################################

markdownLintImage=tmknom/markdownlint
containerMarkdownLint=$(PROJECT_NAME)-markdownlint
containerMarkdownLintFix=$(PROJECT_NAME)-markdownlint-fix

lint:
	@golangci-lint run -c ./.golangci.yml --out-format=tab --issues-exit-code=0
	@# @if $(DOCKER) ps -a --format '{{.Names}}' | grep -Eq "^${containerMarkdownLint}$$"; then $(DOCKER) start -a $(containerMarkdownLint); else $(DOCKER) run --name $(containerMarkdownLint) -i -v "$(CURDIR):/work" $(markdownLintImage); fi


FIND_ARGS := -name '*.go' -type f -not -path "./sample_txs*" -not -path "*.git*" -not -path "./build_report/*" -not -path "./scripts*" -not -name '*.pb.go'

format:
	@find . $(FIND_ARGS) | xargs gofmt -w -s
	@find . $(FIND_ARGS) | xargs goimports -w -local github.com/Pylons-tech/pylons

proto-lint:
	@buf lint --error-format=json


.PHONY: lint format