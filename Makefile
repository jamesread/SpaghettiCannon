compile: daemon-compile-x64-lin

daemon-codestyle:
	go fmt ./...
	go vet ./...
	gocyclo -over 4 internal
	gocritic check ./...

daemon-unittests:

daemon-compile-x64-lin:
	go build ./...

grpc: go-tools
	buf generate

protoc:
	protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. -I .:/usr/include/ SpaghettiCannon.proto

webui-codestyle:
	cd webui.dev && npm install
	cd webui.dev && ./node_modules/.bin/eslint main.ts ts/*
	cd webui.dev && ./node_modules/.bin/stylelint style.css

webui-dist:
	rm -rf webui webui.dev/dist
	cd webui.dev && npm install
	cd webui.dev && parcel build --public-url "." && mv dist ../webui/

go-tools:
	go install "github.com/bufbuild/buf/cmd/buf"
	go install "github.com/fzipp/gocyclo/cmd/gocyclo"
	go install "github.com/go-critic/go-critic/cmd/gocritic"
	go install "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	go install "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	go install "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	go install "google.golang.org/protobuf/cmd/protoc-gen-go"

.PHONY: grpc
