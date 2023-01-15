# BlinkGO Quick Start

## Contents
- [Getting Started](#getting-started)
- [Authentication](#authentication)
    - [Login](#login)
    - [Verify Pin](#verify-pin)
    - [Refresh Token](#refresh-token)
    - [Logout](#logout)
- [Network](#network)
    - [List Networks](#list-networks)
- [Camera](#camera)
    - [List Cameras](#list-cameras)
- [Sync Module](#sync-module)
    - [List Sync Modules](#list-sync-modules)
- [System](#system)
    - [Get Manifest](#get-manifest)
- [Videos (Cloud)](#videos-cloud)
    - [List Videos](#list-videos-cloud)
    - [Downloading Videos](#download-videos-cloud)
- [Videos (Local Storage)](#videos-local-storage)
    - [List Videos](#list-videos-local-storage)
    - [Downloading Videos](#download-videos-local-storage)

## Getting Started

### Create New Account
```go
email := "example@example.com"
password := "PLEASE_DON'T_PLAINTEXT_REAL_PASSWORDS"

// returns account object with email, password, and a unique uuid
// this account object will now be able to do login
account := blink.NewAccount(email, password)
```

## Authentication

### Login
Note: you must verify pin after login before trying any other methods

```go
// this returns a login response that you can use
// however, it is unneccessary for this example
loginResponse, err := account.Login()
if err != nil {
    log.Fatal(err)
}
```

### Verify Pin
```go
account := blink.NewAccount("someUsername", "somepassword")
loginResp, err := account.Login()

if err != nil {
    log.Fatal(err)
}

// if blink wants a 2FA verification you must use the 
// verify pin operation
// 2FA is not always required but it is typically
// required first time on new device
if loginResp.Account.AccountVerificationRequired {
    fmt.Print("Enter Pin: ")
    var pin string
    fmt.Scanln(&pin)

    // this returns a verify pin response that you can use
    // however, it is unneccessary for this example
    if verifyPinRes, err := account.VerifyPin(pin); err != nil {
        log.Fatal(err)
    }
}
```

### Refresh Token
```go
// this returns a login response just like account.Login() that you can use
// however, it is unneccessary for this example
refreshResponse, err := account.RefreshToken()
if err != nil {
    log.Fatal(err)
}
```

### Logout
```go
if err := account.Logout(); err != nil {
    log.Fatal(err)
}
```

## Network

### List Networks
```go
// writes the list of networks to the account object
err := account.GetListOfNetworks()
// can be accessed via:
networks := account.Networks
```

## Camera

### List Cameras
```go
// returns a slice of blink camera objects
cameras, err := account.GetCameras()
```

## Sync Modules

### List Sync Modules
```go
// writes the list of SyncModules to the account object
err := account.GetSyncModules()
// can be accessed via:
syncmodules := account.SyncModules
```

## System

### Get Manifest
Note: this is the same as hitting the homescreen endpoint
```go
// returns manifest object
manifest, err := account.GetManifest()
```

## Videos Cloud
Note: Cloud operations require the blink subscription

### List Videos Cloud
```go 
pages := 1 // number of pages you want

// returns a video event object
videoEvents, err := account.GetVideoEvents(pages)

// you can access the videos via the following:
videos := videoEvents.Media
```

### Download Videos Cloud
```go 
// retrieves media from: https://%s/api/v1/accounts/%i/media/changed?since=2019-04-19T23:11:20+0000&page=%i
// the pages are the number of media pages you would like to download from

numPages := 1

// explicitly state the filepath you would like this to download to no need to add a / after the end
downloadDir := "currentDir/newDir"
err := DownloadVideosByPages(numPages, downloadDir)

if err != nil {
    log.Fatal(err)
}
```

## Videos Local From Storage
I did not discover this myself, this is from [blinkpy](https://github.com/fronzbot/blinkpy)

The steps for pulling videos from local storage

1. Query sync module for information regarding stored clips
2. Upload the clips to the cloud
3. Download the clips from a cloud URL

Beware the upload/download sequence, there must be a waiting period between the two as the operation is not instantenous

### Read Videos From Local Storage
```go
// this manifest will contain your networks and sync modules
manifest, err := account.GetManifest()

// parse out the sync module you want
syncModule := manifest.SyncModules[0]
// extract ids
syncModuleID := syncModule.ID
networkID := syncModule.NetworkID

// returns local storage manifest object, which contains an array of clips
localStorageManifest, err := account.GetLocalStorageManifest(networkID, syncModuleID)
if err != nil {
    log.Fatal(err)
}

// extract clipID from clip list
clipID := localStorageManifest.Clips[0].ID
```

### Download Videos From Local Storage
```go
// this manifest will contain your networks and sync modules
manifest, err := account.GetManifest()

// parse out the sync module you want
syncModule := manifest.SyncModules[0]
// extract ids
syncModuleID := syncModule.ID
networkID := syncModule.NetworkID

// returns local storage manifest object, which contains an array of clips
localStorageManifest, err := account.GetLocalStorageManifest(networkID, syncModuleID)
if err != nil {
    log.Fatal(err)
}
// iterate through the clips and download them
for i := 0; i < len(localStorageManifest.Clips); i++ {
    clipID := localStorageManifest.Clips[i].ID
    fileName := "someVideoTitle-" + fmt.Sprint(i)
    err = account.DownloadVideoByClipID(networkID, syncModuleID, localStorageManifest.ManifestID, clipID, fileName)

    if err != nil {
        log.Fatal(err)
    }
}
```
