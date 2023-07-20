package responses

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func SuccessResponse(response interface{}) (events.APIGatewayProxyResponse, error) {
	response_json, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(response_json),
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
