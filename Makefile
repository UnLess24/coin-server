.SILENT:
.PHONY: up down lint .bin-deps .buf-generate .tidy generate

up:
	docker network create --driver bridge coinclientserver &> /dev/null || true
	docker-compose up -d

down:
	docker-compose down --rmi local
	docker network rm coinclientserver -f &> /dev/null || true

lint:
	golangci-lint run ./...

LOCAL_BIN := $(CURDIR)/bin
BUF_BUILD := $(LOCAL_BIN)/buf

.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	$(info Installing binary dependencies...)

	go install github.com/bufbuild/buf/cmd/buf@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.buf-generate:
	$(info Generating protobuf files...)
	PATH="$(LOCAL_BIN):$(PATH)" $(BUF_BUILD) generate

generate: .buf-generate .tidy

.tidy:
	go mod tidy
