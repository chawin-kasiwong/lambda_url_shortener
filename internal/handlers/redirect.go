package handlers

import (
	"lambda_url_shortener/internal/services"

	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRedirect(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	shortCode := req.PathParameters["shortcode"]
	if shortCode == "" {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: `{"error": "Shortcode is required"}`}, nil
	}

	longURL, err := services.GetLongURL(shortCode)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusNotFound, Body: `{"error": "Short URL not found"}`}, nil
	}

	headers := map[string]string{"Location": longURL}
	return events.APIGatewayProxyResponse{StatusCode: http.StatusFound, Headers: headers}, nil
}
