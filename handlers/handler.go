package handlers

import (
	"encoding/json"
	"log"

	"github.com/DivyanshuVerma98/url-shortener/config"
	"github.com/DivyanshuVerma98/url-shortener/models"
	"github.com/DivyanshuVerma98/url-shortener/responses"
	"github.com/DivyanshuVerma98/url-shortener/utils"
	"github.com/aws/aws-lambda-go/events"
)

func CreateShortUrl(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Inside CreateShortUrl")
	var requestBody models.CreateShortUrlRequest
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		return responses.ValidationError("Invalid Data")
	}
	isValid, errMsg := requestBody.IsValid()
	if !isValid{
		return responses.ValidationError(errMsg)
	}
	var code string
	if requestBody.Alias!=""{
		code = requestBody.Alias
	}else{
		code = utils.GenerateCode(config.LengthOfCode)
	}
	expTime := utils.CreateExpTime(config.ExpTimeInDays)
	urlItem := models.UrlMapperItem{
		UserUrl:  requestBody.Url,
		ShortUrl: code,
		ExpTime:  expTime,
	}
	_ = models.CreateUrlMapperItem(&urlItem)
	shortUrl := config.DomainName + "/" + code
	data := map[string]string{"short_url": shortUrl}
	return responses.SuccessResponse(data)
}

func GetUserUrl(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Inside GetUserUrl")
	code := request.Path[1:]
	item := models.GetUrlMapperItem(code)
	if item.UserUrl == "" {
		return responses.NotFoundResponse()
	}
	return responses.RedirectPermanentlyResponse(item.UserUrl)
}
