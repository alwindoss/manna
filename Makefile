# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_LOC=bin
BINARY_NAME=manna
VERSION=0.0.1

all: build
docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./$(BINARY_LOC)/$(BINARY_NAME) -v
package:
	docker build -t alwindoss/$(BINARY_NAME):$(VERSION) .
publish:
	docker push alwindoss/$(BINARY_NAME):$(VERSION)
setup:
	$(GOGET) -v ./...
build: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./$(BINARY_LOC)/$(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -rf $(BINARY_LOC)
	rm -f *.db
run: clean build
	./$(BINARY_LOC)/$(BINARY_NAME)