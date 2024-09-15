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

	message := req.Body

	// logs the received WebSocket message and request information
	log.Info("Received message on WebSocket",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", req.RequestContext.ConnectionID),
		zap.String("message", message))

	// Initialize API Gateway Management Client
	callbackURL := url.URL{
		Scheme: "https",
		Host:   req.RequestContext.DomainName,
		Path:   req.RequestContext.Stage,
	}
	apigwCli, err := apigw.NewAPIGatewayClient(ctx, callbackURL)
	if err != nil {
		log.Error("Failed to initialize API Gateway client", zap.Error(err))
		return apigw.InternalServerErrorResponse(), err
	}

	svc := mydynamodb.NewDBSession(ctx)
	connections, err := mydynamodb.GetAllConnections(ctx, svc)
	if err != nil {
		return apigw.InternalServerErrorResponse(), err
	}

	// Send the message to all connected clients
	for _, conn := range connections {
		err = apigw.SendMessageToConnection(ctx, apigwCli, conn.ConnectionId, message)
		if err != nil {
			return apigw.InternalServerErrorResponse(), err
		} else {
			log.Info("Message sent successfully to connection",
				zap.String("connectionId", conn.ConnectionId))
		}
	}

	return apigw.OkResponse(), nil
}
