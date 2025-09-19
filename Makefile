.PHONY: lint build

lint:
	golangci-lint run

build:
	go build -o dist/cligen