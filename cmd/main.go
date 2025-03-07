package main

import (
	"lambda_url_shortener/internal/routes"
	"lambda_url_shortener/internal/services"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// Initialize DynamoDB client
	err := services.InitDynamoDB()
	if err != nil {
		log.Fatalf("Failed to initialize DynamoDB: %v", err)
	}

	// Start Lambda handler
	lambda.Start(routes.LambdaHandler)
}
