package apigw

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type Response = events.APIGatewayProxyResponse

func InternalServerErrorResponse() Response {
	return Response{StatusCode: http.StatusInternalServerError}
}

//func BadRequestResponse() Response {
//	return Response{StatusCode: http.StatusBadRequest}
//}

func OkResponse() Response {
	return Response{StatusCode: http.StatusOK}
}
