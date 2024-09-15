package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/norun9/HyBird/backend/lambda/websocket/lib/apigw"
	mydynamodb "github.com/norun9/HyBird/backend/lambda/websocket/lib/dynamodb"
	"github.com/norun9/HyBird/backend/pkg/log"
	"go.uber.org/zap"
	"net/url"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.InitLogger()
	defer log.Sync()

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
	callbackURL := url.URL{
		Scheme: "https",
		Host:   req.RequestContext.DomainName,
		Path:   req.RequestContext.Stage,
	}
	apigwCli, err := apigw.NewAPIGatewayClient(ctx, callbackURL)
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
