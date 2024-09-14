package apigw

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type Response = events.APIGatewayProxyResponse

func InternalServerErrorResponse() Response {
	return Response{StatusCode: http.StatusInternalServerError}
}

func OkResponse() Response {
	return Response{StatusCode: http.StatusOK}
}
