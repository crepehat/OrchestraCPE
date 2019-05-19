package api

import (
	"net/http"
	"time"

	"github.com/crepehat/OrchestraCPE/device"
)

type HeartBeat struct {
	DeviceId string         `json:"id"`                // gives this
	Command  device.Command `json:"command,omitempty"` // receives this
	State    device.State   `json:"state,omitempty"`   // gives this
}

var client http.Client
var configApi, heartbeatApi string

func init() {
	client = http.Client{
		Timeout: 30 * time.Second,
	}
	configApi = "http://ec2-54-245-75-78.us-west-2.compute.amazonaws.com:8088/config"
	heartbeatApi = "http://ec2-54-245-75-78.us-west-2.compute.amazonaws.com:8088/heartbeat"
}
