# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest
LABEL maintainer="Sonia <sonia@rhizicube.ai>"
ENV GO111MODULE=on
# Set the Current Working Directory inside the container
WORKDIR /chat_service
#COPY go.mod .
#COPY go.sum .

# Copy go mod and sum files
RUN go mod init chat_service

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
RUN go get google.golang.org/grpc
RUN go get github.com/SoniaPunjabi/chat_service

# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Build the Go app
# EXPOSE 9093
# Command to run the executable


EXPOSE 50051

CMD ["go", "run", "server.go"]