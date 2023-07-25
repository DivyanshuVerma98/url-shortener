package models

import (
	"strconv"

	"github.com/DivyanshuVerma98/url-shortener/config"
	"github.com/DivyanshuVerma98/url-shortener/utils"
)

type CreateShortUrlRequest struct {
	Url   string `json:"url"`
	Alias string `json:"alias"`
}

func (request CreateShortUrlRequest) IsValid() (bool, string) {
	if request.Url == "" {
		return false, "URL is empty"
	}
	if !utils.ValidateUrl(request.Url) {
		return false, "Invalid URL"
	}
	if request.Alias != "" {
		if len(request.Alias) < config.MinLengthOfAlias {
			return false, "The Alias must have at least " + strconv.Itoa(config.MinLengthOfAlias) + " characters."
		} else if len(request.Alias) > config.MaxLengthOfAlias {
			return false, "The Alias must be less than or equal to " + strconv.Itoa(config.MaxLengthOfAlias) + " characters."
		} else if !utils.IsAlphanumeric(request.Alias) {
			return false, "The Alias should only contain letters and numbers."
		}
		// Checking is the given Alias already exists
		item := GetUrlMapperItem(request.Alias)
		if item.UserUrl != "" {
			return false, "Alias is not available."
		}
	}
	return true, ""
}

type UrlMapperItem struct {
	ShortUrl string `json:"short_url"`
	UserUrl  string `json:"user_url"`
	ExpTime  string `json:"exp_time"`
}

type ResponseItem struct {
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}
