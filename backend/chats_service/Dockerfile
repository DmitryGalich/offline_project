FROM golang:1.22

WORKDIR /app

# For docker-compose
# COPY chats_service/go.mod ./
# For dockerfile
COPY go.mod ./

RUN go mod download
RUN go install -v github.com/go-delve/delve/cmd/dlv@v1.22.1
RUN go get github.com/gorilla/websocket

# For docker-compose
# COPY chats_service/cmd ./cmd
# For dockerfile
COPY cmd ./cmd

RUN go build -o main ./cmd

EXPOSE 80

CMD ["./main"]
