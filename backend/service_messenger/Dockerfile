FROM golang:1.22

RUN go install golang.org/x/tools/gopls@v0.16.1
RUN go install honnef.co/go/tools/cmd/staticcheck@v0.4.7
RUN go install github.com/go-delve/delve/cmd/dlv@v1.23.0

EXPOSE 5001

WORKDIR /service_messenger

COPY . .

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go build -o main ./cmd/messenger

CMD ["./main"]