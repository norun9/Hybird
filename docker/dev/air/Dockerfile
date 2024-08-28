FROM golang:1.22-alpine as builder

WORKDIR /app/go/builder

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/air-verse/air@latest

COPY . .

FROM golang:1.22-alpine

WORKDIR /app/go/base

ENV HYBIRD_ENV_NAME=dev

COPY --from=builder /go/pkg/mod /go/pkg/mod
COPY --from=builder /go/bin/air /go/bin/air