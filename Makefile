.SILENT:
.EXPORT_ALL_VARIABLES:
.PHONY: gen test deps

run:
	go run cmd/refresh-hash/main.go
	
generate:
	protoc 	--go_out=gen \
	--go_opt=paths=source_relative \
	--go-grpc_out=gen \
	--go-grpc_opt=paths=source_relative \
	proto/hash.proto

compose-up:
	docker-compose up

compose-down:
	docker-compose down

test:
	go clean -testcache
	go test ./...

deps:
	go get -u ./...
	go mod tidy
	go mod vendor