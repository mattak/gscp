.PHONY: build test clean

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
BINARY_NAME=gscp

all: clean build

build:
	cd cmd && $(GOBUILD) -o $(BINARY_NAME) -v

test:
	cd internal && $(GOTEST) -v

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
