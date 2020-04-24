# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGENERATE=$(GOCMD) generate
GOGET=$(GOCMD) get
BINARY_LOC=bin
BINARY_NAME=sith

all: get generate test build
build: get
	$(GOBUILD) -o ./$(BINARY_LOC)/$(BINARY_NAME) -v
test: get generate
	$(GOTEST) -v ./...
generate:
	$(GOGENERATE) ./...
get:
	$(GOGET) -u -v github.com/golang/mock/mockgen
clean: 
	$(GOCLEAN)
	rm -f ./$(BINARY_LOC)/$(BINARY_NAME)
run: build
	./$(BINARY_LOC)/$(BINARY_NAME)


# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v    