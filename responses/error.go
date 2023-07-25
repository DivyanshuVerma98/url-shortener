package responses

import (
	"encoding/json"
	"net/http"

	"github.com/DivyanshuVerma98/url-shortener/models"
	"github.com/aws/aws-lambda-go/events"
)

func NotFoundResponse() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Body:       "Not Found",
	}, nil
}

func ServerError() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func ValidationError(message string, data map[string]string) (events.APIGatewayProxyResponse, error) {
	responseItem := models.ResponseItem{
		Message: message,
		Data:    data,
	}
	response_json, _ := json.Marshal(responseItem)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json"},
		Body: string(response_json),
	}, nil
}
