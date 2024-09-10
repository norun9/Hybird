FROM golang:1.23-alpine as build

# from "backend" context
# command: docker build --build-arg LAMBDA_SOURCE_DIR=connect -t latest -f lambda/websocket/Dockerfile .

WORKDIR /app

COPY lambda/websocket ./lambda/websocket/

WORKDIR /app/lambda/websocket

COPY ../../pkg /app/pkg

RUN go mod download

ARG LAMBDA_SOURCE_DIR=""

RUN GOOS=linux CGO_ENABLED=0 go build -o main ./${LAMBDA_SOURCE_DIR}/main.go

FROM public.ecr.aws/lambda/provided:al2023

COPY --from=build /app/lambda/websocket/main ./main

ENTRYPOINT ["./main"]