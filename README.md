# blinkgo

Simple library for interacting with blink cameras, mainly: authentication, listing devices/networks/clips, and downloading clips from local storage

This library was made for my purposes but if you would like to see more features open an issue and I will get to it

## Installation

```shell
go get -u github.com/rekram1-node/blinkgo/blink
```

## Table of contents

* [authentication](#auth)
* [fetch user info](#user)
* [list cameras](#camera)
* [download videos](#videos)

## Usage

```shell
example
```

## Local Storage

I did not discover this myself, this is from [blinkpy](https://github.com/fronzbot/blinkpy)

The steps for pulling videos from local storage

1. Query sync module for information regarding stored clips
2. Upload the clips to the cloud
3. Download the clips from a cloud URL

Beware the upload/download sequence, there must be a waiting period between the two as the operation is not instantenous

## Issues

If you have an issue: report it on the [issue tracker](https://github.com/rekram1-node/blinkgo/issues)
