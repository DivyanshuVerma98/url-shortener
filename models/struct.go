package models

type CreateShortUrlRequest struct {
	Url string `json:"url"`
}

type UrlMapperItem struct {
	ShortUrl string `json:"short_url"`
	UserUrl  string `json:"user_url"`
	ExpTime  string `json:"exp_time"`
}
