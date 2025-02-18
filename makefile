BINARY_NAME=leinao-docs-sync-bot
VERSION=1.0.0

all: run

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/$(BINARY_NAME) cmd/main.go
	sudo docker build -t $(BINARY_NAME):$(VERSION) .

run:
	go run cmd/main.go

test:
	go test ./...

clean:
	rm -rf build/
	rm -f $(BINARY_NAME)

.PHONY: all build run test docker-build clean