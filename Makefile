.PHONY: build run test

BIN_DIR := bin
BIN := $(BIN_DIR)/logmatcher

build:
	go build -o $(BIN) ./cmd/logmatcher

run: build
	./$(BIN)

start:
	./$(BIN)

test:
	go test ./...
