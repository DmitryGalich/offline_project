FROM golang:1.22

ADD . /offline_project

WORKDIR /offline_project

# Debugger
#RUN go install -v github.com/go-delve/delve/cmd/dlv@v1.22.1

EXPOSE 80

# Starting point
CMD ["go","run","cmd/main.go"]
