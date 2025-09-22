.PHONY: lint build build-linux test vhs clean

lint:
	golangci-lint run

build:
	go build -o dist/cligen

build-linux:
	GOOS=linux go build -o dist/cligen

test:
	go test ./...

vhs:
	make build-linux;
	cp demo.tape dist/demo.tape;
	cd dist && docker run --rm -v $$(PWD):/vhs ghcr.io/charmbracelet/vhs demo.tape;
	cp dist/demo.gif demo.gif;
	make clean;

clean:
	rm -rf ./dist