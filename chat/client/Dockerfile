# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest
LABEL maintainer="Sonia <sonia@rhizicube.ai>"

# Set the Current Working Directory inside the container
WORKDIR /chat_service

# Copy go mod and sum files
RUN go mod init chat_service
RUN go mod tidy

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Build the Go app
# EXPOSE 9093
# Command to run the executable

RUN go env -w GO111MODULE=auto
EXPOSE 50051

CMD ["go", "run", "client.go"]