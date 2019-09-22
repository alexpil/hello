GOOS?=darwin
GOARCH?=amd64

GOPROXY?=https://gocenter.io

COMMIT := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
RELEASE := v0.0.1

build:
	GOOS=${GOOS} GOARCH=${GOARCH} \
	GO111MODULE=on GOPROXY=${GOPROXY} \
	CGO_ENABLED=0 go build \
		-ldflags "-s -w -X github.com/alexpil/hello/internal/diagnostics.Version=${RELEASE} \
		-X github.com/alexpil/hello/internal/diagnostics.Hash=${COMMIT} \
		-X github.com/alexpil/hello/internal/diagnostics.BuildTime=${BUILD_TIME}" \
		-o bin/hello ./cmd/hello
