#! /usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

sudo apt install -y protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.47.3

go install github.com/goreleaser/goreleaser@v1.10.3
