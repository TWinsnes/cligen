.PHONY: lint build test

lint:
	golangci-lint run

build:
	go build -o dist/cligen

test:
	go test ./...