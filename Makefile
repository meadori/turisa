.PHONY: all
all: build

src := $(shell find . -type f -name '*.go')
tests := $(shell find . -type f -name '*_test.go')

.PHONY: build
build: $(src)
	go build $(shell dirname $(src)) 

.PHONY: fmt
fmt:
	go fmt $(shell dirname $(src)) 

.PHONY: test
test:
	go test $(shell dirname $(tests)) 

.PHONY: clean
clean:
	rm -rf bin
