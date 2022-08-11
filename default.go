package yc

import (
	"net/http"
	"time"
)

const (
	DefaultRpcPort       = 20443
	DefaultHttpPort      = 20080
	DefaultCallerLevel   = 1000
	SignExpiresInSeconds = 15
)

var (
	httpClient = http.Client{
		Timeout: time.Second * 5,
	}
)
