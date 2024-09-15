FROM golang:1.22-alpine as builder

RUN go install github.com/air-verse/air@latest

FROM golang:1.22-alpine

WORKDIR /app/go/base

ENV HYBIRD_ENV_NAME=dev

COPY --from=builder /go/pkg/mod /go/pkg/mod
COPY --from=builder /go/bin/air /go/bin/air