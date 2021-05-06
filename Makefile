.PHONY: build test clear
.DEFAULT_GOAL := build

NAME = apiserver

build:
	go mod tidy
	go build -v ./cmd/$(NAME)

test:
	go test -v -race -timeout 30s ./...

clear:
	rm $(NAME)