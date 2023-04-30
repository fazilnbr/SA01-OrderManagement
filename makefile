SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage

all: test build


proto: ## Generate the grpc protoc
	protoc pb/*.proto --go_out=plugins=grpc:.