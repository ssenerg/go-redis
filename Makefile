APP_NAME=go-redis
VERSION=1.0.0

ifeq (,$(wildcard ./.env))
    $(warning .env file not found)
else
    include .env
    export $(shell sed 's/=.*//' .env)
endif

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build: tidy
	@go build -o bin/$(APP_NAME)_$(VERSION) cmd/main.go

.PHONY: run
run: build
	@bin/$(APP_NAME)_$(VERSION)
