FROM golang:1.21-alpine3.19

WORKDIR /app

RUN apk add --no-cache \
        alpine-sdk=1.0-r1 bash mysql-client \
    && go install github.com/pilu/fresh@latest \
    && go install github.com/pressly/goose/v3/cmd/goose@v3.17.0

RUN mkdir -p /app/db

COPY . db/