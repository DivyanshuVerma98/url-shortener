package utils

import (
	"net/url"
	"unicode"
)

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

func IsAlphanumeric(givenStr string) bool{
	for _, val := range(givenStr){
		if !unicode.IsLetter(val) && !unicode.IsNumber(val){
			return false
		}
	}
	return true
}