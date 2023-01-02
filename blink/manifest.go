package blink

import (
	"fmt"
	"time"

	"github.com/rekram1-node/blinkgo/internal/client"
)

type DeviceManifest struct {
	Account struct {
		ID                        int  `json:"id"`
		EmailVerified             bool `json:"email_verified"`
		EmailVerificationRequired bool `json:"email_verification_required"`
		AmazonAccountLinked       bool `json:"amazon_account_linked"`
	} `json:"account"`
	Networks []struct {
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Name      string    `json:"name"`
		TimeZone  string    `json:"time_zone"`
		Dst       bool      `json:"dst"`
		Armed     bool      `json:"armed"`
		LvSave    bool      `json:"lv_save"`
	} `json:"networks"`
	SyncModules []SyncModule  `json:"sync_modules"`
	Cameras     []Camera      `json:"cameras"`
	Sirens      []interface{} `json:"sirens"`
	Chimes      []interface{} `json:"chimes"`
	VideoStats  struct {
		Storage              int   `json:"storage"`
		AutoDeleteDays       int   `json:"auto_delete_days"`
		AutoDeleteDayOptions []int `json:"auto_delete_day_options"`
	} `json:"video_stats"`
	DoorbellButtons []interface{} `json:"doorbell_buttons"`
	Owls            []interface{} `json:"owls"`
	Doorbells       []interface{} `json:"doorbells"`
	AppUpdates      struct {
		Message         string `json:"message"`
		Code            int    `json:"code"`
		UpdateAvailable bool   `json:"update_available"`
		UpdateRequired  bool   `json:"update_required"`
	} `json:"app_updates"`
	DeviceLimits struct {
		Camera         int `json:"camera"`
		Chime          int `json:"chime"`
		Doorbell       int `json:"doorbell"`
		DoorbellButton int `json:"doorbell_button"`
		Owl            int `json:"owl"`
		Siren          int `json:"siren"`
		TotalDevices   int `json:"total_devices"`
	} `json:"device_limits"`
	WhatsNew struct {
		UpdatedAt int    `json:"updated_at"`
		URL       string `json:"url"`
	} `json:"whats_new"`
	Subscriptions struct {
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"subscriptions"`
	Entitlements struct {
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"entitlements"`
	TivLockEnable bool `json:"tiv_lock_enable"`
	TivLockStatus struct {
		Locked bool `json:"locked"`
	} `json:"tiv_lock_status"`
	Accessories struct {
		Storm []interface{} `json:"storm"`
		Rosie []interface{} `json:"rosie"`
	} `json:"accessories"`
}

type SyncModule struct {
	ID                     int       `json:"id"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
	Onboarded              bool      `json:"onboarded"`
	Status                 string    `json:"status"`
	Name                   string    `json:"name"`
	Serial                 string    `json:"serial"`
	FwVersion              string    `json:"fw_version"`
	Type                   string    `json:"type"`
	Subtype                string    `json:"subtype"`
	LastHb                 time.Time `json:"last_hb"`
	WifiStrength           int       `json:"wifi_strength"`
	NetworkID              int       `json:"network_id"`
	EnableTempAlerts       bool      `json:"enable_temp_alerts"`
	LocalStorageEnabled    bool      `json:"local_storage_enabled"`
	LocalStorageCompatible bool      `json:"local_storage_compatible"`
	LocalStorageStatus     string    `json:"local_storage_status"`
	Revision               string    `json:"revision"`
}

func (account *Account) GetManifest() (*DeviceManifest, error) {
	c := client.New(account.AuthToken)
	url := fmt.Sprintf("https://rest-%s.immedia-semi.com/api/v3/accounts/%d/homescreen", account.Tier, account.ID)
	manifest := &DeviceManifest{}

	resp, err := c.R().
		SetResult(manifest).
		Get(url)

	if err != nil {
		return nil, err
	} else if !resp.IsSuccess() {
		return nil, fmt.Errorf("failed to get device manifest, status code: %d, response: %s", resp.StatusCode(), resp.String())
	}

	account.SyncModules = &manifest.SyncModules

	return manifest, nil
}

type localStorageManifestIDResponse struct {
	ID        int `json:"id"`
	NetworkID int `json:"network_id"`
}

func (account *Account) GetLocalStorageManifestRequestID(networkID, syncModuleID int) (int, error) {
	localManifestIDResponse := &localStorageManifestIDResponse{}
	c := client.New(account.AuthToken)
	url := fmt.Sprintf("https://rest-%s.immedia-semi.com/api/v1/accounts/%d/networks/%d/sync_modules/%d/local_storage/manifest/request", account.Tier, account.ID, networkID, syncModuleID)

	resp, err := c.R().
		SetResult(localManifestIDResponse).
		Post(url)

	if err != nil {
		return 0, err
	} else if !resp.IsSuccess() {
		return 0, fmt.Errorf("failed to get local storage manifest request ID, status code: %d, response: %s", resp.StatusCode(), resp.String())
	}

	id := localManifestIDResponse.ID

	return id, nil
}

type LocalStorageManifest struct {
	Version    string `json:"version"`
	ManifestID string `json:"manifest_id"`
	Clips      []Clip `json:"clips"`
}

func (account *Account) GetLocalStorageManifest(networkID, syncModuleID, manifestRequestID int) (*LocalStorageManifest, error) {
	c := client.New(account.AuthToken)
	url := fmt.Sprintf("https://rest-%s.immedia-semi.com/api/v1/accounts/%d/networks/%d/sync_modules/%d/local_storage/manifest/request/%d", account.Tier, account.ID, networkID, syncModuleID, manifestRequestID)
	localManifest := &LocalStorageManifest{}

	resp, err := c.R().
		SetResult(localManifest).
		Get(url)

	if err != nil {
		return nil, err
	} else if !resp.IsSuccess() {
		return nil, fmt.Errorf("failed to get local storage manifest, status code: %d, response: %s", resp.StatusCode(), resp.String())
	}

	return localManifest, nil
}
