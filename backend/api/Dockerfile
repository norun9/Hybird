FROM golang:1.22-alpine AS build_base
RUN apk add --no-cache git
WORKDIR /tmp/gin

COPY .. .

WORKDIR /tmp/gin/api

RUN go mod download
RUN GOOS=linux CGO_ENABLED=0 go build -o ../bootstrap ./cmd/main.go

FROM alpine:3.9
RUN apk add ca-certificates tzdata
COPY --from=public.ecr.aws/awsguru/aws-lambda-adapter:0.8.3 /lambda-adapter /opt/extensions/lambda-adapter
COPY --from=build_base /tmp/gin/bootstrap /app/bootstrap
COPY --from=build_base /tmp/gin/api/config.yml /app/config.yml

EXPOSE 8080

CMD ["/app/bootstrap"]