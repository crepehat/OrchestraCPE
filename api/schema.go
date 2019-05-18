package api

import (
	"net/http"
	"time"
)

var client http.Client
var configApi, heartbeatApi string

func init() {
	client = http.Client{
		Timeout: 30 * time.Second,
	}
	configApi = "127.0.0.1:8080/config"
	heartbeatApi = "127.0.0.1:8080/heartbeat"
}
