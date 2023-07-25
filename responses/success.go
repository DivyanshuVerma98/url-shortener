package responses

import (
	"encoding/json"
	"net/http"

	"github.com/DivyanshuVerma98/url-shortener/models"
	"github.com/aws/aws-lambda-go/events"
)

func SuccessResponse(message string, data map[string]string) (events.APIGatewayProxyResponse, error) {
	responseItem := models.ResponseItem{
		Message: message,
		Data:    data,
	}
	response_json, _ := json.Marshal(responseItem)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json"},
		Body: string(response_json),
	}, nil
}

func RedirectPermanentlyResponse(redirect_url string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMovedPermanently,
		Headers: map[string]string{
			"Location": redirect_url,
		},
	}, nil
}

func RenderHTMLResponse(response string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "text/html; charset=UTF-8",
		},
		Body: response,
	}, nil
}
