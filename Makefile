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
	curl -X POST localhost:8080/mockDSPs/19 -d '$(shell cat testRequest.json | jq -c)' -vvv
