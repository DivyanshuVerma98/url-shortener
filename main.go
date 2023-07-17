package main

import (
	"context"
	"fmt"
	"log"

	"github.com/DivyanshuVerma98/url-shortener/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("HTTP Method", request.HTTPMethod)
	log.Println("Path", request.Path)
	if request.Path == "/createurl" && request.HTTPMethod == "POST" {
		return handlers.CreateShortUrl(request)
	}
	return events.APIGatewayProxyResponse{Body: "Given path not found."}, nil
}

func main() {
	fmt.Println("This is my UrlShortner Service")
	lambda.Start(handleRequest)
}

// build command
// GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go

/*
type APIGatewayProxyRequest struct {
    Resource              string                        `json:"resource"` // The resource path defined in API Gateway
    Path                  string                        `json:"path"`     // The url path for the caller
    HTTPMethod            string                        `json:"httpMethod"`
    Headers               map[string]string             `json:"headers"`
    QueryStringParameters map[string]string             `json:"queryStringParameters"`
    PathParameters        map[string]string             `json:"pathParameters"`
    StageVariables        map[string]string             `json:"stageVariables"`
    RequestContext        APIGatewayProxyRequestContext `json:"requestContext"`
    Body                  string                        `json:"body"`
    IsBase64Encoded       bool                          `json:"isBase64Encoded,omitempty"`
}

type APIGatewayProxyResponse struct {
    StatusCode      int               `json:"statusCode"`
    Headers         map[string]string `json:"headers"`
    Body            string            `json:"body"`
    IsBase64Encoded bool              `json:"isBase64Encoded,omitempty"`
}
*/
