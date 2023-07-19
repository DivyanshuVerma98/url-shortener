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
	var requestBody models.CreateShortUrlRequest
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	log.Println("Request Body", requestBody)
	if err != nil {
		return clientError(http.StatusUnprocessableEntity)
	}
	urlItem := models.UrlMapperItem{
		UserUrl:  requestBody.Url,
		ShortUrl: requestBody.Url,
	}
	_ = models.CreateUrlMapperItem(&urlItem)

	response := map[string]string{"short_url": requestBody.Url}
	response_json, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(response_json),
	}, nil
}

func GetUserUrl(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	shortUrl := request.QueryStringParameters["url"]
	item := models.GetUrlMapperItem(shortUrl)
	if item.UserUrl == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Not Found",
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMovedPermanently,
		Headers: map[string]string{
			"Location": item.UserUrl,
		},
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}
