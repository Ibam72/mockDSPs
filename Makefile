REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.revision=$(REVISION)'

.PHONY: setup fmt curl

## Setup
setup:
	go get golang.org/x/tools/cmd/goimports
	go get github.com/Songmu/make2help/cmd/make2help

## Format source codes
fmt:
	find . -name "*.go" -not -path "./vendor/*" | xargs goimports -w

## Run Run away
run:
	go run cmd/main.go

## curl
curl:
	curl -X POST localhost:11115/mockDSPs/19 -d '$(shell cat testRequest.json | jq -c)' | jq

BINDIR=./bin
## build linux binary
build-linux:
	mkdir -p $(BIN)/linux
	GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags "$(LDFLAGS)" -o $(releaseBINDIR)/$(HoneyCat) cmd/$(HoneyCat)/main.go
