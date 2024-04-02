FROM golang:1.22

WORKDIR /app

COPY service_chats/go.mod ./

RUN go mod download
RUN go install -v github.com/go-delve/delve/cmd/dlv@v1.22.1
RUN go get github.com/gorilla/websocket
RUN go get github.com/sirupsen/logrus

COPY service_chats/cmd ./cmd

RUN go get service_chats/cmd

RUN go build -o main ./cmd

EXPOSE 80

CMD ["./main"]
