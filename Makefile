NAME := MATERIALCOLOR
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.revision=$(REVISION)'
TEMPLATEDIR := ./template

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
LINUX=$(BINDIR)/linux
## build linux binary
build-linux:
	mkdir -p $(LINUX)
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(LINUX)/$(NAME) cmd/main.go
	cp -r $(TEMPLATEDIR) $(LINUX)/$(TEMPLATEDIR)
	cp Makefile $(LINUX)/Makefile

linux-run:
	./$(NAME) -p 14514
