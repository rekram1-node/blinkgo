package client

import (
	"github.com/go-resty/resty/v2"
)

const (
	BaseURL  = "https://rest-prod.immedia-semi.com"
	LoginURL = BaseURL + "/api/v5/account/login"
)

type DefaultClient struct {
	Client *resty.Client
}

func Default() *resty.Client {
	c := resty.New().
		SetHeader("Content-Type", "application/json")

	return c
}

func New(token string) *resty.Client {
	c := resty.New().
		SetHeader("Content-Type", "application/json").
		SetHeader("token-auth", token)

	return c
}
