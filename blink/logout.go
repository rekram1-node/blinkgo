package blink

import (
	"fmt"

	"github.com/rekram1-node/blinkgo/internal/client"
)

func (account *Account) Logout() error {
	// logoutRes := &LogoutRes{}
	c := client.New(account.AuthToken)
	url := fmt.Sprintf("https://rest-%s.immedia-semi.com/api/v4/account/%d/client/%d/logout", account.Tier, account.ID, account.ClientID)
	resp, err := c.R().
		// SetResult(logoutRes).
		Post(url)

	if err != nil {
		return err
	} else if !resp.IsSuccess() {
		return fmt.Errorf("failed to logout, status code: %d, response: %s", resp.StatusCode(), resp.String())
	}

	// TODO convert to object
	fmt.Println(resp.String())
	// return logoutRes, nil
	return nil
}
