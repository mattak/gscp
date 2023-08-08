.PHONY: build test clean

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
BINARY_NAME=gscp

all: clean build

build:
	cd cmd/gscp && $(GOBUILD) -o $(BINARY_NAME) -v

test:
	cd pkg/gscp && $(GOTEST) -v

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)


install:
	cd cmd/gscp && $(GOINSTALL)
