NAME := MATERIALCOLOR
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.revision=$(REVISION)'
TEMPLATEDIR := ./template

.PHONY: setup fmt curl build-linux scp linux-run clean

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
	curl -X POST localhost:11115/19 -d '$(shell cat testRequest.json | jq -c)' | jq
## curl fenrir
curl-fenrir:
	curl -X POST fenrir:19191/19 -d '$(shell cat testRequest.json | jq -c)' | jq

BINDIR=./bin
LINUX=$(BINDIR)/linux
## build linux binary
build-linux: clean
	mkdir -p $(LINUX)
	GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(LINUX)/$(NAME) cmd/main.go
	cp -r $(TEMPLATEDIR) $(LINUX)/$(TEMPLATEDIR)
	cp Makefile $(LINUX)/Makefile

scp: build-linux
	scp -r $(LINUX) fenrir:~/$(NAME)

linux-run:
	./$(NAME) -p 19191

clean:
	rm -rf $(BINDIR)
