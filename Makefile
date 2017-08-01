BINARY := keystone-auth-hook
TARGET := build/$(BINARY)
GOOS?=linux

all: build

deps:
	glide install --strip-vendor

clean:
	rm -f $(TARGET)

build: clean deps
	CGO_ENABLED=0 GOOS=$(GOOS) go build -v -x -o $(TARGET) ./cmd

docker: build
	docker build --pull -t $(BINARY) -f build/Dockerfile build

.PHONY: all deps clean build docker
