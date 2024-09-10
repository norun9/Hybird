# ビルドステージ: Goのビルド環境を設定
FROM golang:1.23-alpine AS build

WORKDIR /app

# Assuming that the context is /backend
COPY ../go.mod ./
COPY go.sum ./

# ビルドに必要な依存関係をインストール
RUN apk add --no-cache make gcc g++ git

# 作業ディレクトリを作成
WORKDIR /go/src/app

# Goのモジュールファイルをコピーして依存関係をダウンロード
COPY ../go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# Goアプリケーションをビルド
RUN GOOS=linux GOARCH=amd64 go build -o main main.go