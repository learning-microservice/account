INSTANCE = account

default: build

init:
	go get -u github.com/golang/dep/cmd/dep
	dep init
	dep ensure --update

build: deps
	mkdir -p bin
	go build -o bin/$(INSTANCE) -v

test: build
	go test ./...

clean:
	go clean
	rm -rf bin
	rm -rf vendor

run: build
	bin/$(INSTANCE)

deps:
	dep ensure