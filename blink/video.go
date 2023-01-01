package blink

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rekram1-node/blinkgo/internal/client"
)

type Clip struct {
	ID         string    `json:"id"`
	Size       string    `json:"size"`
	CameraName string    `json:"camera_name"`
	CreatedAt  time.Time `json:"created_at"`
}

type VideoEvents struct {
	Limit        int           `json:"limit"`
	PurgeID      int64         `json:"purge_id"`
	RefreshCount int           `json:"refresh_count"`
	Media        []interface{} `json:"media"`
}

func (account *Account) GetVideoEvents(pages int) (*VideoEvents, error) {
	videoEvents := &VideoEvents{}
	c := client.New(account.AuthToken)
	url := fmt.Sprintf("https://rest-%s.immedia-semi.com/api/v1/accounts/%d/media/changed?since=2015-04-19T23:11:20+0000&page=%d", account.Tier, account.ID, pages)

	resp, err := c.R().
		SetResult(videoEvents).
		Get(url)

	if err != nil {
		return nil, err
	} else if !resp.IsSuccess() {
		return nil, fmt.Errorf("failed to get video events, status code: %d, response: %s", resp.StatusCode(), resp.String())
	}

	return videoEvents, nil
}

func (account *Account) GetClipIDs(networkID, syncModuleID, requestID int) (*[]Clip, error) {
	manifest, err := account.GetLocalStorageManifest(networkID, syncModuleID, requestID)

	if err != nil {
		return nil, err
	}

	return &manifest.Clips, nil
}

func (account *Account) DownloadVideoByClipID(networkID, syncModuleID int, manifestID, clipID, fileName string) error {
	// this filename should have the ".mp4" in it
	out, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer out.Close()

	c := client.New(account.AuthToken)
	url := fmt.Sprintf("https://rest-%s.immedia-semi.com/api/v1/accounts/%d/networks/%d/sync_modules/%d/local_storage/manifest/%s/clip/request/%s", account.Tier, account.ID, networkID, syncModuleID, manifestID, clipID)

	resp, err := c.R().
		SetDoNotParseResponse(true). // necessary to read file
		Get(url)

	if err != nil {
		return err
	} else if !resp.IsSuccess() {
		return fmt.Errorf("failed to get request download clip by ID: %s, status code: %d, response: %s", clipID, resp.StatusCode(), string(resp.Body()))
	}

	// copy mp4 video into file
	if _, err = io.Copy(out, resp.RawBody()); err != nil {
		return err
	}

	return resp.RawBody().Close()
}

type UploadResponse struct {
	ID        int `json:"id"`
	NetworkID int `json:"network_id"`
}

// For local storage use
func (account *Account) RequestUploadByClipID(networkID, syncModuleID int, manifestID, clipID string) (*UploadResponse, error) {
	uploadRes := &UploadResponse{}
	c := client.New(account.AuthToken)
	url := fmt.Sprintf("https://rest-%s.immedia-semi.com/api/v1/accounts/%d/networks/%d/sync_modules/%d/local_storage/manifest/%s/clip/request/%s", account.Tier, account.ID, networkID, syncModuleID, manifestID, clipID)

	resp, err := c.R().
		SetResult(uploadRes).
		Post(url)

	if err != nil {
		return nil, err
	} else if !resp.IsSuccess() {
		return nil, fmt.Errorf("failed to get request upload clip by ID: %s, status code: %d, response: %s", clipID, resp.StatusCode(), resp.String())
	}

	return uploadRes, nil
}

type Media struct {
	Limit        int64 `json:"limit"`
	PurgeID      int64 `json:"purge_id"`
	RefreshCount int64 `json:"refresh_count"`
	Media        []struct {
		AdditionalDevices []interface{} `json:"additional_devices"`
		CreatedAt         string        `json:"created_at"`
		Deleted           bool          `json:"deleted"`
		Device            string        `json:"device"`
		DeviceID          int64         `json:"device_id"`
		DeviceName        string        `json:"device_name"`
		ID                int64         `json:"id"`
		Media             string        `json:"media"`
		NetworkID         int64         `json:"network_id"`
		NetworkName       string        `json:"network_name"`
		Partial           bool          `json:"partial"`
		Source            string        `json:"source"`
		Thumbnail         string        `json:"thumbnail"`
		TimeZone          string        `json:"time_zone"`
		Type              string        `json:"type"`
		UpdatedAt         string        `json:"updated_at"`
		Watched           bool          `json:"watched"`
	} `json:"media"`
	Videos []Video `json:"videos"`
}

type Video struct {
	AccountID       int64       `json:"account_id"`
	Address         string      `json:"address"`
	CameraID        int64       `json:"camera_id"`
	CameraName      string      `json:"camera_name"`
	CreatedAt       string      `json:"created_at"`
	Deleted         bool        `json:"deleted"`
	Description     string      `json:"description"`
	Encryption      string      `json:"encryption"`
	EncryptionKey   interface{} `json:"encryption_key"`
	EventID         interface{} `json:"event_id"`
	ID              int64       `json:"id"`
	Length          int64       `json:"length"`
	Locked          bool        `json:"locked"`
	NetworkID       int64       `json:"network_id"`
	NetworkName     string      `json:"network_name"`
	Partial         bool        `json:"partial"`
	Ready           bool        `json:"ready"`
	Size            int64       `json:"size"`
	StorageLocation string      `json:"storage_location"`
	Thumbnail       string      `json:"thumbnail"`
	TimeZone        string      `json:"time_zone"`
	UpdatedAt       string      `json:"updated_at"`
	UploadTime      int64       `json:"upload_time"`
	Viewed          string      `json:"viewed"`
}

func (account *Account) GetMedia(daysBack, pageNum int) (*Media, error) {
	currTimeStamp := time.Now().UTC().Add(-time.Hour * 24 * time.Duration(daysBack)).Format("2006-01-02T15:04:05-0700")
	media := &Media{}
	c := client.New(account.AuthToken)
	url := fmt.Sprintf("/api/v1/accounts/%d/media/changed?since=%s&page=%d", account.ID, currTimeStamp, pageNum)

	resp, err := c.R().
		SetResult(media).
		Post(url)

	if err != nil {
		return nil, err
	} else if !resp.IsSuccess() {
		return nil, fmt.Errorf("failed to get media list, status code: %d, response: %s", resp.StatusCode(), resp.String())
	}

	return media, nil
}

func (account *Account) GetVideos(daysBack, pageNum int) (*[]Video, error) {
	media, err := account.GetMedia(daysBack, pageNum)

	if err != nil {
		return nil, err
	}

	return &media.Videos, nil
}

func (account *Account) DownloadVideosByPage() {

}
