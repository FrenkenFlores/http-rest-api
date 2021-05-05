.PHONY: build clear
NAME = apiserver

build:
	go build -v ./cmd/$(NAME)

.DEFAULT_GOAL := build

clear:
	rm $(NAME)