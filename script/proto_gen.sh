#!/bin/bash

# Install dependencies
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc -I ./proto \
    --plugin=$(go env GOPATH)/bin/protoc-gen-go \
    --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc \
    --go_out ./proto --go_opt paths=source_relative \
    --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
    ./proto/*.proto

protoc -I ./email_service/proto/email \
    --plugin=$(go env GOPATH)/bin/protoc-gen-go \
    --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc \
    --go_out ./proto --go_opt paths=source_relative \
    --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
    ./proto/*.proto
