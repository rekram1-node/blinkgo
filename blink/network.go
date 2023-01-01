package blink

import (
	"fmt"

	"github.com/rekram1-node/blinkgo/internal/client"
)

type NetworkList struct {
	RangeDays int `json:"range_days"`
	Reference struct {
		Usage int `json:"usage"`
	} `json:"reference"`
	Networks []Network `json:"networks"`
}

type Network struct {
	ID      int    `json:"network_id"`
	Name    string `json:"name"`
	Cameras []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Usage       int    `json:"usage"`
		LvSeconds   int    `json:"lv_seconds"`
		ClipSeconds int    `json:"clip_seconds"`
	} `json:"cameras"`
}

func (account *Account) GetListOfNetworks() error {
	networkList := &NetworkList{}
	c := client.New(account.AuthToken)
	url := fmt.Sprintf("https://rest-%s.immedia-semi.com/api/v1/camera/usage", account.Tier)

	resp, err := c.R().
		SetResult(networkList).
		Get(url)

	if err != nil {
		return err
	} else if !resp.IsSuccess() {
		return fmt.Errorf("failed to get list of networks, status code: %d, response: %s", resp.StatusCode(), resp.String())
	}

	account.Networks = &networkList.Networks

	return nil
}
