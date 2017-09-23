TEST_FILES = $(shell cd src ; find -L * -name '*_test.go' )
TEST_PACKAGES = $(dir $(TEST_FILES))

GOPATH := $(PWD):$(PWD)/vendor
export GOPATH


all:  twittergostalker

twittergostalker: bin/gb
	bin/gb build $@

bin/gb:
	go build -o bin/gb github.com/constabulary/gb/cmd/gb

test:
	go test $(TEST_PACKAGES)

vet:
	go tool vet ./src

clean:
	rm -rf bin/
	rm -rf pkg/
	rm -rf vendor/pkg


.PHONY: all test vet clean twittergostalker

