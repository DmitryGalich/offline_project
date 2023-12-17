FROM golang:1.21

WORKDIR /offline_project/backend/service_gateway

COPY . .

RUN go install -v "golang.org/x/tools/gopls@latest"
RUN go install -v "github.com/cweill/gotests/gotests@v1.6.0"
RUN go install -v "github.com/fatih/gomodifytags@v1.16.0"
RUN go install -v "github.com/josharian/impl@v1.1.0"
RUN go install -v "github.com/haya14busa/goplay/cmd/goplay@v1.0.0"
RUN go install -v "github.com/go-delve/delve/cmd/dlv@latest"
RUN go install -v "honnef.co/go/tools/cmd/staticcheck@latest"
RUN go install -v "golang.org/x/tools/cmd/goimports@latest"

# RUN go get "google.golang.org/grpc/credentials/insecure"
RUN go get "go get google.golang.org/grpc@v1.60.0"
#RUN go get "google.golang.org/grpc/status"
RUN go install -v "google.golang.org/protobuf/cmd/protoc-gen-go@v1.28"
#RUN go install -v "google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2"
#RUN go get "google.golang.org/grpc/codes"
RUN apt update
RUN apt install -y protobuf-compiler
RUN export PATH="$PATH:$(go env GOPATH)/bin"

EXPOSE 8080