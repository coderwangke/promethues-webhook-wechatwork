BINARY_NAME := $(shell basename "$(PWD)")
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

VERSION := $(shell git describe --tags --always)

GOOS := "linux"
GOARCH := "amd64"

all: build

build:
    @echo "  >  Building binary..."
    @GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(GOBIN)/$(BINARY_NAME) main.go

image:
    @echo "  >  Building image..."

push:
    @echo "  >  Pushing image..."
