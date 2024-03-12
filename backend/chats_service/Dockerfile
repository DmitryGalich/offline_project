FROM golang:1.22

# ADD . /offline_project

# WORKDIR /offline_project

# # Debugger
# RUN go install -v github.com/go-delve/delve/cmd/dlv@v1.22.1

# EXPOSE 80

# # Starting point
# CMD ["go","run","cmd/main.go"]


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY chats_service/go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY chats_service/cmd ./cmd

# Build the Go app
RUN go build -o main ./cmd

# Expose port 8080 to the outside world
EXPOSE 80

# Command to run the executable
CMD ["./main"]