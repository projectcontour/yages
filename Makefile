.PHONY: install-tools
install-tools:
	./hack/install-tools.sh

.PHONY: generate-pb-go
generate-pb-go:
	protoc  --proto_path=${PWD} --go-grpc_out=. --go_out=. yages-schema.proto

.PHONY: build
build:
	goreleaser build --snapshot --rm-dist

.PHONY: release
release:
	goreleaser release --rm-dist

.PHONY: clean
clean:
	rm -rf ${PWD}/dist
