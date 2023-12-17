#!/bin/bash

echo "ENTRYPOINT KEK"

go install -v "golang.org/x/tools/gopls@latest" 
go install -v "github.com/cweill/gotests/gotests@v1.6.0"
go install -v "github.com/fatih/gomodifytags@v1.16.0"
go install -v "github.com/josharian/impl@v1.1.0"
go install -v "github.com/haya14busa/goplay/cmd/goplay@v1.0.0"
go install -v "github.com/go-delve/delve/cmd/dlv@latest"
go install -v "honnef.co/go/tools/cmd/staticcheck@latest"
go install -v "golang.org/x/tools/cmd/goimports@latest"

go install -v "google.golang.org/grpc@v1.60.0"
go install -v "google.golang.org/protobuf/cmd/protoc-gen-go@v1.28"
go install -v "google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2"

apt update
apt install -y protobuf-compiler
export PATH="$PATH:$(go env GOPATH)/bin"
