GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOFMT=gofmt -w

BINARY_NAME=xeno

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/xeno

test:
	$(GOTEST) -v .

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/xeno
	./$(BINARY_NAME)

install:
	$(GOINSTALL) ./cmd/xeno

fmt:
	$(GOFMT) .

.PHONY: all build test clean run install fmt
