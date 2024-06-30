FROM golang:1.22

# RUN go install -v github.com/go-delve/delve/cmd/dlv@v1.22.1

EXPOSE 5001

WORKDIR /app
COPY . .
COPY go.mod ./
RUN go mod download
RUN go build -o main ./cmd/messenger
CMD ["./main"]