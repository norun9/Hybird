module github.com/norun9/HyBird/backend/lambda/websocket

go 1.23.0

replace github.com/norun9/HyBird/backend/pkg/log => ../../pkg/log

require (
	github.com/aws/aws-lambda-go v1.47.0
	github.com/aws/aws-sdk-go v1.55.5
	github.com/norun9/HyBird/backend/pkg/log v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.27.0
)

require (
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
)
