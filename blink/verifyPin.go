package blink

import (
	"fmt"

	"github.com/rekram1-node/blinkgo/internal/client"
)

type VerifyResponse struct {
	Valid         bool   `json:"valid"`
	RequireNewPin bool   `json:"require_new_pin"`
	Message       string `json:"message"`
	Code          int    `json:"code"`
}

func (account *Account) VerifyPin(pin string) (*VerifyResponse, error) {
	verifyRes := &VerifyResponse{}
	c := client.New(account.AuthToken)
	url := fmt.Sprintf("https://rest-%s.immedia-semi.com/api/v4/account/%d/client/%d/pin/verify", account.Tier, account.ID, account.ClientID)
	body := map[string]string{
		"pin": pin,
	}

	resp, err := c.R().
		SetResult(verifyRes).
		SetBody(body).
		Post(url)

	if err != nil {
		return nil, err
	} else if !resp.IsSuccess() {
		return nil, fmt.Errorf("failed to verify pin, status code: %d, response: %s", resp.StatusCode(), resp.String())
	}

	return verifyRes, nil
}
