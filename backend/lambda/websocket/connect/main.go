package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/norun9/HyBird/backend/lambda/websocket/lib/apigw"
	mydynamodb "github.com/norun9/HyBird/backend/lambda/websocket/lib/dynamodb"
	"github.com/norun9/HyBird/backend/pkg/log"
	"go.uber.org/zap"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.InitLogger()

	defer log.Sync()

	connectionId := req.RequestContext.ConnectionID

	log.Info("websocket connected",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", connectionId))

	item := mydynamodb.Item{
		ConnectionId: connectionId,
	}

	svc := mydynamodb.NewDBSession(ctx)

	err := item.PutConnectionId(ctx, svc)
	if err != nil {
		return apigw.InternalServerErrorResponse(), err
	}

	return apigw.OkResponse(), nil
}
