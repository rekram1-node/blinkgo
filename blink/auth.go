package blink

import (
	"fmt"

	"github.com/rekram1-node/blinkgo/internal/client"
)

type LoginResponse struct {
	Account struct {
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
		User                        struct {
			UserID  int    `json:"user_id"`
			Country string `json:"country"`
		} `json:"user"`
		AmazonAccountLinked bool `json:"amazon_account_linked"`
	} `json:"account"`
	Auth struct {
		Token string `json:"token"`
	} `json:"auth"`
	Phone struct {
		Number             string `json:"number"`
		Last4Digits        string `json:"last_4_digits"`
		CountryCallingCode string `json:"country_calling_code"`
		Valid              bool   `json:"valid"`
	} `json:"phone"`
	Verification struct {
		Email struct {
			Required bool `json:"required"`
		} `json:"email"`
		Phone struct {
			Required bool   `json:"required"`
			Channel  string `json:"channel"`
		} `json:"phone"`
	} `json:"verification"`
	LockoutTimeRemaining  int  `json:"lockout_time_remaining"`
	ForcePasswordReset    bool `json:"force_password_reset"`
	AllowPinResendSeconds int  `json:"allow_pin_resend_seconds"`
}

// Completes a login request, might require a verify pin authentication as well
func (account *Account) Login() (*LoginResponse, error) {
	body := map[string]interface{}{
		"email":     account.Email,
		"password":  account.Password,
		"unique_id": account.UniqueID,
		"reauth":    true,
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

// Fetches new auth token
func (account *Account) RefreshToken() (*LoginResponse, error) {
	// this just for readability of code I created RefreshToken
	return account.Login()
}

type VerifyResponse struct {
	Valid         bool   `json:"valid"`
	RequireNewPin bool   `json:"require_new_pin"`
	Message       string `json:"message"`
	Code          int    `json:"code"`
}

// Verifies pin sent via email or sms
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

// Logs user out
func (account *Account) Logout() error {
	c := client.New(account.AuthToken)
	url := fmt.Sprintf("https://rest-%s.immedia-semi.com/api/v4/account/%d/client/%d/logout", account.Tier, account.ID, account.ClientID)
	resp, err := c.R().Post(url)

	if err != nil {
		return err
	} else if !resp.IsSuccess() {
		return fmt.Errorf("failed to logout, status code: %d, response: %s", resp.StatusCode(), resp.String())
	}

	return nil
}
