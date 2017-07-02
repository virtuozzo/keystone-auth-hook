BINARY := keystone-auth-hook
TARGET := build/$(BINARY)

$(TARGET):
	glide install
	CGO_ENABLED=0 GOOS=linux go build -v -x -o $(TARGET) ./cmd

	docker build --pull -t $(BINARY) -f build/Dockerfile build

clean:
	rm -f $(TARGET)
