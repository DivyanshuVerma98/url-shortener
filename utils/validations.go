package utils

import "net/url"

func ValidateUrl(inputUrl string) bool {
	parsedUrl, err := url.Parse(inputUrl)
	if err != nil {
		return false
	}
	if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		return false
	}
	return true
}
