GOCMD=go
GOGET=$(GOCMD) get
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
BINARY_NAME=bin/backend

BIN_DIR=bin

.PHONY: all test build clean fmt

all: get build

get:
	$(GOGET) ./...

run:
	$(GORUN) ./...

build:
	$(GOBUILD) -o $(BINARY_NAME) ./main.go

fmt:
	$(GOFMT) ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

