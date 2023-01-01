package blink

import (
	"time"
)

type Camera struct {
	ID        int           `json:"id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Name      string        `json:"name"`
	Serial    string        `json:"serial"`
	FwVersion string        `json:"fw_version"`
	Type      string        `json:"type"`
	Enabled   bool          `json:"enabled"`
	Thumbnail string        `json:"thumbnail"`
	Status    string        `json:"status"`
	Battery   string        `json:"battery"`
	UsageRate bool          `json:"usage_rate"`
	NetworkID int           `json:"network_id"`
	Issues    []interface{} `json:"issues"`
	Signals   struct {
		Lfr     int `json:"lfr"`
		Wifi    int `json:"wifi"`
		Temp    int `json:"temp"`
		Battery int `json:"battery"`
	} `json:"signals"`
	LocalStorageEnabled    bool        `json:"local_storage_enabled"`
	LocalStorageCompatible bool        `json:"local_storage_compatible"`
	Snooze                 bool        `json:"snooze"`
	SnoozeTimeRemaining    interface{} `json:"snooze_time_remaining"`
	Revision               string      `json:"revision"`
	Color                  string      `json:"color"`
}

func (account *Account) GetCameras() ([]Camera, error) {
	manifest, err := account.GetManifest()

	if err != nil {
		return nil, err
	}

	return manifest.Cameras, nil
}
