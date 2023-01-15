package client

import (
	"github.com/go-resty/resty/v2"
)

const (
	BaseURL  = "https://rest-prod.immedia-semi.com"
	LoginURL = BaseURL + "/api/v5/account/login"
)

func Default() *resty.Client {
	return resty.New().SetHeader("Content-Type", "application/json")
}

func New(token string) *resty.Client {
	return Default().SetHeader("token-auth", token)
}
