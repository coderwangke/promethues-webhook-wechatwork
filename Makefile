BINARY_NAME := $(shell basename "$(PWD)")
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

VERSION := "latest"

GOOS := "linux"
GOARCH := "amd64"

REGISTRY := "ccr.ccs.tencentyun.com"
IMAGE_NAME := hale/$(BINARY_NAME)

all: build


test:
	@echo "  >  Testing..."
	@go test

lint:
	@echo "  >  Linting..."
	@golint

build:
	@echo "  >  Building binary..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(GOBIN)/$(BINARY_NAME) *.go

image:
	@echo "  >  Building image..."
	@docker build -t $(REGISTRY)/$(IMAGE_NAME):$(VERSION) .

push:
	@echo "  >  Pushing image..."
	@docker push $(REGISTRY)/$(IMAGE_NAME):$(VERSION)
