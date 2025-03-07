package routes

import (
	"lambda_url_shortener/internal/handlers"

	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func LambdaHandler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "POST":
		return handlers.HandleShortenURL(req)
	case "GET":
		return handlers.HandleRedirect(req)
	default:
		return events.APIGatewayProxyResponse{StatusCode: http.StatusMethodNotAllowed}, nil
	}
}
