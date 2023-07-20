package responses

import (
	"net/http"

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

func ValidationError(msg string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       msg,
	}, nil
}
