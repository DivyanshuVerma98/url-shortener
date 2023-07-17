package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DivyanshuVerma98/url-shortener/models"
	"github.com/aws/aws-lambda-go/events"
)

func CreateShortUrl(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Print("Inside CreateShortUrl")
	var requestbody models.CreateShortUrlRequest
	err := json.Unmarshal([]byte(request.Body), &requestbody)
	log.Println("Request Body", requestbody)
	if err != nil {
		return clientError(http.StatusUnprocessableEntity)
	}

	return events.APIGatewayProxyResponse{
        StatusCode: http.StatusOK,
        Body:       string(request.Body),
    }, nil
	// return events.APIGatewayProxyResponse{}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}
