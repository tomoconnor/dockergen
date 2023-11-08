# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=dockergen

# Build for all platforms
all: build-linux build-windows build-macos

# Build for Linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o dist/$(BINARY_NAME)-linux-amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=386 $(GOBUILD) -o dist/$(BINARY_NAME)-linux-386

# Build for Windows
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o dist/$(BINARY_NAME)-windows-amd64.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=386 $(GOBUILD) -o dist/$(BINARY_NAME)-windows-386.exe

# Build for MacOS
build-macos:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o dist/$(BINARY_NAME)-macos-amd64
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOBUILD) -o dist/$(BINARY_NAME)-macos-arm64

# Clean up
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)-*

# Run tests
test:
	$(GOCMD) test -v .

# Fetch dependencies
deps:
	$(GOGET) -v ./...

# Run
run:
	$(GOCMD) run main.go

# Cross-compile
cross-compile: build-linux build-windows build-macos
