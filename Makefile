.PHONY: build clear
NAME = apiserver

build:
	go mod tidy
	go build -v ./cmd/$(NAME)

.DEFAULT_GOAL := build

clear:
	rm $(NAME)