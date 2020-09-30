.PHONY: all build gotool

BINARY="k8s"

all: build gotool

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY} .

gotool:
	@go fmt ./
	@go vet ./