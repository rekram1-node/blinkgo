# blinkgo

[![Go Report](https://goreportcard.com/badge/github.com/rekram1-node/blinkgo)](https://goreportcard.com/report/github.com/rekram1-node/blinkgo) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/rekram1-node/blinkgo/main/LICENSE) ![Build Status](https://github.com/rekram1-node/blinkgo/actions/workflows/main.yml/badge.svg)

![blinkgo](docs/assets/blinkgo-logo.png)

Simple library for interacting with blink cameras, mainly: authentication, listing devices/networks/clips, and downloading clips from local storage

This library was made for my purposes but if you would like to see more features open an issue and I will get to it

## Features

* authentication
* read networks, cameras, sync modules
* list videos
* download videos

## Getting Started

### Prerequisites
- [Go](https://go.dev/)

### Getting blinkgo

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```go
import "github.com/rekram1-node/blinkgo/blink"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following to install the `blinkgo` library

```shell
$ go get -u github.com/rekram1-node/blinkgo/blink
```

## Documentation

Read the [documentation](https:/github.com/rekram1-node/blinkgo/docs/docs.md) for usage instructions


## Authentication

#### Simple login and 2FA verify pin
```go
package main

import (
	"fmt"
	"log"

	"github.com/rekram1-node/blinkgo/blink"
)

func main() {
	email := "example@example.com"
	password := "PLEASE_DON'T_PLAINTEXT_REAL_PASSWORDS"

    // returns account object with: email, password, uuid
	account := blink.NewAccount(email, password)

	// this returns a login response that you can use
	// however, it is unneccessary for this example
	if _, err := account.Login(); err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter Pin: ")
	var pin string
	fmt.Scanln(&pin)

	// this returns a verify pin response that you can use
	// however, it is unneccessary for this example
	if _, err = account.VerifyPin(pin); err != nil {
		log.Fatal(err)
	}
}
```

## Local-Storage

I did not discover this myself, this is from [blinkpy](https://github.com/fronzbot/blinkpy)

The steps for pulling videos from local storage

1. Query sync module for information regarding stored clips
2. Upload the clips to the cloud
3. Download the clips from a cloud URL

Beware the upload/download sequence, there must be a waiting period between the two as the operation is not instantenous

## Issues

If you have an issue: report it on the [issue tracker](https://github.com/rekram1-node/blinkgo/issues)
