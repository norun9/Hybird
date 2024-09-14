# Stage 1: Build the Go application
FROM golang:1.22-alpine AS build

# from "backend" context
# command: docker build --build-arg LAMBDA_SOURCE_DIR=connect -t latest -f lambda/websocket/Dockerfile .

# Set the working directory inside the container
WORKDIR /app

# Copy the source code and dependencies
COPY lambda/websocket /app/lambda/websocket/
COPY ../../pkg /app/pkg

# Change to the websocket directory
WORKDIR /app/lambda/websocket

# Download Go modules
RUN go mod download

# Build the Go application
ARG LAMBDA_SOURCE_DIR=""
RUN GOOS=linux CGO_ENABLED=0 go build -o bootstrap ./${LAMBDA_SOURCE_DIR}/main.go

# Stage 2: Prepare the AWS Lambda runtime
FROM public.ecr.aws/lambda/go:1

# Copy the compiled Go binary from the build stage
COPY --from=build /app/lambda/websocket/bootstrap ${LAMBDA_TASK_ROOT}/bootstrap

# The entrypoint is automatically handled by the Lambda base image, so no need to specify CMD/ENTRYPOINT
CMD ["/${LAMBDA_TASK_ROOT}/bootstrap"]