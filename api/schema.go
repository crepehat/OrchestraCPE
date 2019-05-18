package api

import (
	"net/http"
	"time"
)

var client http.Client

func init() {
	client = http.Client{
		Timeout: 30 * time.Second,
	}
}
