.PHONY: proto proto-gen proto-clean
PROTO_SRC_DIR := api/proto
PROTO_OUT_DIR := internal/gen

.DEFAULT_GOAL := all

.PHONY: all
all: rebuild tools-check proto-gen build

build:
	go build -o bin/main cmd/hello-svc/main.go

run: build
	bin/main

.PHONY: clean
clean: proto-clean
	rm -rf bin/

.PHONY: fmt
fmt:
	@echo "Formatting with gofmt"
	@gofmt -s -w .

.PHONY: fumpt
fumpt:
	go install mvdan.cc/gofumpt@latest
	@gofumpt -l -w .

.PHONY: lint
lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@golangci-lint run --skip-dirs=internal/gen ./...

.PHONY: rebuild
rebuild: clean build

.PHONY: vet
vet:
	@go vet ./...

proto: tools-check proto-clean proto-gen

.PHONY: install-proto
install-proto:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go@latest

proto-gen:
	@echo "Generating protos → $(PROTO_OUT_DIR)"
	@mkdir -p $(PROTO_OUT_DIR)
	protoc --proto_path=$(PROTO_SRC_DIR) \
		--go_out=$(PROTO_OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_OUT_DIR) \
		--go-grpc_opt=paths=source_relative \
		$$(find $(PROTO_SRC_DIR) -name '*.proto' -print)

.PHONY: proto-clean
proto-clean:
	rm -rf $(PROTO_OUT_DIR)/*

tools-check:
	@if [ -z "$$(protoc-gen-go --version 2>/dev/null)" ]; then \
		echo "\033[0;31mprotoc-gen-go not found or version blank\033[0m"; \
		exit 1; \
	else \
		echo "\033[0;32mprotoc-gen-go OK\033[0m"; \
	fi
	@if [ -z "$$(protoc-gen-go-grpc -version 2>/dev/null)" ]; then \
		echo "\033[0;31mprotoc-gen-go-grpc not found or version blank\033[0m"; \
		exit 1; \
	else \
		echo "\033[0;32mprotoc-gen-go-grpc OK\033[0m"; \
	fi
	@if [ -z "$$(protoc --version 2>/dev/null)" ]; then \
		echo "\033[0;31mprotoc not found or version blank\033[0m"; \
		exit 1; \
	else \
		echo "\033[0;32mprotoc OK\033[0m"; \
	fi

