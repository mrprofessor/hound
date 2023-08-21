.PHONY: build install release
# Initialize Githooks, along with anything else needed to dev and build this repo
init:
	git config core.hooksPath .githooks

default: build

dep:
	@go mod tidy && go mod download


build:
	@go build


install:
	@go install

test:
	@go test ./... -race  -cover -coverprofile coverage.txt -covermode=atomic

release:
	goreleaser release
