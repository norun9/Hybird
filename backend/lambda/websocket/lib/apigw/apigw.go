package apigw

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/norun9/HyBird/backend/pkg/log"
	"go.uber.org/zap"
	"net/url"
)

// NewAPIGatewayClient creates a new API Gateway Management API client
func NewAPIGatewayClient(ctx context.Context, callbackURL url.URL) (*apigatewaymanagementapi.Client, error) {
	// Load default AWS config
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	// Create the ApiGatewayManagementApi client with the specified endpoint
	return apigatewaymanagementapi.NewFromConfig(cfg, func(o *apigatewaymanagementapi.Options) {
		o.Region = cfg.Region
		o.Credentials = cfg.Credentials
		o.BaseEndpoint = aws.String(callbackURL.String())
	}), nil
}

// SendMessageToConnection sends a message to a specified WebSocket connection
func SendMessageToConnection(ctx context.Context, client *apigatewaymanagementapi.Client, connectionId, message string) error {
	input := &apigatewaymanagementapi.PostToConnectionInput{
		ConnectionId: aws.String(connectionId),
		Data:         []byte(message),
	}

	_, err := client.PostToConnection(ctx, input)
	if err != nil {
		logError("Failed to send message", connectionId, err)
		return err
	}

	logSuccess("Message sent successfully", connectionId)
	return nil
}

// logError logs errors with connection ID
func logError(msg, connectionId string, err error) {
	log.Error(msg,
		zap.String("connectionId", connectionId),
		zap.Error(err),
	)
}

// logSuccess logs success with connection ID
func logSuccess(msg, connectionId string) {
	log.Info(msg,
		zap.String("connectionId", connectionId),
	)
}
