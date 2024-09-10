module github.com/norun9/HyBird/backend/lambda/websocket

go 1.23.0

replace github.com/norun9/HyBird/backend/pkg/log => ../../pkg/log

require (
	github.com/aws/aws-lambda-go v1.47.0
	github.com/aws/aws-sdk-go-v2 v1.30.5
	github.com/aws/aws-sdk-go-v2/config v1.15.0
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.15.3
	github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi v1.21.6
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.34.9
	github.com/norun9/HyBird/backend/pkg/log v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.27.0
)

require (
	github.com/aws/aws-sdk-go-v2/credentials v1.10.0 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.17 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.17 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.22.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.11.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.9.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.16.0 // indirect
	github.com/aws/smithy-go v1.20.4 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
)
