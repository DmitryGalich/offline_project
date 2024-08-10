FROM golang:1.22

EXPOSE 5001

RUN go install golang.org/x/tools/gopls@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
