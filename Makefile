# Enable Go Module support.
export GO111MODULE=on
export GOFLAGS=-mod=vendor

# Default Go linker flags.
GO_LDFLAGS := -ldflags="-s -w"

.PHONY: usage
usage:
	@echo
	@echo "usage:"
	@echo "   make all     			---- Remove go binary and .zip files"
	@echo "   make build     		---- Build artifact"
	@echo "   make zip     			---- zip up artifact for aws lambda usage"
	@echo "   make vendor    		---- Update dependencies"

.PHONY: all
all: clean build zip

.PHONY: clean
clean:
	rm -f main main.zip

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build $(GO_LDFLAGS) -o main

.PHONY: zip
zip: build
	zip main.zip main

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

