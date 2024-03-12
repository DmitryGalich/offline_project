FROM golang:1.22


ADD . /home

WORKDIR /home

CMD ["go","run","cmd/main.go"]

# RUN go build cmd/
# RUN ./cmd/main

# Debugger
# RUN go install -v github.com/go-delve/delve/cmd/dlv@v1.22.1