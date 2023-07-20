package handlers

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/DivyanshuVerma98/url-shortener/config"
	"github.com/DivyanshuVerma98/url-shortener/models"
	"github.com/DivyanshuVerma98/url-shortener/responses"
	"github.com/DivyanshuVerma98/url-shortener/utils"
	"github.com/aws/aws-lambda-go/events"
)

func CreateShortUrl(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Print("Inside CreateShortUrl")
	var requestBody models.CreateShortUrlRequest
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		return responses.ValidationError("Invalid Data")
	}
	if requestBody.Url == "" || !utils.ValidateUrl(requestBody.Url) {
		return responses.ValidationError("Invalid or empty URL given")
	}
	code := utils.GenerateCode(config.LengthOfCode)
	urlItem := models.UrlMapperItem{
		UserUrl:  requestBody.Url,
		ShortUrl: code,
	}
	_ = models.CreateUrlMapperItem(&urlItem)
	shortUrl := config.DomainName + "/getuserurl/" + code
	data := map[string]string{"short_url": shortUrl}
	return responses.SuccessResponse(data)
}

func GetUserUrl(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("request", request)
	index := strings.Index(request.Path, "getuserurl/")
	code := request.Path[index+len("getuserurl/"):]
	item := models.GetUrlMapperItem(code)
	if item.UserUrl == "" {
		return responses.NotFoundResponse()
	}
	return responses.RedirectPermanentlyResponse(item.UserUrl)
}
