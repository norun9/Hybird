# Stage 1: Build the Go application
FROM golang:1.22-alpine AS build_base

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies first (caching purpose)
COPY lambda/websocket/go.mod lambda/websocket/go.sum /app/lambda/websocket/
COPY pkg/log/go.mod pkg/log/go.sum /app/pkg/log/

# Change to the websocket directory
WORKDIR /app/lambda/websocket

# Download Go modules
RUN go mod download

# Copy the rest of the source code
COPY lambda/websocket /app/lambda/websocket/
COPY pkg /app/pkg/

# Build the Go application
ARG LAMBDA_SOURCE_DIR=""
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /app/bootstrap ./${LAMBDA_SOURCE_DIR}/main.go

# Stage 2: Prepare the AWS Lambda runtime
FROM public.ecr.aws/lambda/provided:al2023

# Copy the compiled Go binary from the build stage to the Lambda task directory
COPY --from=build_base /app/bootstrap ./main

ENTRYPOINT ["./main"]

