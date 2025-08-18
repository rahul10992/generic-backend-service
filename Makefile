build:
	go build -o bin/main cmd/hello-svc/main.go

run: build
	bin/main

clean:
	rm -rf bin/

install-proto:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go@latest

proto-gen: tools-check
	ls api/proto/health/v1/health.proto
	ls internal/gen
	protoc --proto_path=api/proto \
		--go_out=internal/gen \
		--go_opt=paths=source_relative \
		--go-grpc_out=internal/gen \
		--go-grpc_opt=paths=source_relative \
		api/proto/health/v1/health.proto

proto-clean:
	rm -rf internal/gen/

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

