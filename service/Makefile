cicd: go-tools compile unittests
	rm SpaghettiCannon

compile:
	go build

codestyle:
	go fmt ./...
	go vet ./...
	gocyclo -over 4 internal
	gocritic check ./...

unittests:
	go test -v ./...

go-tools:
	go install "github.com/bufbuild/buf/cmd/buf"
	go install "github.com/fzipp/gocyclo/cmd/gocyclo"
	go install "github.com/go-critic/go-critic/cmd/gocritic"
	go install "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	go install "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	go install "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	go install "google.golang.org/protobuf/cmd/protoc-gen-go"
	go install "connectrpc.com/connect/cmd/protoc-gen-connect-go"
