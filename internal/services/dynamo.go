package services

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type URLMapping struct {
	ShortCode string `dynamodbav:"short_code"`
	LongURL   string `dynamodbav:"long_url"`
	Expiry    int64  `dynamodbav:"expiry"`
}

var (
	dbClient  *dynamodb.Client
	tableName string
)

// InitDynamoDB initializes the DynamoDB client
func InitDynamoDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to load AWS configuration: %v", err)
	}

	dbClient = dynamodb.NewFromConfig(cfg)
	if dbClient == nil {
		return fmt.Errorf("failed to create DynamoDB client")
	}

	// Set DYNAMODB_TABLE in Lambda environment
	tableName = os.Getenv("DYNAMODB_TABLE")

	return nil
}

// StoreShortURL inserts a short URL mapping into DynamoDB
func StoreShortURL(shortCode, longURL string, expiry int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	urlItem := URLMapping{
		ShortCode: shortCode,
		LongURL:   longURL,
		Expiry:    expiry,
	}

	item, err := attributevalue.MarshalMap(urlItem)
	if err != nil {
		return fmt.Errorf("failed to marshal URLMapping: %v", err)
	}

	_, err = dbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
		// Prevent overwriting existing short codes
		ConditionExpression: aws.String("attribute_not_exists(short_code)"),
	})

	if err != nil {
		return fmt.Errorf("failed to put item into DynamoDB: %v", err)
	}

	return nil
}

// GetLongURL retrieves the original URL from DynamoDB based on short code
func GetLongURL(shortCode string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Fetch URL mapping
	result, err := dbClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"short_code": &types.AttributeValueMemberS{Value: shortCode},
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to get item from DynamoDB: %v", err)
	}

	// Check if result is empty
	if result.Item == nil || len(result.Item) == 0 {
		return "", fmt.Errorf("short URL not found")
	}

	var urlMapping URLMapping
	err = attributevalue.UnmarshalMap(result.Item, &urlMapping)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal item from DynamoDB: %v", err)
	}

	return urlMapping.LongURL, nil
}
