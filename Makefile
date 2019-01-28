GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOVERSION=$(shell go version)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

NAME=go-practice
BIN=./bin/$(NAME)
# VERSION=$(shell git describe --tags --abbrev=0)
VERSION=0.0.1
REVISION=$(shell git rev-parse --short HEAD)
LDFLAGS="-w -s -X main.version=$(VERSION) -X main.revision=$(REVISION)"

clean:
	$(GOCLEAN)
	rm -f $(BIN)

test:
	$(GOTEST) -v ./...

build:
	$(GOBUILD) -ldflags $(LDFLAGS) -o $(BIN)

run: test build
	@$(BIN)
