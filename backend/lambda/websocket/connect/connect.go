package connect

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

func handleRequest(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.InitLogger()

	defer log.Sync()

	svc := mydynamodb.NewDBSession()

	connectionId := req.RequestContext.ConnectionID

	log.Info("websocket connected!",
		zap.String("requestId", req.RequestContext.RequestID),
		zap.String("connectionId", connectionId))

	item := mydynamodb.Item{
		ConnectionId: connectionId,
	}

	err := item.PutConnectionId(svc)
	if err != nil {
		return apigw.InternalServerErrorResponse(), err
	}

	return apigw.OkResponse(), nil
}
