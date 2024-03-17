compile: daemon-x64-lin

daemon-codestyle:
	go fmt ./...
	go vet ./...
	gocyclo -over 4 internal
	gocritic check ./...

daemon-x64-lin:
	go build ./...

grpc:
	buf generate

protoc:
	protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. -I .:/usr/include/ SpaghettiCannon.proto

webui-dist:
	rm -rf webui webui.dev/dist
	cd webui.dev && npm install
	cd webui.dev && parcel build --public-url "." && mv dist ../webui/

