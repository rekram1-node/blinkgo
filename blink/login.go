package blink

import (
	"fmt"

	"github.com/rekram1-node/blinkgo/internal/client"
)

type LoginResponse struct {
	Account               AccountInformation `json:"account"`
	Auth                  Auth               `json:"auth"`
	Phone                 Phone              `json:"phone"`
	Verification          Verification       `json:"verification"`
	LockoutTimeRemaining  int                `json:"lockout_time_remaining"`
	ForcePasswordReset    bool               `json:"force_password_reset"`
	AllowPinResendSeconds int                `json:"allow_pin_resend_seconds"`
}

type AccountInformation struct {
	AccountID                   int    `json:"account_id"`
	UserID                      int    `json:"user_id"`
	ClientID                    int    `json:"client_id"`
	ClientTrusted               bool   `json:"client_trusted"`
	NewAccount                  bool   `json:"new_account"`
	Tier                        string `json:"tier"`
	Region                      string `json:"region"`
	AccountVerificationRequired bool   `json:"account_verification_required"`
	PhoneVerificationRequired   bool   `json:"phone_verification_required"`
	ClientVerificationRequired  bool   `json:"client_verification_required"`
	RequireTrustClientDevice    bool   `json:"require_trust_client_device"`
	CountryRequired             bool   `json:"country_required"`
	VerificationChannel         string `json:"verification_channel"`
	User                        User   `json:"user"`
	AmazonAccountLinked         bool   `json:"amazon_account_linked"`
}

type User struct {
	UserID  int    `json:"user_id"`
	Country string `json:"country"`
}

type Auth struct {
	Token string `json:"token"`
}

type Phone struct {
	Number             string `json:"number"`
	Last4Digits        string `json:"last_4_digits"`
	CountryCallingCode string `json:"country_calling_code"`
	Valid              bool   `json:"valid"`
}

type Verification struct {
	Email struct {
		Required bool `json:"required"`
	} `json:"email"`
	Phone struct {
		Required bool   `json:"required"`
		Channel  string `json:"channel"`
	} `json:"phone"`
}

func (account *Account) Login() (*LoginResponse, error) {
	body := map[string]interface{}{
		"reauth":   true,
		"email":    account.Email,
		"password": account.Password,
	}

	loginResp := &LoginResponse{}
	c := client.Default()
	url := client.LoginURL

	resp, err := c.R().
		SetBody(body).
		SetResult(loginResp).
		Post(url)

	if err != nil {
		return nil, err
	} else if !resp.IsSuccess() {
		return nil, fmt.Errorf("failed to login, status code: %d, response: %s", resp.StatusCode(), resp.String())
	}

	account.updateAccountInfo(loginResp)

	return loginResp, nil
}

func (account *Account) updateAccountInfo(resp *LoginResponse) {
	account.Tier = resp.Account.Tier
	account.ID = resp.Account.AccountID
	account.ClientID = resp.Account.ClientID
	account.AuthToken = resp.Auth.Token
}

func (account *Account) RefreshToken() (*LoginResponse, error) {
	// this method is incomplete
	// this just for readability of code I created RefreshToken
	return account.Login()
}
