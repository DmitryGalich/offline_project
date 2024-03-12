FROM golang:1.22

WORKDIR /app

COPY chats_service/go.mod ./

RUN go mod download
RUN go install -v github.com/go-delve/delve/cmd/dlv@v1.22.1

COPY chats_service/cmd ./cmd
RUN go build -o main ./cmd

EXPOSE 80

CMD ["./main"]
