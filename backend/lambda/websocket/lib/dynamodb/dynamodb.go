package dynamodb

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/norun9/HyBird/backend/pkg/log"
	"go.uber.org/zap"
	"os"
)

type Item struct {
	ConnectionId string
}

func GetAllConnections(ctx context.Context, svc *dynamodb.Client) ([]Item, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("Connections"),
	}

	result, err := svc.Scan(ctx, input)
	if err != nil {
		log.Error("failed to scan table", zap.Error(err))
		return nil, err
	}

	var items []Item
	err = attributevalue.UnmarshalListOfMaps(result.Items, &items)
	if err != nil {
		log.Error("failed to unmarshal items", zap.Error(err))
		return nil, err
	}

	return items, nil
}

func (i Item) PutConnectionId(ctx context.Context, svc *dynamodb.Client) error {
	itemAttributes, err := attributevalue.MarshalMap(i)
	if err != nil {
		log.Error("Got error marshalling new ConnectionId: %s", zap.Error(err))
		return err
	}

	tableName := os.Getenv("TABLE_NAME")

	input := &dynamodb.PutItemInput{
		Item:      itemAttributes,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(ctx, input)
	if err != nil {
		log.Error("Got error calling PutItem using DynamoDB", zap.Error(err))
		return err
	}

	return nil
}

func DeleteConnectionId(ctx context.Context, svc *dynamodb.Client, connectionId string) error {
	tableName := os.Getenv("TABLE_NAME")
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"ConnectionId": &types.AttributeValueMemberS{Value: connectionId},
		},
	}

	_, err := svc.DeleteItem(ctx, input)
	if err != nil {
		log.Error("Got error calling DeleteItem using DynamoDB", zap.Error(err))
		return err
	}
	return nil
}

func NewDBSession(ctx context.Context) *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("ap-northeast-1"))
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}
	return dynamodb.NewFromConfig(cfg)
}
