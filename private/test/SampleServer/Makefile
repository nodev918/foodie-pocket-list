GOPATH ?= $(shell go env GOPATH)
export PATH := $(GOPATH)/bin:$(PATH)

.PHONY: gen-mock gen-proto test build lint clean

gen-proto:
	mkdir -p internal/gen/pb/notification
	protoc -I./submodules/rocks --go_out=./internal/gen/pb --go_opt=paths=source_relative --go-grpc_out=./internal/gen/pb --go-grpc_opt=paths=source_relative ./submodules/rocks/notification/club_notification.proto
	protoc -I./submodules/rocks --go_out=./internal/gen/pb --go_opt=paths=source_relative --go-grpc_out=./internal/gen/pb --go-grpc_opt=paths=source_relative ./submodules/rocks/notification/club_metadata.proto
	protoc -I./submodules/rocks --go_out=./internal/gen/pb --go_opt=paths=source_relative --go-grpc_out=./internal/gen/pb --go-grpc_opt=paths=source_relative ./submodules/rocks/notification/notification.proto

	mkdir -p internal/gen/pb/health
	protoc -I./submodules/rocks --go_out=./internal/gen/pb/health --go_opt=paths=source_relative --go-grpc_out=./internal/gen/pb/health --go-grpc_opt=paths=source_relative ./submodules/rocks/healthCheck.proto

gen-mock:
	make -C submodules/caerus $@
	mockgen -source internal/worker/message_builder_factory/builder.go -destination internal/worker/message_builder_factory/builder_mock.go -package messagebuilderfactory
	mockgen -source internal/worker/message_builder_factory/message_factory.go -destination internal/worker/message_builder_factory/message_factory_mock.go  -package messagebuilderfactory
	mockgen -source internal/kafka/notification.go  -destination internal/kafka/notification_mock.go -package kafka
	mockgen -source internal/repository/notification.go  -destination internal/repository/notification_mock.go -package repository
	mockgen -source internal/grpc/controller/last_resend_cache/last_resend_cache.go  -destination internal/grpc/controller/last_resend_cache/last_resend_cache_mock.go -package lastresendcache

test: gen-proto gen-mock
	go test -v -race ./... -cover -coverprofile c.out
	go tool cover -html c.out

lint: gen-proto gen-mock
	golangci-lint run ./...

build: gen-proto
	go build -o stork_grpc ./cmd/server
	go build -o stork_worker ./cmd/worker

clean:
	make -C submodules/caerus $@
	rm -rf internal/gen build
	rm -f stork_grpc stork_worker
	rm -f c.out
	find .  -name "*_mock.go" -delete

