FROM golang:1.22

# RUN go install -v github.com/go-delve/delve/cmd/dlv@v1.22.1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .