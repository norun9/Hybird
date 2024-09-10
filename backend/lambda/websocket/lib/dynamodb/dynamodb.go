package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/norun9/HyBird/backend/lambda/websocket/lib/dynamodb/constants"
	"github.com/norun9/HyBird/backend/pkg/log"
	"go.uber.org/zap"
)

type Item struct {
	ConnectionId string
}

func (i Item) PutConnectionId(svc *dynamodb.DynamoDB) error {
	itemAttributes, err := dynamodbattribute.MarshalMap(i)
	if err != nil {
		log.Error("Got error marshalling new ConnectionId: %s", zap.Error(err))
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      itemAttributes,
		TableName: aws.String(constants.TableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Error("Got error calling PutItem using DynamoDB", zap.Error(err))
		return err
	}

	return nil
}

func NewDBSession() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{}))
	return dynamodb.New(sess)
}
