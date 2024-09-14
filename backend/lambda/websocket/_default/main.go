package _default

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/norun9/HyBird/backend/lambda/websocket/lib/apigw"
	mydynamodb "github.com/norun9/HyBird/backend/lambda/websocket/lib/dynamodb"
	"github.com/norun9/HyBird/backend/pkg/log"
	"go.uber.org/zap"
	"os"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.InitLogger()
	defer log.Sync()

	// Get the API Gateway endpoint from environment variables
	endpoint, err := getEnv("API_GATEWAY_ENDPOINT")
	if err != nil {
		logErrorAndReturn(err, "API_GATEWAY_ENDPOINT is not set")
		return apigw.InternalServerErrorResponse(), nil
	}

	// Log the received message from WebSocket
	message := req.Body
	logMessageReceived(req, message)

	// Fetch all WebSocket connections from DynamoDB
	svc := mydynamodb.NewDBSession(ctx)
	connections, err := mydynamodb.GetAllConnections(ctx, svc)
	if err != nil {
		logErrorAndReturn(err, "Failed to retrieve connections from DynamoDB")
		return apigw.InternalServerErrorResponse(), err
	}

	// Initialize API Gateway Management Client
	apigwCli, err := apigw.NewAPIGatewayClient(ctx, endpoint)
	if err != nil {
		logErrorAndReturn(err, "Failed to initialize API Gateway client")
		return apigw.InternalServerErrorResponse(), err
	}

	// Send the message to all connected clients
	for _, conn := range connections {
		connectionId := conn.ConnectionId
		err := apigw.SendMessageToConnection(ctx, apigwCli, connectionId, message)
		if err != nil {
			log.Error("Failed to send message to connection",
				zap.String("connectionId", connectionId),
				zap.Error(err))
		} else {
			log.Info("Message sent successfully to connection",
				zap.String("connectionId", connectionId))
		}
	}

	return apigw.OkResponse(), nil
}

// getEnv fetches the value of the environment variable or returns an error if not set
func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%s environment variable is not set", key)
	}
	return value, nil
}

// logErrorAndReturn logs an error message and returns it
func logErrorAndReturn(err error, message string) {
	log.Error(message, zap.Error(err))
}

// logMessageReceived logs the received WebSocket message and request information
func logMessageReceived(req *events.APIGatewayWebsocketProxyRequest, message string) {
	log.Info("Received message on WebSocket",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID),
		zap.String("message", message))
}
