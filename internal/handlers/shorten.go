package handlers

import (
	"encoding/json"
	"lambda_url_shortener/internal/dto"
	"lambda_url_shortener/internal/services"
	"lambda_url_shortener/pkg/utils"

	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func HandleShortenURL(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var shortenReq dto.ShortenRequest

	err := json.Unmarshal([]byte(req.Body), &shortenReq)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: `{"error": "Invalid request"}`}, nil
	}

	if err := utils.ValidateRequest(shortenReq); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Body: `{"error": "Invalid URL format"}`}, nil
	}

	shortCode := utils.GenerateShortCode()

	err = services.StoreShortURL(shortCode, shortenReq.LongURL, shortenReq.Expiry)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Body: `{"error": "Failed to store URL"}`}, nil
	}

	response := dto.ShortenResponse{
		ShortURL: "https://" + req.Headers["Host"] + "/" + shortCode,
	}
	respJSON, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: string(respJSON)}, nil
}
