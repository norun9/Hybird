FROM golang:1.21-alpine3.19 AS build

WORKDIR /app

RUN go install github.com/pilu/fresh@latest \
    && go install github.com/pressly/goose/v3/cmd/goose@v3.17.0

# マイグレーションファイルをコピー
COPY migrations /app/db/migrations

# 最終イメージ
FROM alpine:3.19

WORKDIR /app

# 必要なファイルとバイナリのみをコピー
COPY --from=build /app/db /app/db
COPY --from=build /go/bin/fresh /usr/local/bin/fresh
COPY --from=build /go/bin/goose /usr/local/bin/goose

# 最小限のパッケージをインストール
RUN apk add --no-cache bash mysql-client
